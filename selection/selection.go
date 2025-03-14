package selection

import (
	"ai-assignment-2/variables"
	"math/rand"
)

// TournamentSelection selects the best individual from a randomly chosen subset of the population.
// Params:
// - population: a slice of variables.Individual representing the current population
// - tournamentSize: an integer representing the size of the tournament
// Returns:
// - the best individual of type variables.Individual from the tournament
func TournamentSelection(population []variables.Individual, tournamentSize int) variables.Individual {
	best := population[rand.Intn(len(population))]
	for i := 1; i < tournamentSize; i++ {
		competitor := population[rand.Intn(len(population))]
		if competitor.Score > best.Score {
			best = competitor
		}
	}
	return best
}