package websocket

import (
	"ByZe/pkg/logging"
	"encoding/json"
)

type WSpackage struct {
	Method string
	Type   int
	Data   interface{}
}

func ProcessMessage(message []byte, wsp *WSpackage) []byte {

	if err := json.Unmarshal(message, &wsp); err != nil {
		logging.Warn("can not unmarshal message: %v", err)
	}
	wsp.Method = "test"
	data, _ := json.Marshal(wsp)

	return data
}
