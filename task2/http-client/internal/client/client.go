package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type DecodeResponse struct {
	OutputString string `json:"outputString"`
}

func Run() {
	client := &http.Client{}

	// GET /version
	resp, err := client.Get("http://localhost:8080/version")
	if err != nil {
		fmt.Printf("Ошибка при запросе версии: %v\n", err)
		return
	}
	defer resp.Body.Close()
	version, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Версия:", string(version))

	// POST /decode
	decodeReq := map[string]string{"inputString": "aGVsbG8gd29ybGQ="}
	jsonData, _ := json.Marshal(decodeReq)
	resp, err = client.Post("http://localhost:8080/decode", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Ошибка при запросе декодирования: %v\n", err)
		return
	}
	defer resp.Body.Close()
	var decodeRes DecodeResponse
	json.NewDecoder(resp.Body).Decode(&decodeRes)
	fmt.Println("Расшифрованная строка:", decodeRes.OutputString)

	// GET /hard-op
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/hard-op", nil)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Запрос был отменен:", err)
		return
	}
	defer resp.Body.Close()
	status := resp.StatusCode == http.StatusOK
	fmt.Printf("Hard-op запрос: успешность=%v, статус код=%d\n", status, resp.StatusCode)
}
