package main

import (
	"ai-assignment-2/vars"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Action int

const (
	Cooperate Action = iota
	Defect
)

type Strategy interface {
	Play(opponentHistory []Action) Action
	Name() string
}

type AlwaysCooperate struct{}

func (ac AlwaysCooperate) Play(_ []Action) Action {
	return Cooperate
}

func (ac AlwaysCooperate) Name() string {
	return "Always Cooperate"
}

type AlwaysDefect struct{}

func (ad AlwaysDefect) Play(_ []Action) Action {
	return Defect
}

func (ad AlwaysDefect) Name() string {
	return "Always Defect"
}

type TitForTat struct{}

func (tft TitForTat) Play(history []Action) Action {
	if len(history) == 0 {
		return Cooperate
	}
	return history[len(history)-1]
}

func (tft TitForTat) Name() string {
	return "Tit-for-Tat"
}

var payoffMatrix = map[Action]map[Action][2]int{
	Cooperate: {Cooperate: {3, 3}, Defect: {0, 5}},
	Defect:    {Cooperate: {5, 0}, Defect: {1, 1}},
}

type Individual struct {
	Gene   []Action
	Score  int
}

func playMatch(s1, s2 Strategy, rounds int) (int, int) {
	history1 := []Action{}
	history2 := []Action{}
	
	score1, score2 := 0, 0
	for i := 0; i < rounds; i++ {
		move1 := s1.Play(history2)
		move2 := s2.Play(history1)
		result := payoffMatrix[move1][move2]
		score1 += result[0]
		score2 += result[1]
		history1 = append(history1, move1)
		history2 = append(history2, move2)
	}
	return score1, score2
}

func generateRandomIndividual(length int) Individual {
	gene := make([]Action, length)
	for i := range gene {
		gene[i] = Action(rand.Intn(2))
	}
	return Individual{Gene: gene, Score: 0}
}

func evaluateFitness(population []Individual, opponents []Strategy) {
	for i := range population {
		totalScore := 0
		for _, opponent := range opponents {
			individualStrategy := RandomStrategy{Gene: population[i].Gene}
			score, _ := playMatch(individualStrategy, opponent, 100)
			totalScore += score
		}
		population[i].Score = totalScore
	}
}

type RandomStrategy struct {
	Gene []Action
}

func (rs RandomStrategy) Play(history []Action) Action {
	if len(history) < len(rs.Gene) {
		return rs.Gene[len(history)]
	}
	return rs.Gene[rand.Intn(len(rs.Gene))]
}

func (rs RandomStrategy) Name() string {
	return "Evolved Strategy"
}

func tournamentSelection(pop []Individual, k int) Individual {
	best := pop[rand.Intn(len(pop))]
	for i := 1; i < k; i++ {
		contender := pop[rand.Intn(len(pop))]
		if contender.Score > best.Score {
			best = contender
		}
	}
	return best
}

func crossover(parent1, parent2 Individual) Individual {
	point := rand.Intn(len(parent1.Gene))
	newGene := append(parent1.Gene[:point], parent2.Gene[point:]...)
	return Individual{Gene: newGene, Score: 0}
}

func mutate(ind Individual, mutationRate float64) Individual {
	for i := range ind.Gene {
		if rand.Float64() < mutationRate {
			ind.Gene[i] = Action(rand.Intn(2))
		}
	}
	return ind
}

func calculateAverageScore(population []Individual) float64 {
	totalScore := 0
	for _, individual := range population {
		totalScore += individual.Score
	}
	return float64(totalScore) / float64(len(population))
}

func logScoresToFile(filename string, generation int, bestScore int, averageScore float64) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	logLine := fmt.Sprintf("Generation %d: Best Score = %d, Average Score = %.2f\n", generation, bestScore, averageScore)
	if _, err := file.WriteString(logLine); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

func evolvePopulation(population []Individual, opponents []Strategy, generations int) {
	logFilename := "evolution_log.txt"
	elitismCount := len(population) / 10 // Preserve top 10%
	for gen := 0; gen < generations; gen++ {
		evaluateFitness(population, opponents)
		sort.Slice(population, func(i, j int) bool {
			return population[i].Score > population[j].Score
		})
		
		bestScore := population[0].Score
		averageScore := calculateAverageScore(population)
		fmt.Printf("Generation %d: Best Score = %d, Average Score = %.2f\n", gen, bestScore, averageScore)
		logScoresToFile(logFilename, gen, bestScore, averageScore)

		newPopulation := make([]Individual, 0, len(population))
		newPopulation = append(newPopulation, population[:elitismCount]...)
		
		for i := elitismCount; i < len(population); i++ {
			parent1 := tournamentSelection(population, 3)
			parent2 := tournamentSelection(population, 3)
			child := crossover(parent1, parent2)
			child = mutate(child, 0.02)
			newPopulation = append(newPopulation, child)
		}
		population = newPopulation
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	strategies := []Strategy{AlwaysCooperate{}, AlwaysDefect{}, TitForTat{}}
	
	variables := vars.GetVariables()

	population := make([]Individual, variables.Population)
	for i := range population {
		population[i] = generateRandomIndividual(variables.GeneLength)
	}
	evolvePopulation(population, strategies, variables.NumGenerations)
}
