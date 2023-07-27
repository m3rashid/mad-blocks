package utils

const (
	MINING_DIFFICULTY = 4
	MINING_SENDER     = "MadBlocks"
	MINING_REWARD     = 0.5
	VERBOSE           = true
)

type DefaultFuncParamsType struct {
	Verbose bool `json:"verbose"`
}

var DefaultFuncParams = DefaultFuncParamsType{
	Verbose: false,
}
