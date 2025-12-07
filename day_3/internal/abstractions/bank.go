package abstractions

import "math"

type batteryIndex = uint

type Bank struct {
	Batteries []Battery
}

func (b *Bank) GetHighestVoltage() uint {

	expectedDigits := 12
	totalVoltage := uint(0)

	currentBatteryIndex := batteryIndex(0)
	batteryCount := batteryIndex(len(b.Batteries))

	for currentDigit := 0; currentDigit < expectedDigits; currentDigit++ {

		/* At the current digit, the minimum remaining digits are 1 less than expected */
		minimumRemainingDigits := expectedDigits - currentDigit - 1

		/* Figure out the range to work with for the current digit */
		fromIndex := currentBatteryIndex
		toIndex := batteryCount - batteryIndex(minimumRemainingDigits)

		highestVoltage, highestVoltageIndex := b.getHighestVoltage(fromIndex, toIndex)

		/* Adds the current digit to the total voltage with the appropriate power of 10 */
		totalVoltage += uint(highestVoltage) * uint(math.Pow10(expectedDigits-currentDigit-1))

		/* Moves the current battery index to the next one */
		currentBatteryIndex = highestVoltageIndex + batteryIndex(1)
	}

	return totalVoltage
}

func (b *Bank) getHighestVoltage(fromIndex batteryIndex, remainingIndex batteryIndex) (VoltageRating, batteryIndex) {

	highestVoltageIndex := batteryIndex(0)
	highestVoltage := VoltageRating(0)

	for index, battery := range b.Batteries[fromIndex:remainingIndex] {
		if battery.Voltage > highestVoltage {
			highestVoltage = battery.Voltage
			/* The returned index is relative to the subslide */
			absoluteIndex := fromIndex + batteryIndex(index)
			highestVoltageIndex = absoluteIndex
		}
	}

	return highestVoltage, highestVoltageIndex
}
