package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type Probability struct {
	Values map[interface{}]int
	Total  int
}

func NewProbability(values ...interface{}) *Probability {
	var prb Probability

	prb.Values = make(map[interface{}]int)

	for _, value := range values {
		prb.Values[value]++
	}

	prb.Total = len(values)

	return &prb
}

func NewProbabilityFromFreqs(values map[interface{}]int) *Probability {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return &Probability{
		Values: values,
		Total:  sum,
	}
}

func (prb *Probability) Sample(size int) []interface{} {
	sample := make([]interface{}, size)

	for i := 0; i < size; i++ {
		r := rand.Intn(prb.Total)

		for value, count := range prb.Values {
			if count > r {
				sample[i] = value
				break
			}

			r -= count
		}
	}

	return sample
}

func (prb *Probability) String() string {
	var data []string

	for value, count := range prb.Values {
		data = append(data, fmt.Sprintf("%v: %d/%d", value, count, prb.Total))
	}

	return strings.Join(data, "\n")
}

func main() {
	coin := NewProbability("Heads", "Tails", "Tails")

	fmt.Println(coin)
	fmt.Println(coin.Sample(10))

	die := NewProbability(1, 2, 3, 4, 5, 6)

	fmt.Println(die)
	fmt.Println(die.Sample(5))

	die1 := NewProbability(1, 1, 5, 5, 9, 9)
	die2 := NewProbability(2, 2, 4, 4, 7, 7)
	die3 := NewProbability(3, 3, 6, 6, 8, 8)

	fmt.Println(die1)
	fmt.Println(die2)
	fmt.Println(die3)
}
