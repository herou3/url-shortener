// пакеты исполняемых приложений должны называться main
package main

import (
	"flag"
	"github.com/herou3/url-shortener/internal/app/internal/config"
	internal "github.com/herou3/url-shortener/internal/app/internal/server"
	"net/http"
)

type configForLaunch struct {
	serverURL string
	baseURL   string
}

// функция main вызывается автоматически при запуске приложения
func main() {
	conf := readConsoleData()
	if err := run(conf.serverURL, conf.baseURL); err != nil {
		panic(err)
	}
}

func readConsoleData() configForLaunch {
	htf := flag.String("a", "localhost:8080", "default host")
	su := flag.String("b", "https://shorturl.ru", "default host")

	flag.Parse()
	return configForLaunch{
		serverURL: *htf,
		baseURL:   *su,
	}
}

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func run(host string, url string) error {
	s := internal.Init()
	config.GetConf().SetConf(map[string]string{"host": host, "shortURL": url})
	configForServer := config.GetConf()
	return http.ListenAndServe(configForServer.HOST, s.Mux)
}
