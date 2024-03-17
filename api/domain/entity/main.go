package entity

type UserInfo struct {
	TaxCode      string `json:"tax_code"`
	Username     string `json:"username"`
	TaxAuthority string `json:"tax_authority"`
	CCCD         string `json:"cccd"`
	DateRange    string `json:"date_range"`
	Status       string `json:"status"`
}
