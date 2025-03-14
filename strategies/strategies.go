package strategies

import (
	"ai-assignment-2/variables"
	"log"
)

func AlwaysCooperate(_ string) string { return "C" }
func AlwaysDefect(_ string) string { return "D" }

// TitForTat returns the opponent's last move, or "C" if there was no last move.
// Params:
// - lastMove: a string representing the opponent's last move
// Returns:
// - the opponent's last move, or "C" if there was no last move
func TitForTat(lastMove string) string {
    if lastMove == "" {
        return "C"
    }
    return lastMove
}

// GrimTrigger returns "D" if the opponent's last move was "D", otherwise returns "C".
// Params:
// - lastMove: a string representing the opponent's last move
// Returns:
// - "D" if the opponent's last move was "D", otherwise "C"
func GrimTrigger(lastMove string) string {
    if lastMove == "D" {
        return "D"
    }
    return "C"
}

// GetStrategy returns the strategy function corresponding to the given strategy name.
// Params:
// - strategy: a string representing the name of the strategy
// Returns:
// - a function that takes a string (the opponent's last move) and returns a string (the next move)
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