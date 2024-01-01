package lasagna

// PreparationTime returns the total time to prepare all layers.
func PreparationTime(layers []string, prepTime int) int {
	if prepTime <= 0 {
		prepTime = 2
	}
	return len(layers) * prepTime
}

// Quantities returns the required quantity of noodles and sauce.
func Quantities(layers []string) (nood int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			nood += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

// AddSecretIngredient() substitutes the last "?" in my list with
// the last item in friendsList.
func AddSecretIngredient(friendsList, myList []string) {
	if len(friendsList) > 0 && len(myList) > 0 && myList[len(myList)-1] == "?" {
		myList[len(myList)-1] = friendsList[len(friendsList)-1]
	}
}

// ScaleRecipe creates a new list of quantities for the given portion for 2.
func ScaleRecipe(quantities []float64, portions int) []float64 {
	if portions < 1 {
		return nil
	}

	scaled := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		scaled[i] = quantities[i] * float64(portions) / 2
	}
	return scaled
}
