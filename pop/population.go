package pop

import (
	"ai-assignment-2/crossover"
	"ai-assignment-2/mutate"
	"ai-assignment-2/selection"
	"ai-assignment-2/variables"
	"math/rand"
)

func InitPopulation(populationSize int) []variables.Individual {
	strategies := []string{variables.AlwaysCooperateStr, variables.AlwaysDefectStr, variables.TitForTatStr, variables.GrimTriggerStr}
	population := make([]variables.Individual, populationSize)
	for i := range population {
		population[i] = variables.Individual{Strategy: strategies[rand.Intn(len(strategies))]}
	}
	return population
}

func EvolvePopulation(population []variables.Individual, tournamentSize int, mutationRate float64) []variables.Individual {
	newPopulation := make([]variables.Individual, len(population))
	for i := range population {
		parent1 := selection.TournamentSelection(population, tournamentSize)
		parent2 := selection.TournamentSelection(population, tournamentSize)
		child := crossover.Crossover(parent1, parent2)
		child = mutate.Mutate(child, mutationRate)
		newPopulation[i] = child
	}
	return newPopulation
}