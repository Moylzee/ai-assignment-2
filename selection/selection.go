package selection

import (
	"ai-assignment-2/variables"
	"math/rand"
)

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