package variables

type Individual struct {
	Strategy string
	Score    float64
}

const (
	AlwaysCooperateStr = "AlwaysCooperate"
	AlwaysDefectStr    = "AlwaysDefect"
	GrimTriggerStr     = "GrimTrigger"
	TitForTatStr   = "TitForTat"
)

type Variables struct {
	Population     int
	GeneLength     int
	NumGenerations int
	TournamentSize int
	MutationRate   float64
	Noise float64
}
	
func GetVariables() Variables {
	return Variables{
		Population:     100,
		GeneLength:     10,
		NumGenerations: 500,
		TournamentSize: 5,
		MutationRate:   0.1,
		Noise: 0.0,
	}
}

var PayoffMatrix = map[string]map[string][2]int{
	"C": {"C": {3, 3}, "D": {0, 5}},
	"D": {"C": {5, 0}, "D": {1, 1}},
}
