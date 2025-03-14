package fitness

import (
	"ai-assignment-2/match"
	"ai-assignment-2/strategies"
	"ai-assignment-2/variables"
)

// EvaluateFitness evaluates the fitness of a population of individuals.
// Params:
// - population: a slice of variables.Individual representing the population to be evaluated
// - fixedStrategies: a slice of functions that take a string and return a string, representing fixed strategies to play against
// - noise: a float64 representing the noise level in the match
// Returns:
// - a slice of variables.Individual with updated scores based on their performance against the fixed strategies
func EvaluateFitness(population []variables.Individual, fixedStrategies []func(string) string, noise float64) []variables.Individual {
    for i := range population {
        population[i].Score = 0
        for _, fixedStrategy := range fixedStrategies {
            score := match.PlayMatch(strategies.GetStrategy(population[i].Strategy), fixedStrategy, 50, noise)
            population[i].Score += score
        }
    }
    return population
}