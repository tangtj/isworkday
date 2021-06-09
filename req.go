package main

import (
	"encoding/json"
	"net/http"
)

func Req(writer http.ResponseWriter, request *http.Request) {
	day, err := Today()
	if err != nil {
		bytes, _ := json.Marshal(map[string]interface{}{
			"ok":  false,
			"msg": err.Error(),
		})
		writer.Write(bytes)
		return
	} else {
		bytes, _ := json.Marshal(struct {
			Ok bool `json:"ok"`
			*Day
		}{
			Ok:  true,
			Day: day,
		})
		writer.Write(bytes)
		return
	}
}
