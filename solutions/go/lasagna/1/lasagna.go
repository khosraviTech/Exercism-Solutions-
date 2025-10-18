package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime =40
// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    remain:= OvenTime - actualMinutesInOven
	return remain
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
    prepareTime := numberOfLayers * 2
	return prepareTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    El := PreparationTime(numberOfLayers) + actualMinutesInOven
	return El
}
