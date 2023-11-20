package config

type Configuration struct {
	HOST     string
	ShortURL string
}

var config *Configuration

func GetConf() *Configuration {
	if config == nil {
		config = &Configuration{
			HOST:     "localhost:8080/",
			ShortURL: "local:8989/",
		}
	}

	return config
}

func (config *Configuration) SetConf(params map[string]string) {
	host, errHost := params["host"]
	if errHost != false {
		config.HOST = host
	}
	shortURL, errUrl := params["shortURL"]
	if errUrl != false {
		config.ShortURL = shortURL
	}
}
