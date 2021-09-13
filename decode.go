package goterra

import (
	"bytes"
	"encoding/base32"
	"encoding/binary"
	"log"
	"strings"
)

func decodeBase32(data string) *bytes.Buffer {
	data = strings.Trim(data, "=")
	p := len(data) % 8

	for i := 0; i < 8-p && p != 0; i++ {
		data += "="
	}

	decode, err := base32.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return bytes.NewBuffer(decode)
}

func decodeCardsBlock(numOfCards uint8, data *bytes.Buffer, deck *Deck) {
	numGroupOfs, _ := binary.ReadUvarint(data)

	for c := 0; c < int(numGroupOfs); c++ {
		numOfsInThisGroup, _ := binary.ReadUvarint(data)
		set, _ := binary.ReadUvarint(data)
		faction, _ := binary.ReadUvarint(data)

		for i := 0; i < int(numOfsInThisGroup); i++ {
			cardId, _ := binary.ReadUvarint(data)

			deck.CardGroups = append(deck.CardGroups, CardGroup{
				Count: numOfCards,
				Card: Card{
					Set:     uint8(set),
					Faction: uint8(faction),
					Id:      uint8(cardId),
				},
			})
		}
	}
}

func decodeDeckCardsFromBuffer(data *bytes.Buffer, deck *Deck) {
	var n uint8
	for n = 3; n > 0; n-- {
		decodeCardsBlock(n, data, deck)
	}
}

func DeckFromCode(code string) *Deck {
	buf := decodeBase32(code)
	formatAndVersionByte, err := buf.ReadByte()
	if err != nil {
		log.Fatalln("Error decoding the deck")
	}
	format := formatAndVersionByte >> 4
	version := formatAndVersionByte & 0xF

	deck := NewDeck()
	deck.Format = format
	deck.Version = version
	decodeDeckCardsFromBuffer(buf, deck)
	return deck
}
