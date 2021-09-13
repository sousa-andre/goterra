package goterra

import (
	"bytes"
	"encoding/base32"
	"fmt"
)

func encodeBase32(data []byte) string {
	encode := base32.StdEncoding.EncodeToString(data)
	return encode
}

func encodeCardBlock(numOfCards uint8, dst *bytes.Buffer, deck *Deck) {
	tmpDeck := *NewDeck()

	for _, cg := range deck.CardGroups {
		if cg.Count == numOfCards {
			tmpDeck.CardGroups = append(tmpDeck.CardGroups, cg)
		}
	}

	dm := make(map[string]*Deck)
	for _, cg := range tmpDeck.CardGroups {
		groupId := cg.Card.getGroupId()
		if dt, ok := dm[groupId]; ok {
			dt.CardGroups = append(dt.CardGroups, cg)
		} else {
			dm[groupId] = NewDeck()
			dm[groupId].CardGroups = append(dm[groupId].CardGroups, cg)
		}
	}

	dst.WriteByte(byte(len(dm)))
	for _, d := range dm {
		dst.WriteByte(byte(len(d.CardGroups)))
		dst.WriteByte(d.CardGroups[0].Card.Set)
		dst.WriteByte(d.CardGroups[0].Card.Faction)

		for _, c := range d.CardGroups {
			dst.WriteByte(c.Card.Id)
		}
	}
}

func encodeDeckCardsFromBuffer(dest *bytes.Buffer, deck *Deck) {
	for n := 3; n > 0; n-- {
		encodeCardBlock(uint8(n), dest, deck)
	}
}

func (d *Deck) Encode() string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteByte((d.Format << 4) + d.Version)
	encodeDeckCardsFromBuffer(buf, d)

	s := buf.String()
	fmt.Println([]byte(s))

	return encodeBase32([]byte(s))
}
