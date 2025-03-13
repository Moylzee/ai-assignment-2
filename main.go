package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	AlwaysCooperateStr = "AlwaysCooperate"
	AlwaysDefectStr    = "AlwaysDefect"
	GrimTriggerStr     = "GrimTrigger"
	TitForTatStr       = "TitForTat"
)


var (
	fixedStrategies = []func(string) string{AlwaysCooperate, AlwaysDefect, TitForTat, GrimTrigger}
	strategies = []string{AlwaysCooperateStr, AlwaysDefectStr, TitForTatStr, GrimTriggerStr}
)

type Variables struct {
	Population     int
	NumGenerations int
	TournamentSize int
	MutationRate   float64
	Noise          float64
}

func GetVariables() Variables {
	return Variables{
		Population:     100,
		NumGenerations: 100,
		TournamentSize: 5,
		MutationRate:   0.1,
		Noise:          0.8,
	}
}

type Individual struct {
	Strategies []string
	Score      float64
}

var payoffMatrix = map[string]map[string][2]int{
	"C": {"C": {3, 3}, "D": {0, 5}},
	"D": {"C": {5, 0}, "D": {1, 1}},
}

func AlwaysCooperate(_ string) string { return "C" }
func AlwaysDefect(_ string) string    { return "D" }
func TitForTat(lastMove string) string {
	if lastMove == "" {
		return "C"
	}
	return lastMove
}
func GrimTrigger(lastMove string) string {
	if lastMove == "D" {
		return "D"
	}
	return "C"
}

func getStrategyFunction(strategy string) func(string) string {
	switch strategy {
	case AlwaysCooperateStr:
		return AlwaysCooperate
	case AlwaysDefectStr:
		return AlwaysDefect
	case TitForTatStr:
		return TitForTat
	case GrimTriggerStr:
		return GrimTrigger
	}
	log.Println("Invalid strategy, defaulting to TitForTat")
	return TitForTat
}

func initPopulation(populationSize int) []Individual {
	population := make([]Individual, populationSize)

	for i := range population {

		// Pick 2 Unique Strategies and Assign them to that individual
		strat1 := rand.Intn(len(strategies))
		strat2 := rand.Intn(len(strategies))
		if strat1 == strat2 {
			strat2 = (strat2 + 1) % len(strategies)
		}
		
		strats := []string{strategies[strat1], strategies[strat2]}
		population[i] = Individual{Strategies: strats}
	}
	return population
}

func playMatch(agent Individual, opponent func(string) string, rounds int, noise float64) float64 {
	var score int
	var lastMove1, lastMove2 string
	for i := 0; i < rounds; i++ {
		currentStrategy := getStrategyFunction(agent.Strategies[rand.Intn(len(agent.Strategies))])

		if rand.Float64() < noise {
			m1 := currentStrategy(lastMove1)
			m2 := []string{"C", "D"}[rand.Intn(2)]
			score += payoffMatrix[m1][m2][0]
		} else {
			m1, m2 := currentStrategy(lastMove2), opponent(lastMove1)
			score += payoffMatrix[m1][m2][0]
			lastMove1, lastMove2 = m1, m2
		}
	}
	return float64(score) / float64(rounds)
}

func evaluateFitness(population []Individual, fixedStrategies []func(string) string, noise float64) []Individual {
	for i := range population {
		population[i].Score = 0
		for _, fixedStrategy := range fixedStrategies {
			score := playMatch(population[i], fixedStrategy, 50, noise)
			population[i].Score += score
		}
	}
	return population
}

func crossover(parent1, parent2 Individual) Individual {
	childStrategies := append(parent1.Strategies[:len(parent1.Strategies)/2], parent2.Strategies[len(parent2.Strategies)/2:]...)
	return Individual{Strategies: childStrategies}
}

func mutate(ind Individual, mutationRate float64) Individual {
	if rand.Float64() < mutationRate {
		newStrategy := strategies[rand.Intn(len(strategies))]
		ind.Strategies = append(ind.Strategies, newStrategy)
	}
	return ind
}

func TournamentSelection(population []Individual, tournamentSize int) Individual {
	best := population[rand.Intn(len(population))]
	for i := 1; i < tournamentSize; i++ {
		competitor := population[rand.Intn(len(population))]
		if competitor.Score > best.Score {
			best = competitor
		}
	}
	return best
}

func evolvePopulation(population []Individual, tournamentSize int, mutationRate float64) []Individual {
	newPopulation := make([]Individual, len(population))
	for i := range population {
		parent1 := TournamentSelection(population, tournamentSize)
		parent2 := TournamentSelection(population, tournamentSize)
		child := crossover(parent1, parent2)
		child = mutate(child, mutationRate)
		newPopulation[i] = child
	}
	return newPopulation
}

func geneticAlgorithm(vars Variables) {
	population := initPopulation(vars.Population)

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
		population = evaluateFitness(population, fixedStrategies, vars.Noise)

		var totalFitness float64
		for _, ind := range population {
			totalFitness += ind.Score
		}
		avgFitness := totalFitness / float64(len(population))
		fmt.Printf("Generation %d: Average Fitness = %.2f\n", gen, avgFitness)
		writer.Write([]string{fmt.Sprintf("%d", gen), fmt.Sprintf("%.2f", avgFitness)})

		population = evolvePopulation(population, vars.TournamentSize, vars.MutationRate)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	vars := GetVariables()
	geneticAlgorithm(vars)
}
