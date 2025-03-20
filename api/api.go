package api
import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int
	Balance int64
}

type ErrorMessage struct {
	Code int 
	Msg string
}

func writeError(w http.ResponseWriter, code int, msg string) {
	resp:= ErrorMessage{
		Code: code,
		Msg: msg,
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w,http.StatusBadRequest,err.Error())
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w,http.StatusInternalServerError,"Unexpected Error")
	}
)