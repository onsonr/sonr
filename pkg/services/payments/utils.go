package payments

import "fmt"

func formatAmount(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}
