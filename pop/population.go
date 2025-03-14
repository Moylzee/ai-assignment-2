package pop

import (
	"ai-assignment-2/crossover"
	"ai-assignment-2/mutate"
	"ai-assignment-2/selection"
	"ai-assignment-2/variables"
	"math/rand"
)

// InitPopulation initializes a population of individuals with random strategies.
// Params:
// - populationSize: an integer representing the size of the population to be initialized
// Returns:
// - a slice of variables.Individual representing the initialized population
func InitPopulation(populationSize int) []variables.Individual {
    strategies := []string{variables.AlwaysCooperateStr, variables.AlwaysDefectStr, variables.TitForTatStr, variables.GrimTriggerStr}
    population := make([]variables.Individual, populationSize)
    for i := range population {
        population[i] = variables.Individual{Strategy: strategies[rand.Intn(len(strategies))]}
    }
    return population
}

// EvolvePopulation evolves a population of individuals through selection, crossover, and mutation.
// Params:
// - population: a slice of variables.Individual representing the current population
// - tournamentSize: an integer representing the size of the tournament for selection
// - mutationRate: a float64 representing the probability of mutation
// Returns:
// - a slice of variables.Individual representing the evolved population
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