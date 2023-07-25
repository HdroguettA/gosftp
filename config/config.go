package config

type SftpConfig struct {
	Host     string
	Port     string
	Username string
	Password string
}

func NewConfig() *SftpConfig {
	return &SftpConfig{
		Host:     "",
		Port:     "",
		Username: "",
		Password: "",
	}
}
