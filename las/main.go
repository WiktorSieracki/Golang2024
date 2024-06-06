package main

const (
	SIZE      = 40
	DENSITY   = 0.7
	TREE      = "🌲"
	FIRE      = "🔥"
	LIGHTNING = "⚡"
)

func main() {
	forest := Make_forest(SIZE, DENSITY)
	Strike_tree(forest)
	Spread_fire(forest)
	Print_forest(forest, SIZE)

}