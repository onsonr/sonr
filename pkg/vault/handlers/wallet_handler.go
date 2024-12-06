package handlers

//
// // Constants for supported payment methods
// const (
// 	MethodCard       = "basic-card"
// 	MethodGooglePay  = "https://google.com/pay"
// 	MethodApplePay   = "https://apple.com/apple-pay"
// 	MethodSonrWallet = "https://sonr.id/wallet"
// )
//
// // InitiatePayment starts the payment request flow
// func InitiatePayment(c echo.Context, request PaymentRequest) error {
// 	return StartPayment(request).Render(c.Request().Context(), c.Response().Writer)
// }
//
// // Helper functions to create payment requests
// func NewBasicCardPayment(amount float64, currency string, label string) PaymentRequest {
// 	return PaymentRequest{
// 		MethodData: []PaymentMethodData{
// 			{
// 				SupportedMethods: MethodCard,
// 				Data: map[string]any{
// 					"supportedNetworks": []string{"visa", "mastercard"},
// 					"supportedTypes":    []string{"credit", "debit"},
// 				},
// 			},
// 		},
// 		Details: PaymentDetails{
// 			Total: PaymentItem{
// 				Label: label,
// 				Amount: Money{
// 					Currency: currency,
// 					Value:    formatAmount(amount),
// 				},
// 			},
// 		},
// 		Options: PaymentOptions{
// 			RequestPayerName:  true,
// 			RequestPayerEmail: true,
// 		},
// 	}
// }
//
// // Example usage:
// func PaymentHandler(c echo.Context) error {
// 	request := NewBasicCardPayment(99.99, "USD", "Product Purchase")
//
// 	// Add display items
// 	request.Details.DisplayItems = []PaymentItem{
// 		{
// 			Label: "Product Price",
// 			Amount: Money{
// 				Currency: "USD",
// 				Value:    "89.99",
// 			},
// 		},
// 		{
// 			Label: "Tax",
// 			Amount: Money{
// 				Currency: "USD",
// 				Value:    "10.00",
// 			},
// 		},
// 	}
//
// 	// Add shipping options
// 	request.Details.ShippingOptions = []ShippingOption{
// 		{
// 			ID:    "standard",
// 			Label: "Standard Shipping",
// 			Amount: Money{
// 				Currency: "USD",
// 				Value:    "0.00",
// 			},
// 			Selected: true,
// 		},
// 		{
// 			ID:    "express",
// 			Label: "Express Shipping",
// 			Amount: Money{
// 				Currency: "USD",
// 				Value:    "10.00",
// 			},
// 		},
// 	}
//
// 	return InitiatePayment(c, request)
// }
//
// func formatAmount(amount float64) string {
// 	return fmt.Sprintf("%.2f", amount)
// }
