package abstractions

type batteryIndex = uint

type Bank struct {
	Batteries []Battery
}

func (b *Bank) GetHighestVoltage() uint {

	/* Find the 1st battery with the highest voltage (the last one must be excluded) */
	highestVoltage, highestVoltageIndex := b.getHighestVoltage(0, batteryIndex(len(b.Batteries)-1))

	/* Finds the biggest voltage for the 2nd battery located after the 1st battery, including the last battery */
	secondHighestVoltage, _ := b.getHighestVoltage(highestVoltageIndex+1, batteryIndex(len(b.Batteries)))

	return uint(highestVoltage*10 + secondHighestVoltage)
}

func (b *Bank) getHighestVoltage(fromIndex batteryIndex, toIndex batteryIndex) (VoltageRating, batteryIndex) {

	highestVoltageIndex := batteryIndex(0)
	highestVoltage := VoltageRating(0)

	for index, battery := range b.Batteries[fromIndex:toIndex] {
		if battery.Voltage > highestVoltage {
			highestVoltage = battery.Voltage
			/* The returned index is relative to the subslide */
			absoluteIndex := fromIndex + batteryIndex(index)
			highestVoltageIndex = absoluteIndex
		}
	}

	return highestVoltage, highestVoltageIndex
}
