package amazonpaytesthelper

type AmazonPayConfig struct {
	MerchantID   string
	AccessKey    string
	SecretKey    string
	ClientID     string
	ClientSecret string
	Sandbox      bool
	CurrencyCode string
}

type AmazonPayTestAccount struct {
	Email         string
	EmailPassword string
}
