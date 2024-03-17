package usecase

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"megabot/domain/entity"
	capcha2 "megabot/pkg/capcha"
)

type TaxCodeInfo struct {
	CCCD string `bson:"cccd" json:"cccd" binding:"required"`
}

func (t *Config) Search(taxCodeInput string) ([]entity.UserInfo, error) {
	id := uuid.New().String()
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		t.log.Error(id, fmt.Sprintf("init selenium error"), err.Error())

	}

	defer service.Stop()

	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless", // comment out this line for testing
		"--window-size=1920,1080",
		"--no-sandbox",
		"--disable-extensions",
		"--dns-prefetch-disable",
		"--disable-gpu",
	}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		t.log.Error(id, fmt.Sprintf("create a new remote client error"), err.Error())
		return nil, err
	}

	// visit the target page
	err = driver.Get("https://tracuunnt.gdt.gov.vn/tcnnt/mstcn.jsp")
	if err != nil {
		t.log.Error(id, fmt.Sprintf("visit the target page error"), err.Error())
		return nil, err
	}

	// 2.25s
	// get btn Mã số thuế
	inputCCCD, err := driver.FindElement(selenium.ByName, "cmt2")
	if err != nil {
		t.log.Error(id, fmt.Sprintf("get btn Mã số thuế error"), err.Error())
		return nil, err
	}
	if err := inputCCCD.SendKeys(taxCodeInput); err != nil {
		t.log.Error(id, fmt.Sprintf("set btn Mã số thuế error"), err.Error())
		return nil, err
	}

	data, err := t.getCapchaAndHandel(driver, ".inputBtn[value='Tra cứu']", 5)
	if err != nil {
		t.log.Error(id, fmt.Sprintf("get captcha error"), err.Error())
		return nil, err
	}

	return data, nil
}

func (t *Config) getCapchaAndHandel(driver selenium.WebDriver, nameBtn string, retry int) ([]entity.UserInfo, error) {
	var error error
	for i := 0; i < retry; i++ {
		// get capcha
		capcha, err := driver.FindElement(selenium.ByXPATH, "/html/body/div/div[1]/div[4]/div[2]/div[2]/div/div/form/table/tbody/tr[6]/td[2]/table/tbody/tr/td[2]/div/img")
		if err != nil {
			error = err
		}
		capchaScreenShot, err := capcha.Screenshot(true)
		if err != nil {
			error = err
		}

		imgBase64Str := base64.StdEncoding.EncodeToString(capchaScreenShot)
		capchaCode, err := capcha2.Recapcha(t.cfg, &imgBase64Str)
		if err != nil {
			error = err
		}

		// set input capcha
		inputCapcha, err := driver.FindElement(selenium.ByCSSSelector, "#captcha")
		if err != nil {
			error = err
		}
		if err := inputCapcha.SendKeys(*capchaCode); err != nil {
			error = err
		}

		// click button tra cuu
		btnSearch, err := driver.FindElement(selenium.ByCSSSelector, ".subBtn")
		if err != nil {
			error = err
		}
		if err := btnSearch.Click(); err != nil {
			error = err
		}

		data, err := driver.FindElement(selenium.ByXPATH, "/html/body/div/div[1]/div[4]/div[2]/div[2]/div/div/table")
		if err != nil {
			error = err
		}

		// check error after submit captcha code
		if i <= retry && error != nil {
			_ = inputCapcha.SendKeys("")
			error = nil
			continue
		}

		userInfos, err := data.FindElements(selenium.ByTagName, "td")
		if err != nil {
			return nil, err
		}

		var infoTaxCodes []entity.UserInfo
		var infoTaxCode entity.UserInfo
		var index = 0
		for i2, info := range userInfos {
			if i2 == index {
				continue
			}

			if i2 == index+1 {
				infoTaxCode.TaxCode, _ = info.Text()
				continue
			}
			if i2 == index+2 {
				infoTaxCode.Username, _ = info.Text()
				continue
			}
			if i2 == index+3 {
				infoTaxCode.TaxAuthority, _ = info.Text()
				continue
			}
			if i2 == index+4 {
				infoTaxCode.CCCD, _ = info.Text()
				continue
			}
			if i2 == index+5 {
				infoTaxCode.DateRange, _ = info.Text()
				continue
			}
			if i2 == index+6 {
				infoTaxCode.Status, _ = info.Text()
				continue
			}

			if i2 == index+7 {
				infoTaxCodes = append(infoTaxCodes, infoTaxCode)
				infoTaxCode = entity.UserInfo{}
				index = index + 7
			}
		}
		return infoTaxCodes, nil
	}

	return nil, nil
}
