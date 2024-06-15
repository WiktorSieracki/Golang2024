package main

import (
	"fmt"
	"os"
)

func AverageAmountOfBurnedTrees() [3]float64 {
	var average float64
	for i := 0; i < 1000; i++ {
		forest := Make_forest(*SIZE, *DENSITY)
		Strike_tree(forest)
		Spread_fire(forest)
		burnedTrees := 0
		totalTrees := float64(*SIZE) * float64(*SIZE) * *DENSITY
		for _, tree := range forest {
			if tree == FIRE || tree == LIGHTNING {
				burnedTrees++
			}
		}
		average += float64(burnedTrees) / totalTrees
	}
	average /= 1000
	fmt.Printf("Average amount of burned trees for 1000 forests of size %d and density %.2f: %.2f%%\n", *SIZE, *DENSITY, average*100)
	return [3]float64{float64(*SIZE), *DENSITY, average}

}

func EvaluateDensityVariations() {
    file, err := os.Create("density_results.csv")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()
	_, err = file.WriteString("Size,Density,Percentage_of_trees_burned\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

    for i := 1; i <= 100; i++ {
        *DENSITY = float64(i) / 100
        result := AverageAmountOfBurnedTrees()
        _, err := file.WriteString(fmt.Sprintf("%f,%.2f,%.2f\n", result[0], result[1], result[2]*100))
        if err != nil {
            fmt.Println("Error writing to file:", err)
            return
        }
    }
}