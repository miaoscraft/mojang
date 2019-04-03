package mojang

import (
	"testing"
	"time"
)

func TestGetUUID(t *testing.T) {
	nid, err := GetUUID("Tnze", time.Now())
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

func TestGetUUID_oldName(t *testing.T) {
	nid, err := GetUUID("Denius", time.Date(2016, 10, 10, 1, 0, 0, 0, time.Local))
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
