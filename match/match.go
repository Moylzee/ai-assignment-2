package match

import (
	"ai-assignment-2/variables"
	"math/rand"
)

func PlayMatch(s1, s2 func(string) string, rounds int, noise float64) float64 {
	var score int
	var lastMove1, lastMove2 string
	var m1, m2 string
	for i := 0; i < rounds; i++ {

		// Introduce noise: random choice if noise < random value
		if rand.Float64() > noise {
			m1 := []string{"C", "D"}[rand.Intn(2)]
			m2 := []string{"C", "D"}[rand.Intn(2)]
			score += variables.PayoffMatrix[m1][m2][0]
		} else {
			m1, m2 := s1(lastMove2), s2(lastMove1)
			score += variables.PayoffMatrix[m1][m2][0]
		}
		lastMove1, lastMove2 = m1, m2
	}
	return float64(score) / float64(rounds)
}
