package crypto

import (
	"fmt"
	"testing"
)

func Test_MPCCreate(t *testing.T) {
	w, err := Generate(WithParticipants("bio1"))
	fmt.Println(err)
	fmt.Println(w.Config)
}
