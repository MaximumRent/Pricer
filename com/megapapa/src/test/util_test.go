package test

import (
	"testing"
	"../main/util"
	)

func TestItsNotTrashString_True(t *testing.T) {
	expected := true
	actual := util.IsNonTrashString("Some string")
	if expected != actual {
		t.Fail()
	}
}

func TestItsNotTrashString_False(t *testing.T) {
	expected := false
	actual := util.IsNonTrashString("    \n\t")
	if expected != actual {
		t.Fail()
	}
}

func TestPickoutLettersFromString(t *testing.T) {
	expected := "normalstringwithoutspaces12-3"
	actual := util.PickoutLettersFromString("normal string without spaces 12-3")
	if expected != actual {
		t.Fail()
	}
}
