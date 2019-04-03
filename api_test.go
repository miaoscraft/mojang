package mojang

import (
	"testing"
)

func TestGetUUID(t *testing.T) {
	nid, err := GetUUID("Tnze", nil)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	want := NameAndUUID{
		UUID: "58f6356eb30c48118bfcd72a9ee99e73",
		Name: "Tnze",
	}
	if nid != want {
		t.Errorf("get %v, want %v", nid, want)
	}
}
