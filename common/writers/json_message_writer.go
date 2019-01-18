package writers

import (
	"encoding/json"
)

type JSONMessageWriter struct {
	Message string `json:"message"`
}

func NewMessageWriter(message string) *JSONMessageWriter {
	return &JSONMessageWriter{
		Message: message,
	}
}

func (jw *JSONMessageWriter) JSONString() (string, error) {
	messageResponse := map[string]interface{}{
		"data": map[string]string{
			"message": jw.Message,
		},
	}
	bytesValue, err := json.Marshal(messageResponse)
	if err != nil {
		return "", err
	}
	return string(bytesValue), nil
}
