package controller

import (
	"fmt"
	"testing"
)

func TestSignUp(t *testing.T) {
	uri := "/web/passage/detail/1"
	body, err := Get(uri)

	fmt.Println(string(body), err)
}
