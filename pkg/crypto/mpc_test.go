package crypto

import (
	"fmt"
	"testing"
)

func Test_MPC(t *testing.T) {
	err := Generate("bio1test")
	fmt.Println(err)
}
