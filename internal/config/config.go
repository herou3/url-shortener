package config

import (
	"regexp"
)

type Configuration struct {
	HOST     string
	ShortURL string
}

var config *Configuration

func GetConf() *Configuration {
	if config == nil {
		config = &Configuration{
			HOST:     "localhost:8080",
			ShortURL: "localhost:8080",
		}
	}

	return config
}

func (config *Configuration) SetConf(params map[string]string) {
	host, errHost := params["host"]
	if errHost {
		config.HOST = host
	}
	shortURL, errURL := params["shortURL"]
	if errURL {
		isMatch, _ := regexp.MatchString("((http|https)://)", shortURL)
		if !isMatch {
			shortURL = "http://" + shortURL
		}
		config.ShortURL = shortURL
	}
}
