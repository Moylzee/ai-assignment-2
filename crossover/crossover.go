package crossover

import (
	"ai-assignment-2/variables"
	"math/rand/v2"
)

func Crossover(parent1, parent2 variables.Individual) variables.Individual {
	childStrategy := parent1.Strategy
	if rand.Float32() < 0.5 {
		childStrategy = parent2.Strategy
	}
	return variables.Individual{Strategy: childStrategy}
}
