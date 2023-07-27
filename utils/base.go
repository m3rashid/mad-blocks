package utils

import (
	"encoding/json"
	"log"
)

const (
	MINING_DIFFICULTY = 4
	MINING_SENDER     = "MadBlocks"
	MINING_REWARD     = 0.5
	VERBOSE           = true
)

func JsonStatus(message string) string {
	data, err := json.Marshal(struct {
		Message string `json:"message"`
	}{Message: message})
	if err != nil {
		log.Println(err.Error())
	}
	return string(data)
}
