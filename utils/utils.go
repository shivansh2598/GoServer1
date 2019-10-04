package utils

import (
	"books-list/model"
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter , status int , err model.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}){
	json.NewEncoder(w).Encode(data)
}
