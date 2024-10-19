package model

type RequestModel struct {
	InputNumbers []float64 `json:"input_numbers"`
}

type ResponseModel struct {
	Message string  `json:"message"`
	Average float64 `json:"average"`
}

