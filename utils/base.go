package utils

import (
	"encoding/json"
	"log"
)

const (
	MINING_DIFFICULTY               = 4
	MINING_SENDER                   = "MadBlocks"
	MINING_REWARD                   = 0.5
	VERBOSE                         = true
	MINING_TIMER_SECONDS            = 20
	BLOCKCHAIN_PORT_RANGE_START     = 5000
	BLOCKCHAIN_PORT_RANGE_END       = 5003
	NEIGHBOR_IP_RANGE_START         = 1
	NEIGHBOR_IP_RANGE_END           = 2
	NEIGHBORS_SYNC_TIME_SET_SECONDS = 20
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
