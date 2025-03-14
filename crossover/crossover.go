package crossover

import (
    "ai-assignment-2/variables"
    "math/rand/v2"
)

// Crossover takes two parent individuals and produces a child individual.
// Params:
// - parent1: the first parent individual of type variables.Individual
// - parent2: the second parent individual of type variables.Individual
// Returns:
// - a new individual of type variables.Individual with a strategy inherited from one of the parents
func Crossover(parent1, parent2 variables.Individual) variables.Individual {
    childStrategy := parent1.Strategy
    if rand.Float32() < 0.5 {
        childStrategy = parent2.Strategy
    }
    return variables.Individual{Strategy: childStrategy}
}