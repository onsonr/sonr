package rest

import "fmt"

func challengeUuidStoreKey(origin, uuid string)	string {
	return fmt.Sprintf("challenge/%s:%s", origin, uuid)
}

func authorizedUserStoreKey(origin, uuid string) string {
	return fmt.Sprintf("authorized/%s:%s", origin, uuid)
}
