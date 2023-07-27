package utils

const (
	MINING_DIFFICULTY = 4
	MINING_SENDER     = "MadBlocks"
	MINING_REWARD     = 1.0
	VERBOSE           = true
)

type DefaultFuncParams struct {
	Verbose bool `json:"verbose"`
}
