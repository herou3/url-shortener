// пакеты исполняемых приложений должны называться main
package main

import (
	internal "github.com/herou3/url-shortener/internal/app/internal/server"
	"net/http"
)

// функция main вызывается автоматически при запуске приложения
func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {

	s := internal.Init()

	return http.ListenAndServe("127.0.0.1:8080", s.Mux)
}
