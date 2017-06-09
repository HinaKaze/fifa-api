package api

import "testing"

func TestFUT(t *testing.T) {
	fut := new(Core)
	fut.Init("aaa", "bbbb", "cccc")
	err := fut.Login()
	if err != nil {
		t.Fatal(err)
	}
}
