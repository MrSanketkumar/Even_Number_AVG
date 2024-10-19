package handler

import (
	"GOLANG/constant"
	// "GOLANG/handler"
	"GOLANG/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkAverageHandler(b *testing.B) {
	inputData := model.RequestModel{InputNumbers: []float64{5, 10, 15, 20}}
	requestBody, _ := json.Marshal(inputData)

	for i := 0; i < b.N; i++ {
		request := httptest.NewRequest(http.MethodPost, constant.AverageEndpoint, bytes.NewReader(requestBody))
		recorder := httptest.NewRecorder()

		HandleAverage(recorder, request)
	}
}

func BenchmarkAverageHandlerWithInvalidMethod(b *testing.B) {
	for i := 0; i < b.N; i++ {
		request := httptest.NewRequest(http.MethodGet, constant.AverageEndpoint, nil)
		recorder := httptest.NewRecorder()
		HandleAverage(recorder, request)
	}
}
