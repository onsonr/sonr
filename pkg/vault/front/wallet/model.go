package wallet

// Payment types
type PaymentMethodData struct {
	SupportedMethods string         `json:"supportedMethods"`
	Data             map[string]any `json:"data,omitempty"`
}

type PaymentItem struct {
	Label  string `json:"label"`
	Amount Money  `json:"amount"`
}

type Money struct {
	Currency string `json:"currency"`
	Value    string `json:"value"` // Decimal as string for precision
}

type PaymentOptions struct {
	RequestPayerName  bool `json:"requestPayerName,omitempty"`
	RequestPayerEmail bool `json:"requestPayerEmail,omitempty"`
	RequestPayerPhone bool `json:"requestPayerPhone,omitempty"`
	RequestShipping   bool `json:"requestShipping,omitempty"`
}

type PaymentDetails struct {
	Total           PaymentItem      `json:"total"`
	DisplayItems    []PaymentItem    `json:"displayItems,omitempty"`
	ShippingOptions []ShippingOption `json:"shippingOptions,omitempty"`
}

type ShippingOption struct {
	ID       string `json:"id"`
	Label    string `json:"label"`
	Amount   Money  `json:"amount"`
	Selected bool   `json:"selected,omitempty"`
}

type PaymentRequest struct {
	MethodData []PaymentMethodData `json:"methodData"`
	Details    PaymentDetails      `json:"details"`
	Options    PaymentOptions      `json:"options,omitempty"`
}
