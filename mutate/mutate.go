package mutate

import (
	"ai-assignment-2/variables"
	"math/rand"
)

// Mutate applies a mutation to an individual's strategy based on a mutation rate.
// Params:
// - ind: the individual of type variables.Individual to be mutated
// - mutationRate: a float64 representing the probability of mutation
// Returns:
// - the mutated individual of type variables.Individual
func Mutate(ind variables.Individual, mutationRate float64) variables.Individual {
	strategies := []string{variables.AlwaysCooperateStr, variables.AlwaysDefectStr, variables.TitForTatStr, variables.GrimTriggerStr}
	if rand.Float64() < mutationRate {
		ind.Strategy = strategies[rand.Intn(len(strategies))]
	}
	return ind
}