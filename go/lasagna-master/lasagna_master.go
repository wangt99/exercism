package lasagna

func PreparationTime(layers []string, time int) int {
	if time == 0 {
		time = 2
	}
	return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
	const per_n1 int = 50
	const per_n2 float64 = 0.2
	var n1, n2 int
	for _, layer := range layers {
		if layer == "noodles" {
			n1 += 1
		}
		if layer == "sauce" {
			n2 += 1
		}
	}
	return per_n1 * n1, per_n2 * float64(n2)
}

func AddSecretIngredient(layers1, layers2 []string) {
	if len(layers2) == 0 || len(layers1) == 0 {
		return
	}
	layers2[len(layers2)-1] = layers1[len(layers1)-1]
}

func ScaleRecipe(quantities []float64, num int) []float64 {
	b := make([]float64, len(quantities))

	for i := range quantities {
		b[i] = (quantities[i] / 2) * float64(num)
	}
	return b
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
