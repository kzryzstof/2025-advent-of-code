package abstractions

import "testing"

func TestProduct_IsValid(t *testing.T) {
	tests := []struct {
		name            string
		productId       string
		expectedIsValid bool
	}{
		/* Initial use cases */
		{
			name:            "Valid product id",
			productId:       "101",
			expectedIsValid: true,
		},
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
		/* New use cases (Part 2) */
		{
			name:            "Invalid product id",
			productId:       "12341234",
			expectedIsValid: false,
		},
		{
			name:            "Invalid product id",
			productId:       "123123123",
			expectedIsValid: false,
		},
		{
			name:            "Invalid product id",
			productId:       "1212121212",
			expectedIsValid: false,
		},
		{
			name:            "Invalid product id",
			productId:       "1111111",
			expectedIsValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, _ := NewProduct(tt.productId)
			actualIsValid := product.IsValid()

			if actualIsValid != tt.expectedIsValid {
				t.Errorf("'%s' | Expected result %t, got %t", tt.productId, tt.expectedIsValid, actualIsValid)
			}
		})
	}
}
