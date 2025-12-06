package abstractions

import "testing"

func TestProductId_IsValid(t *testing.T) {
	tests := []struct {
		name            string
		productId       string
		expectedIsValid bool
	}{
		/* Initial use cases */
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
		{
			name:            "Valid product id",
			productId:       "1011",
			expectedIsValid: true,
		},
		{
			name:            "Invalid product id",
			productId:       "1010",
			expectedIsValid: false,
		},
		{
			name:            "Valid product id",
			productId:       "1188511884",
			expectedIsValid: true,
		},
		{
			name:            "Invalid product id",
			productId:       "1188511885",
			expectedIsValid: false,
		},
		{
			name:            "Valid product id",
			productId:       "222223",
			expectedIsValid: true,
		},
		{
			name:            "Invalid product id",
			productId:       "222222",
			expectedIsValid: false,
		},
		{
			name:            "Valid product id",
			productId:       "445446",
			expectedIsValid: true,
		},
		{
			name:            "Invalid product id",
			productId:       "446446",
			expectedIsValid: false,
		},
		{
			name:            "Valid product id",
			productId:       "38593858",
			expectedIsValid: true,
		},
		{
			name:            "Invalid product id",
			productId:       "38593859",
			expectedIsValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			productId := ProductId(tt.productId)
			actualIsValid := productId.IsValid()

			if actualIsValid != tt.expectedIsValid {
				t.Errorf("Expected result %t, got %t", tt.expectedIsValid, actualIsValid)
			}
		})
	}
}
