package test

import (
	"bufio"
	"goterra"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type testDeck struct {
	deckCode   string
	cardGroups []goterra.CardGroup
}

func getDeckCodesTestData(testDecks chan<- testDeck) {
	_, b, _, _ := runtime.Caller(0)
	currDir := filepath.Dir(b)

	file, err := os.Open(filepath.Join(currDir, "testdata/DeckCodesTestData.txt"))
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	go func() {
		currLineIsNewDeck := true
		newDeck := testDeck{}
		for scanner.Scan() {
			if scanner.Text() == "" {
				currLineIsNewDeck = true
				testDecks <- newDeck
				newDeck = testDeck{}
			} else if currLineIsNewDeck {
				currLineIsNewDeck = false
				newDeck.deckCode = scanner.Text()
			} else {
				data := strings.Split(scanner.Text(), ":")
				count, _ := strconv.Atoi(data[0])

				newDeck.cardGroups = append(newDeck.cardGroups, goterra.CardGroup{
					Count: uint8(count),
					Card:  *goterra.CardFromCode(data[1]),
				})
			}
		}
		testDecks <- newDeck
		close(testDecks)
	}()
}

func cardsGroupIncludesCard(cg goterra.CardGroup, cgs []goterra.CardGroup) bool {
	for _, icg := range cgs {
		if cg == icg {
			return true
		}
	}
	return false
}

func cardsMatch(cgs1 []goterra.CardGroup, cgs2 []goterra.CardGroup) bool {
	if len(cgs1) != len(cgs2) {
		return false
	}

	for i := 0; i < len(cgs1); i++ {
		if !cardsGroupIncludesCard(cgs1[i], cgs2) {
			return false
		}
	}
	return true
}
