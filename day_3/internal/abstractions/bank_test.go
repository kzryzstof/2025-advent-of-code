package abstractions

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBank_(t *testing.T) {
	tests := []struct {
		ratings         string
		expectedVoltage uint
	}{
		/* Initial use cases */
		{
			ratings:         "987654321111111",
			expectedVoltage: 98,
		},
		{
			ratings:         "811111111111119",
			expectedVoltage: 89,
		},
		{
			ratings:         "234234234234278",
			expectedVoltage: 78,
		},
		{
			ratings:         "818181911112111",
			expectedVoltage: 92,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s -> %d", tt.ratings, tt.expectedVoltage), func(t *testing.T) {

			batteries := make([]Battery, 0)

			for _, batteryRating := range tt.ratings {
				batteryVoltageRating, _ := strconv.Atoi(string(batteryRating))
				batteries = append(batteries, Battery{Voltage: VoltageRating(batteryVoltageRating), On: false})
			}

			bank := Bank{Batteries: batteries}
			actualVoltage := bank.GetHighestVoltage()

			if actualVoltage != tt.expectedVoltage {
				t.Errorf("'%s' | Expected result %d, got %d", tt.ratings, tt.expectedVoltage, actualVoltage)
			}
		})
	}
}
