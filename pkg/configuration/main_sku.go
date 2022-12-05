package configuration

type MainSku struct {
	Package       string
	CountryCode   string
	PercentileMin uint
	PercentileMax uint
	Sku           string `json:"main_sku"`
}
