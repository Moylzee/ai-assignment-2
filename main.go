package main

import (
	"ai-assignment-2/fitness"
	"ai-assignment-2/pop"
	"ai-assignment-2/strategies"
	"ai-assignment-2/variables"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func geneticAlgorithm(vars variables.Variables) {
	population := pop.InitPopulation(vars.Population)
	fixedStrategies := []func(string) string{strategies.AlwaysCooperate, strategies.AlwaysDefect, strategies.TitForTat, strategies.GrimTrigger}

	// CSV File Setup
	file, err := os.Create("fitness_data.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Generation", "AverageFitness"})

	for gen := 0; gen < vars.NumGenerations; gen++ {
		population = fitness.EvaluateFitness(population, fixedStrategies, vars.Noise)

		// Calculate average fitness
		var totalFitness float64
		for _, ind := range population {
			totalFitness += ind.Score
		}
		avgFitness := totalFitness / float64(len(population))
		fmt.Printf("Generation %d: Average Fitness = %.2f\n", gen, avgFitness)

		// Log to CSV
		writer.Write([]string{fmt.Sprintf("%d", gen), fmt.Sprintf("%.2f", avgFitness)})

		// Evolve the population
		population = pop.EvolvePopulation(population, vars.TournamentSize, vars.MutationRate)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	vars := variables.GetVariables()
	geneticAlgorithm(vars)
}
