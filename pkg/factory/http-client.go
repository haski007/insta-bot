package factory

type HTTPClient struct {
	Addr             string `json:"addr" yaml:"addr"`
	Username         string `json:"username" yaml:"username"`
	Password         string `json:"password" yaml:"password"`
	VerificationCode string `json:"verification_code" yaml:"verification_code"`
}
