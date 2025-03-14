package strategies

import (
	"ai-assignment-2/variables"
	"log"
)

func AlwaysCooperate(_ string) string { return "C" }
func AlwaysDefect(_ string) string    { return "D" }
func TitForTat(lastMove string) string {
	if lastMove == "" {
		return "C"
	}
	return lastMove
}
func GrimTrigger(lastMove string) string {
	if lastMove == "D" {
		return "D"
	}
	return "C"
}

func GetStrategy(strategy string) func(string) string {
	switch strategy {
	case variables.AlwaysCooperateStr:
		return AlwaysCooperate
	case variables.AlwaysDefectStr:
		return AlwaysDefect
	case variables.TitForTatStr:
		return TitForTat
	case variables.GrimTriggerStr:
		return GrimTrigger
	}
	log.Println("Invalid strategy, defaulting to TitForTat")
	return TitForTat
}