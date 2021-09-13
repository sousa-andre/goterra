package goterra

import (
	"fmt"
	"sort"
	"strconv"
)

type Deck struct {
	Format     uint8
	Version    uint8
	CardGroups []CardGroup
}

func NewDeck() *Deck {
	d := Deck{
		Version: uint8(1),
		Format:  uint8(1),
	}
	return &d
}

func (d *Deck) AddCard(cardCode string, cardCount uint8) {
	if cardCount > 3 || cardCount < 1 {
		panic("The argument cardCount must be within the values 1 and 3")
	}

	c := *CardFromCode(cardCode)
	if v := c.GetCardMinVersion(); v > d.Version {
		d.Version = v
	}
	if f := c.GetCardMinFormat(); f > d.Format {
		d.Format = f
	}

	d.CardGroups = append(d.CardGroups, CardGroup{
		Count: cardCount,
		Card:  c,
	})
}
func (d *Deck) SortDeck() {
	sort.Slice(d.CardGroups, func(i, j int) bool {
		return d.CardGroups[i].Card.GetCardCode() > d.CardGroups[j].Card.GetCardCode()
	})
}

func CardFromCode(cardCode string) *Card {
	set, _ := strconv.Atoi(cardCode[:2])
	faction := idFromFactionString(cardCode[2:4])
	id, _ := strconv.Atoi(cardCode[4:])

	return &Card{
		Set:     uint8(set),
		Faction: uint8(faction),
		Id:      uint8(id),
	}
}

func (c Card) GetFactionStringId() string {
	return Factions[int(c.Faction)]
}

func (c Card) GetCardCode() string {
	return fmt.Sprintf("%02d%s%03d", c.Set, c.GetFactionStringId(), c.Id)
}

func (c Card) getGroupId() string {
	return fmt.Sprintf("%d.%d", c.Set, c.Faction)
}

func (c Card) GetCardMinVersion() uint8 {
	return FactionsVersion[c.Faction][1]
}

func (c Card) GetCardMinFormat() uint8 {
	return FactionsVersion[c.Faction][0]
}

type CardGroup struct {
	Count uint8
	Card  Card
}

type Card struct {
	Set     uint8
	Faction uint8
	Id      uint8
}

type Faction struct {
	Id     uint8
	TextId string
}
