package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, i interface{}) {
	//baca JSON
	decoder := json.NewDecoder(r.Body)

	//parsing
	err := decoder.Decode(i)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, i interface{}) {
	//memberitahu bahwa bentuknya json
	w.Header().Add("Content-Type", "application/json")

	//menulis webResponse langsung kedalam writer
	encoder := json.NewEncoder(w)
	err := encoder.Encode(i)
	PanicIfError(err)
}
