package config

type VmessConfig struct {
	Add      string `json:"add"`
	Id       string `json:"id"`
	Port     string `json:"port"`
	Ps       string `json:"ps"`
	Security string `json:"security"`
	Net      string `json:"tcp"`
	Sni      string `json:"sni"`
	V        string `json:"v"`
	Fp       string `json:"fp"`
	Type     string `json:"type"`
	Aid      string `json:"aid"`
	Host     string `json:"host"`
	Tls      string `json:"tls"`
}
