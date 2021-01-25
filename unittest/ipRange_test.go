package unittest

import (
	"testing"

)

func TestSetiprange(t *testing.T) {
	got := GetipList()
	want := "12"
	if got != want {
		t.Errorf("test error")
	}
}
