package hdns

import (
	"fmt"
	"net"
)

func LookupSName(sName string) {
	txts, err := net.LookupTXT("gmail.com")
	if err != nil {
		panic(err)
	}
	if len(txts) == 0 {
		fmt.Printf("no record")
	}
	for _, txt := range txts {
		//dig +short gmail.com txt
		fmt.Printf("%s\n", txt)
	}
}
