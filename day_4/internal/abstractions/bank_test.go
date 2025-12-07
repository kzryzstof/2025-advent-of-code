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
			expectedVoltage: 987654321111,
		},
		{
			ratings:         "811111111111119",
			expectedVoltage: 811111111119,
		},
		{
			ratings:         "234234234234278",
			expectedVoltage: 434234234278,
		},
		{
			/*
				 	8 181 81911112111
					818 18 1911112111
					818181 911112111
			*/
			ratings:         "818181911112111",
			expectedVoltage: 888911112111,
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
