package handler

import (
	"GOLANG/businessLogic"
	"GOLANG/constant"
	"GOLANG/logging"
	"GOLANG/model"
	"encoding/json"
	"net/http"
)

func HandleAverage(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)

	if r.Method != http.MethodPost {
		response := model.ResponseModel{
			Message: constant.ErrorInvalidHttpMethodMessage,
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(response)
		logging.ErrorLog(constant.ErrorInvalidHttpMethodMessage)
		return
	}

	var avg model.RequestModel
	err := json.NewDecoder(r.Body).Decode(&avg) 
	if err != nil {
		response := model.ResponseModel{
			Message: constant.ErrorInvalidInputMessage,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		logging.ErrorLog(constant.ErrorInvalidInputMessage)
		return
	}

	if len(avg.InputNumbers) == 0 {
		response := model.ResponseModel{
			Message: constant.ErrorNoNumbersProvidedMessage,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		logging.ErrorLog(constant.ErrorNoNumbersProvidedMessage)
		return
	}

	inputNumbers := avg.InputNumbers
	evenNumbers, err := businesslogic.Numbers(inputNumbers).EvenNumber()
	if err != nil {
		response := model.ResponseModel{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		logging.ErrorLog(err.Error())
		return
	}

	average, err := evenNumbers.EvenCalulateAverage()
	if err != nil {
		response := model.ResponseModel{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		logging.ErrorLog(err.Error())
		return
	}

	response := model.ResponseModel{
		Message: constant.OperationSuccessfulMessage,
		Average: average,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	logging.LogResponse(w, r)
}
