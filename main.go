package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Получаем порт из переменной окружения или используем 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Обработчик для корневого пути
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"status": "ok", "message": "Hello from Go!", "version": "1.0.0"}`)
		w.Header().Set("Content-Type", "application/json")
	})

	// Обработчик для health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"healthy": true}`)
		w.Header().Set("Content-Type", "application/json")
	})

	// Запуск сервера
	fmt.Printf("🚀 Server starting on :%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("❌ Server failed: %v\n", err)
		os.Exit(1)
	}
}
