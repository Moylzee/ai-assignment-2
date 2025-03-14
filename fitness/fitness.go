package fitness

import (
	"ai-assignment-2/match"
	"ai-assignment-2/strategies"
	"ai-assignment-2/variables"
)

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