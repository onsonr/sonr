package router

import (
	"fmt"
)

func endpointAuth(b string) string {
	return fmt.Sprintf("%s/highway/auth", b)
}

func endpointDB(b string) string {
	return fmt.Sprintf("%s/highway/db", b)
}

func endpointStore(b string) string {
	return fmt.Sprintf("%s/highway/store", b)
}

func endpointWallet(b string) string {
	return fmt.Sprintf("%s/highway/wallet", b)
}
