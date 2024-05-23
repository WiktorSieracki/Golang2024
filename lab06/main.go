package main

import (
	"fmt"
	"math/rand"
)



func plant_trees(tree_amount int, forest []string) []string {
	if tree_amount>SIZE*SIZE {
		fmt.Println(SIZE)
		fmt.Println("TOO MANY TREES!")
	}
	for i := range tree_amount {
		forest[i]=tree
	}
    rand.Shuffle(len(forest), func(i, j int) { forest[i], forest[j] = forest[j], forest[i] })
	return forest
}

func Print_forest(forest []string){
	for i := range SIZE{
		for z := range SIZE{
			fmt.Print(forest[i*SIZE+z])
		}
		fmt.Println()
	}
}
func make_forest() []string {
	forest := make([]string,SIZE*SIZE)
	for i := range forest {
		forest[i] = "  "
	}
	plant_trees(int(float64(SIZE*SIZE)*0.5),forest)
	return forest
}

// func strike_tree(forest []string){
// }

const (
	SIZE = 40
	tree = "🌲"
	fire = "🔥"
	lightning = "⚡"
)
func main() {
	forest := make_forest()
	Print_forest(forest)



	// fmt.Println(forest)


}