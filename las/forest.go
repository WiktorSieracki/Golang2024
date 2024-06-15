package main

import (
	"fmt"
	"math/rand"
)

func Print_forest(forest []string, size int) {
	for i := range size {
		for z := range size {
			fmt.Print(forest[i*size+z])
		}
		fmt.Println()
	}
    number_of_trees := 0
    number_of_fires := 0
    for _, tree := range forest {
        switch tree {
        case TREE:
            number_of_trees++
        case FIRE, LIGHTNING:
            number_of_fires++
        }
    }
    fmt.Printf("Trees: %d, Fires: %d\n", number_of_trees, number_of_fires)
    percentage := float64(number_of_fires) / float64(number_of_trees) * 100
    fmt.Printf("Percentage of trees burned: %.2f%%\n", percentage)
}
func Make_forest(size int,density float64) []string {
	forest := make([]string, size*size)
	for i := range forest {
		forest[i] = "  "
	}
	plant_trees(int(float64(size*size)*density), forest)
	return forest
}

func plant_trees(tree_amount int, forest []string) []string {
    if tree_amount > *SIZE**SIZE {
		fmt.Println(SIZE)
		fmt.Println("TOO MANY TREES!")
	}
	for i := range tree_amount {
		forest[i]=TREE
	}
    rand.Shuffle(len(forest), func(i, j int) { forest[i], forest[j] = forest[j], forest[i] })
	return forest
}

func Strike_tree(forest []string) []string {
    index := rand.Intn(len(forest))
	for forest[index] != TREE {
		index = rand.Intn(len(forest))
	}
	forest[index] = LIGHTNING
    return forest
}

func Spread_fire(forest []string) []string {
    queue := make([]int, 0)
    for i, tree := range forest {
        if tree == LIGHTNING {
            queue = append(queue, i)
        }
    }
    for len(queue) > 0 {
        i := queue[0]
        queue = queue[1:]
        if i%*SIZE != 0 && forest[i-1] == TREE {
            forest[i-1] = FIRE
            queue = append(queue, i-1)
        }
        if i%*SIZE != *SIZE-1 && forest[i+1] == TREE {
            forest[i+1] = FIRE
            queue = append(queue, i+1)
        }
        if i >= *SIZE && forest[i-*SIZE] == TREE {
            forest[i-*SIZE] = FIRE
            queue = append(queue, i-*SIZE)
        }
        if i < *SIZE*(*SIZE-1) && forest[i+*SIZE] == TREE {
            forest[i+*SIZE] = FIRE
            queue = append(queue, i+*SIZE)
        }
    }
    return forest
}