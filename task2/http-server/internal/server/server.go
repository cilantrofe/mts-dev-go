package server

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const version = "v1.0.0"

type DecodeRequest struct {
	InputString string `json:"inputString"`
}

type DecodeResponse struct {
	OutputString string `json:"outputString"`
}

func Start() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	mux := http.NewServeMux()
	mux.HandleFunc("/version", handleVersion)
	mux.HandleFunc("/decode", handleDecode)
	mux.HandleFunc("/hard-op", handleHardOp)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Сервер завершился с ошибкой: %v\n", err)
		}
	}()

	fmt.Println("Сервер запущен на порту 8080")

	<-ctx.Done()
	stop()

	fmt.Println("Выключение сервера...")
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctxShutdown); err != nil {
		fmt.Printf("Ошибка при выключении: %v\n", err)
	}
	fmt.Println("Сервер остановлен")
}

func handleVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"version": "%s"}`, version)
}

func handleDecode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Разрешен только POST метод", http.StatusMethodNotAllowed)
		return
	}

	var decodeReq DecodeRequest
	if err := json.NewDecoder(r.Body).Decode(&decodeReq); err != nil {
		http.Error(w, "Некорретный JSON", http.StatusBadRequest)
		return
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(decodeReq.InputString)
	if err != nil {
		http.Error(w, "Некорректный base64 ввод", http.StatusBadRequest)
		return
	}

	decodeRes := DecodeResponse{
		OutputString: string(decodedBytes),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(decodeRes)
}

func handleHardOp(w http.ResponseWriter, r *http.Request) {
	delay := time.Duration(10+rand.Intn(11)) * time.Second
	time.Sleep(delay)

	if rand.Intn(2) == 0 {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Успешно")
	}
}
