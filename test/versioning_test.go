package test

import (
	"github.com/sousa-andre/goterra"
	"testing"
)

func TestInitialVersion(t *testing.T) {
	d := goterra.NewDeck()

	if d.Version != 1 || d.Format != 1 {
		t.Error("The initial version and format should be 1 and 1 respectively")
	}
}

func TestMinimumVersion(t *testing.T) {
	d := goterra.NewDeck()
	d.AddCard("03BW006", 1)

	if d.Version != 2 || d.Format != 1 {
		t.Error("A wrong version is being used for this deck")
	}
}
