// пакеты исполняемых приложений должны называться main
package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/herou3/url-shortener/internal/config"
	"github.com/herou3/url-shortener/internal/logger"
	internal "github.com/herou3/url-shortener/internal/server"
)

type configForLaunch struct {
	serverURL string
	baseURL   string
}

// функция main вызывается автоматически при запуске приложения
func main() {
	conf := readConsoleData()

	if envServer := os.Getenv("SERVER_ADDRESS"); envServer != "" {
		conf.serverURL = envServer
	}
	if envBaseURL := os.Getenv("BASE_URL"); envBaseURL != "" {
		conf.baseURL = envBaseURL
	}

	if err := run(conf.serverURL, conf.baseURL); err != nil {
		panic(err)
	}
}

func readConsoleData() configForLaunch {
	htf := flag.String("a", "localhost:8080", "default host")
	su := flag.String("b", "localhost:8080", "default host")

	flag.Parse()
	return configForLaunch{
		serverURL: *htf,
		baseURL:   *su,
	}
}

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func run(host string, url string) error {
	s := internal.CreateServerInstance()
	config.GetConf().SetConf(map[string]string{"host": host, "shortURL": url})
	configForServer := config.GetConf()
	err := logger.InitializeLogger("info")
	if err != nil {
		return err
	}
	return http.ListenAndServe(configForServer.HOST, logger.WithLogging(s.Mux))
}
