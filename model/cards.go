package model

import (
	"fmt"
	"sort"
)

type CardRank string

const (
	Two   CardRank = "2"
	Three CardRank = "3"
	Four  CardRank = "4"
	Five  CardRank = "5"
	Six   CardRank = "6"
	Seven CardRank = "7"
	Eight CardRank = "8"
	Nine  CardRank = "9"
	Ten   CardRank = "10"
	Jack  CardRank = "J"
	Queen CardRank = "Q"
	King  CardRank = "K"
	Ace   CardRank = "A"
)

var CardRanks = []CardRank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

var CardRanksValueMap = map[CardRank]int{
	Two:   2,
	Three: 3,
	Four:  4,
	Five:  5,
	Six:   6,
	Seven: 7,
	Eight: 8,
	Nine:  9,
	Ten:   10,
	Jack:  11,
	Queen: 12,
	King:  13,
	Ace:   14,
}

type CardSuit string

const (
	Spade   CardSuit = "♠"
	Heart   CardSuit = "♥"
	Club    CardSuit = "♣"
	Diamond CardSuit = "♦"
)

var CardSuits = []CardSuit{Club, Diamond, Heart, Spade}

var CardSuitsValueMap = map[CardSuit]int{
	Spade:   4,
	Heart:   3,
	Club:    2,
	Diamond: 1,
}

type Card struct {
	Rank CardRank
	Suit CardSuit
}

type Cards []Card

func (c Cards) Sort() {
	sort.Slice(c, func(i, j int) bool {
		var cardI, cardJ = c[i], c[j]
		if cardI.Rank != cardJ.Rank {
			return CardRanksValueMap[cardI.Rank] > CardRanksValueMap[cardJ.Rank]
		}
		return CardSuitsValueMap[cardI.Suit] > CardSuitsValueMap[cardJ.Suit]
	})
}

const maxCardsRankValue = 15

// Counter ...
func (c Cards) Counter() (counter [maxCardsRankValue]int) {
	for _, card := range c {
		counter[CardRanksValueMap[card.Rank]]++
	}
	return counter
}

// Check 检查牌型是否合法
func (c Cards) Check() (ok bool) {
	var uniqCardsMap = make(map[string]bool, len(c))
	for _, card := range c {
		// check rank
		if _, ok = CardRanksValueMap[card.Rank]; !ok {
			return false
		}
		// check suit
		if _, ok = CardSuitsValueMap[card.Suit]; !ok {
			return false
		}
		var uniqKey = fmt.Sprintf("%s:%s", card.Rank, card.Suit)
		// 牌面重复
		if _, ok = uniqCardsMap[uniqKey]; ok {
			return false
		}
		uniqCardsMap[uniqKey] = true
	}
	return true
}
