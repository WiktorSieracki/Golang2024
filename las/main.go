package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var (
    SIZE      = flag.Int("s", 50, "Size of the forest")
    DENSITY   = flag.Float64("d", 0.5, "Density of the forest")
	recursive = flag.Bool("r", false, "Run recursively")
	calculate_average = flag.Bool("c", false, "Calculate average amount of burned trees for 1000 forests with different densities")
    TREE      = "🌲"
    FIRE      = "🔥"
    LIGHTNING = "⚡"
)

func main() {
    flag.Parse()

	for {
		forest := Make_forest(*SIZE, *DENSITY)
		Strike_tree(forest)
		Spread_fire(forest)
		Print_forest(forest, *SIZE)
		
		if !*recursive {
			break
		}
		
		fmt.Println("Press 'Enter' to generate another forest or Ctrl+C to exit.")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
	if *calculate_average {
		EvaluateDensityVariations()
	}
}