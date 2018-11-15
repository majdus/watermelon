package main

import (
	"sort"
	"strconv"
)

var maxFlour = 100
var optimalJump = 10

type Result struct {
	Resistance int
	Iterations int
}

func main() {

	var results []Result
	var result Result
	for i := 0; i <= maxFlour; i++  {
		result.Resistance = i
		result.Iterations = getNbIteration(i)
		results = append(results, result)
	}

	//sort by nb iteration
	sort.Slice(results, func(i, j int) bool {
		return results[i].Iterations < results[j].Iterations
	})

	for _, result := range results {
		println("If watermelon resistance equals : " + strconv.Itoa(result.Resistance) + " gives solution in " + strconv.Itoa(result.Iterations) +" iterations");
	}
}

func getNbIteration(resistance int) (int)  {
	nbIteration := 0
	i := 0
	for ; i <= resistance; i = i + optimalJump {
		nbIteration++
	}

	if i == resistance {
		return nbIteration
	} else {
		i = i - optimalJump + 1
		for ; i <= resistance; i++ {
			nbIteration++
		}
	}

	return nbIteration
}