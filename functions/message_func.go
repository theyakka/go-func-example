package functions

import (
	"log"
	"net/http"

	"mysite.com/myfunctions/common/writers"
)

func OutputMessage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	message := query.Get("message")
	if message == "" {
		message = "Look's like someone forgot to write a message..."
	}
	jw := writers.NewMessageWriter(message)
	jsonString, err := jw.JSONString()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	// all good. write our message.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonString))
}
