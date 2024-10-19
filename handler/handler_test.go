package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"GOLANG/constant"
	// "GOLANG/handler"
	"GOLANG/model"
)

func TestAverageCalculationWithValidInput(t *testing.T) {
	inputData := model.RequestModel{InputNumbers: []float64{5, 10, 15, 20}}
	requestBody, err := json.Marshal(inputData)
	if err != nil {
		t.Error(err)
	}
	request := httptest.NewRequest(http.MethodPost, constant.AverageEndpoint, bytes.NewReader(requestBody))
	recorder := httptest.NewRecorder()

	HandleAverage(recorder, request)

	response := recorder.Result()
	var responseBody model.ResponseModel
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatalf(constant.ErrorResponseDecodeMessage, err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf(constant.ExpectedHttpStatusOKMessage, response.StatusCode)
	}

	expectedAverage := 15.0
	if responseBody.Average != expectedAverage {
		t.Errorf(constant.ExpectedAverageValueMessage, expectedAverage, responseBody.Average)
	}
	if responseBody.Message != constant.OperationSuccessfulMessage {
		t.Errorf(constant.ExpectedResponseMessageMessage, constant.OperationSuccessfulMessage, responseBody.Message)
	}
}

func TestAverageHandlerWithInvalidMethod(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, constant.AverageEndpoint, nil)
	recorder := httptest.NewRecorder()

	HandleAverage(recorder, request)

	response := recorder.Result()
	if response.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf(constant.ExpectedHttpMethodNotAllowedMessage, response.StatusCode)
	}
}

func TestAverageHandlerWithEmptyBody(t *testing.T) {

	request := httptest.NewRequest(http.MethodPost, constant.AverageEndpoint, nil)
	recorder := httptest.NewRecorder()

	HandleAverage(recorder, request)

	response := recorder.Result()
	var responseBody model.ResponseModel
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatalf(constant.ErrorResponseDecodeMessage, err)
	}

	if response.StatusCode != http.StatusBadRequest {
		t.Errorf(constant.ExpectedHttpStatusBadRequestMessage, response.StatusCode)
	}

	if responseBody.Message != constant.ErrorInvalidInputMessage {

		t.Errorf(constant.ErrorMessageFormat, constant.ErrorInvalidInputMessage, responseBody.Message)
	}
}

func TestAverageHandlerWithNoValue(t *testing.T) {
	inputData := model.RequestModel{InputNumbers: []float64{}}
	requestBody, err := json.Marshal(inputData)
	if err != nil {
		t.Error(err)
	}
	request := httptest.NewRequest(http.MethodPost, constant.AverageEndpoint, bytes.NewReader(requestBody))
	recorder := httptest.NewRecorder()
	HandleAverage(recorder, request)

	response := recorder.Result()
	var responseBody model.ResponseModel
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatalf(constant.ErrorResponseDecodeMessage, err)
	}

	if response.StatusCode != http.StatusBadRequest {
		t.Errorf(constant.ExpectedHttpStatusBadRequestMessage, response.StatusCode)
	}

	if constant.ErrorNoNumbersProvidedMessage != responseBody.Message {
		t.Errorf(constant.ErrorMessageFormat, constant.ErrorNoNumbersProvidedMessage, responseBody.Message)

	}

}

func TestAverageHandlerWithInvalidJSON(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, constant.AverageEndpoint, bytes.NewBufferString("invalid JSON"))
	recorder := httptest.NewRecorder()

	HandleAverage(recorder, request)

	response := recorder.Result()
	var responseBody model.ResponseModel
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatalf(constant.ErrorResponseDecodeMessage, err)
	}

	if response.StatusCode != http.StatusBadRequest {
		t.Errorf(constant.ExpectedHttpStatusBadRequestMessage, response.StatusCode)
	}

	if responseBody.Message != constant.ErrorInvalidInputMessage {
		t.Errorf(constant.ErrorMessageFormat, constant.ErrorInvalidInputMessage, responseBody.Message)
	}
}

func TestAverageHandlerWithNoEvenNumbers(t *testing.T) {
	inputData := model.RequestModel{InputNumbers: []float64{1, 3, 5}}
	requestBody, _ := json.Marshal(inputData)
	request := httptest.NewRequest(http.MethodPost, constant.AverageEndpoint, bytes.NewReader(requestBody))
	recorder := httptest.NewRecorder()

	HandleAverage(recorder, request)

	response := recorder.Result()
	var responseBody model.ResponseModel
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatalf(constant.ErrorResponseDecodeMessage, err)
	}

	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf(constant.ExpectedHttpStatusInternalServerMessage, response.StatusCode)
	}

	if responseBody.Message != constant.ErrorNoNumbersProvidedMessage {
		t.Errorf(constant.ErrorMessageFormat, constant.ErrorNoNumbersProvidedMessage, responseBody.Message)
	}
}

func TestAverageHandlerWithNegativeNumbers(t *testing.T) {
	inputData := model.RequestModel{InputNumbers: []float64{-4, -2, 2, 4}}
	requestBody, _ := json.Marshal(inputData)
	request := httptest.NewRequest(http.MethodPost, constant.AverageEndpoint, bytes.NewReader(requestBody))
	recorder := httptest.NewRecorder()

	HandleAverage(recorder, request)

	response := recorder.Result()
	var responseBody model.ResponseModel
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		t.Fatalf(constant.ErrorResponseDecodeMessage, err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf(constant.ExpectedHttpStatusOKMessage, response.StatusCode)
	}

	expectedAverage := 0.0
	if responseBody.Average != expectedAverage {
		t.Errorf(constant.ExpectedAverageValueMessage, expectedAverage, responseBody.Average)
	}
}
