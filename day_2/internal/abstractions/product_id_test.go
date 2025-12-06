package abstractions

import "testing"

func TestProductId_IsValid(t *testing.T) {
	tests := []struct {
		name            string
		productId       string
		expectedIsValid bool
	}{
		{
			name:            "Valid product id",
			productId:       "10",
			expectedIsValid: true,
		},
		{
			name:            "Invalid product id",
			productId:       "11",
			expectedIsValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productId := ProductId(tt.productId)
			actualIsValid := productId.IsValid()

			if actualIsValid != tt.expectedIsValid {
				t.Errorf("Expected result %d, got %d", tt.expectedIsValid, actualIsValid)
			}
		})
	}
}
