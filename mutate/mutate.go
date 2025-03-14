package mutate

import (
	"ai-assignment-2/variables"
	"math/rand"
)

func Mutate(ind variables.Individual, mutationRate float64) variables.Individual {
	strategies := []string{variables.AlwaysCooperateStr, variables.AlwaysDefectStr, variables.TitForTatStr, variables.GrimTriggerStr}
	if rand.Float64() < mutationRate {
		ind.Strategy = strategies[rand.Intn(len(strategies))]
	}
	return ind
}