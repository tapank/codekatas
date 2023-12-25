package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	opt := option1
	if option1 > option2 {
		opt = option2
	}
	return opt + " is clearly the better choice."
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	multiplier := 1.0
	if age < 3.0 {
		multiplier = 0.8
	} else if age < 10.0 {
		multiplier = 0.7
	} else {
		multiplier = 0.5
	}
	return originalPrice * multiplier
}
