package main

import (
	"fmt"
	"math/rand"
	"time"
)



func plant_trees(tree_amount int, forest []string) []string {
	if tree_amount>SIZE*SIZE {
		fmt.Println(SIZE)
		fmt.Println("TOO MANY TREES!")
	}
	for i := range tree_amount {
		forest[i]=tree
	}
	rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(forest), func(i, j int) { forest[i], forest[j] = forest[j], forest[i] })
	return forest
}

func strike_tree(){
	
}


var SIZE = 10
var tree = "🌲"
var fire = "🔥"
var lightning = "⚡"
func main() {
	
	forest := make([]string,SIZE*SIZE)

	for i := range forest {
        forest[i] = "  "
    }

	plant_trees(50,forest)

	for i := range SIZE{
		fmt.Println()
		for z := range SIZE{
			fmt.Print(forest[i*10+z])
		}
	}
	fmt.Println()


	// fmt.Println(forest)


}