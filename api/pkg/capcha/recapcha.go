package capcha

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"megabot/config"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Capcha struct {
	Status  int64 `json:"status"`
	Request int64 `json:"request"`
}

const (
	inputUrl    = "http://azcaptcha.com/in.php"
	responseUrl = "http://azcaptcha.com/res.php"
)

func Recapcha(cfg *config.Environment, capchaBase64 *string) (*string, error) {
	capchaInfo, err := requestReCapcha(capchaBase64, cfg)
	if err != nil {
		return nil, err

	}
	capcha, err := requestGetCapcha(capchaInfo.Request, cfg)
	if err != nil {
		return nil, err
	}
	return capcha, nil
}

func requestGetCapcha(id int64, cfg *config.Environment) (*string, error) {
	time.Sleep(1 * time.Second)
	url := fmt.Sprintf("%s?key=%s&action=get&id=%v", responseUrl, cfg.AzcapchaApiKey, id)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 300 {
		return nil, errors.New(string(body))
	}
	output := strings.Replace(string(body), "OK|", "", -1)
	return &output, nil
}

func requestReCapcha(input *string, cfg *config.Environment) (*Capcha, error) {
	response, err := http.PostForm(inputUrl, url.Values{
		"key":    {cfg.AzcapchaApiKey},
		"body":   {*input},
		"method": {"base64"},
		"json":   {"1"},
	})
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 300 {
		return nil, errors.New(string(body))
	}

	output := &Capcha{}
	err = json.Unmarshal(body, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
