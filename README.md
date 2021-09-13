# goterra

goterra is a package for Go lang that makes encoding/decoding **Legends of Runterra** decks easy. This library gets updated according [this](https://github.com/RiotGames/LoRDeckCodes) repository.


##### If you have any questions about how to use this library or related to LOR itself feel free to join [Riot API Dev Community](https://discord.gg/riotgamesdevrel) Discord server or send a private message to André#5360


## Download
- `$ go get github.com/sousa-andre/goterra`

## Usage

```go
package main

import (
	"fmt"
	"github.com/sousa-andre/goterra"
)

func main() {
	// Import deck
	id := goterra.DeckFromCode("CEBAIAIFB4WDANQIAEAQGDAUDAQSIJZUAIAQCAIEAEAQKBIA")

	for _, cg := range id.CardGroups {
		fmt.Printf("Card Code: %s. Card Count %d\n", cg.Card.GetCardCode(), cg.Count)
	}

	// Create deck
	cd := goterra.NewDeck()
	cd.AddCard("01PZ019", 2)
	cd.AddCard("01IO018", 1)

	fmt.Println("Deck Code: ", cd.Encode())
}
```

## Endorsement
goterra isn’t endorsed by Riot Games and doesn’t reflect the views or opinions of Riot Games or anyone officially involved in producing or managing League of Legends. League of Legends and Riot Games are trademarks or registered trademarks of Riot Games, Inc. League of Legends © Riot Games, Inc.