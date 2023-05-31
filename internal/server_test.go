package internal

import (
	"encoding/json"
	"fmt"
	"github.com/steinfletcher/apitest"
	"go-labs-game-platform/internal/bootstrap"
	"go-labs-game-platform/internal/config"
	"go-labs-game-platform/internal/httpserver"
	"go-labs-game-platform/internal/models"
	"net/http"
	"testing"
	"time"
)

type ResponseObject struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Приклад API тестів
// Використовуються для перевірки функціоналу на всіх рівнях від інтерфейсу до взаємодії з іншими системами.
// Є два варіанти реалізувати такі тести
// 1) створити http клієнт та робити запити напряму на запущену систему
// 2) підняти тестовий сервер за допомогою тестів, та робити запити за допомогою бібліотек, (наприклад github.com/steinfletcher/apitest)
// Перший варіант кращий тим що тестується реальне середовище, його не потрібно додатково налаштовувати в тестах.
// Другий варіант дозволяє більш гнучко налаштовувати сервер, при потребі частину даних можна емулювати, але
// це вимагає додаткового налаштування в тестах, та не дозволяє перевіряти працюючий сервер, що інколи може бути корисно.
// Тут розглядається другий варіант.

// Для різних видів тестів є сенс визначати теги збірки.
// Якщо у файлі визначений тег збірки, цей файл виключається зі збірки якщо він явно не вказаний.
// Запустити такий тест можна командою `go test ./... -tags=api_test`
// Також можна використовувати теги для визначення окремих тестів, наприклад `go test ./... -tags=api_test -run=TestAll`
// Якщо тести запускаються з IDE, то там є можливість вказати теги збірки.

var dbURI string

func TestAll(t *testing.T) {

	deps, err := bootstrap.Up()
	if err != nil {
		panic(err)
	}
	s := httpserver.New(deps)
	server := &http.Server{
		Addr:           s.Addr(),
		Handler:        s.Router(),
		ReadTimeout:    config.Get().HTTP.ReadTimeout,
		WriteTimeout:   config.Get().HTTP.WriteTimeout,
		IdleTimeout:    time.Second * 10,
		MaxHeaderBytes: 256,
	}
	go server.ListenAndServe()

	client := &http.Client{}
	req, err := http.NewRequest(
		"POST", "http://localhost:8080/api/v1/room/create", nil,
	)
	req.Header.Add("Authorization", "Bearer ACAGNAAPKNFONQJLZBOPAN6MTI")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Parse the response JSON into an object
	var responseObject models.Room
	err = json.NewDecoder(resp.Body).Decode(&responseObject)
	if err != nil {
		fmt.Println("Error parsing response:", err)
		return
	}

	// Print the response object
	fmt.Println("Response object:", responseObject.ID)

	router := s.Router()

	//Тестування створення ордеру неуспішне
	apitest.New().
		Handler(router).
		Post("/api/v1/register").
		Body(`{"username": "test", "password": "test"}`).
		Expect(t).
		Status(http.StatusInternalServerError).
		End()

	apitest.New().
		Handler(router).
		Post("/api/v1/login").
		Body(`{"username": "test", "password": "test"}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	//Тестування створення ордеру
	apitest.New().
		Handler(router).
		Get("/api/v1/room").
		Headers(map[string]string{"accept:": "application/json", "Authorization": "Bearer ACAGNAAPKNFONQJLZBOPAN6MTI"}).
		Expect(t).
		Status(http.StatusOK).
		End()

	// Create New Room
	apitest.New().
		Handler(router).
		Post("/api/v1/room/create").
		Headers(map[string]string{"accept:": "application/json", "Authorization": "Bearer ACAGNAAPKNFONQJLZBOPAN6MTI"}).
		Expect(t).
		Status(http.StatusCreated).
		End()

	// Delete Room
	apitest.New().
		Handler(router).
		Delete("/api/v1/room/" + responseObject.ID.String()).
		Headers(map[string]string{"accept:": "application/json", "Authorization": "Bearer ACAGNAAPKNFONQJLZBOPAN6MTI"}).
		Expect(t).
		Status(http.StatusOK).
		End()

}
