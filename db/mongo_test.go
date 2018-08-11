package db

import "testing"

func TestSave(t *testing.T) {
	err := Save("the world is my country", "MazroE")
	if err != nil {
		t.Error("faild to save qoute")
	}

	if len(myqoute) != 1 {
		t.Error("nothing is saved")
	}
}

func TestAll(t *testing.T) {
	qs, err := All()
	if err != nil {
		t.Error("faild to query all qoutes")
	}

	if len(qs) != 1 {
		t.Error("no qoute")
	}

	if qs[0].Text != "the world is my country" {
		t.Error("wrong value")
	}
}
