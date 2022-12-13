package test

import (
	"io/ioutil"
)

func ReadTestFile(name string) []byte {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return data
}
