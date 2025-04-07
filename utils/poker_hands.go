package utils

/*
	手牌类型判定;
	1. 皇家同花顺
	2. 同花顺
	3. 四条
	4. 葫芦
	5. 同花
	6. 顺子
	7. 三条
	8. 两对
	9. 一对
	10. 高牌
*/

import (
	"github.com/Ginstein/TexasHoldemHackathon/model"
)

// PokerHandsJudge 手牌类型判定
func PokerHandsJudge(cards model.Cards) (pokerHands model.PokerHands, score int, err error) {
	// default value
	pokerHands = model.HighCard
	if len(cards) != model.PickedCardsCountLimit {
		err = model.PickedCardsCountErr
		return
	}
	if !cards.Check() {
		err = model.CardsCheckErr
		return
	}

	cards.Sort()
	for _, curPokerHands := range model.PokerHandsList {
		if pokerHandsFuncMap[curPokerHands](cards) {
			pokerHands = curPokerHands
			score = pokerHandsScore(curPokerHands, cards)
			return
		}
	}
	return
}

var pokerHandsFuncMap = map[model.PokerHands]func(cards model.Cards) bool{
	model.RoyalFlush:    royalFlush,
	model.StraightFlush: straightFlush,
	model.FourOfAKind:   fourOfAKind,
	model.FullHouse:     fullHouse,
	model.Flush:         flush,
	model.Straight:      straight,
	model.ThreeOfAKind:  threeOfAKind,
	model.TwoPairs:      twoPairs,
	model.OnePair:       onePair,
	model.HighCard:      highCard,
}

// royalFlush 皇家同花顺
func royalFlush(cards model.Cards) (hit bool) {
	return cards[0].Rank == model.Ace && straightFlush(cards)
}

// straightFlush 同花顺
func straightFlush(cards model.Cards) (hit bool) {
	return straight(cards) && flush(cards)
}

// fourOfAKind 四条
func fourOfAKind(cards model.Cards) (hit bool) {
	var counter = cards.Counter()
	var fourCount = 0
	for _, count := range counter {
		if count == 4 {
			fourCount++
		}
	}
	return fourCount == 1
}

// fullHouse 葫芦
func fullHouse(cards model.Cards) (hit bool) {
	var counter = cards.Counter()
	var threeCount = 0
	var pairCount = 0
	for _, count := range counter {
		if count == 3 {
			threeCount++
		}
		if count == 2 {
			pairCount++
		}
	}
	return threeCount == 1 && pairCount == 1
}

// flush 同花
func flush(cards model.Cards) (hit bool) {
	for index := 1; index < len(cards); index++ {
		if cards[index].Suit != cards[0].Suit {
			return false
		}
	}
	return true
}

// straight 顺子
func straight(cards model.Cards) (hit bool) {
	// from large to small
	for index := 1; index < len(cards); index++ {
		var card, preCard = cards[index].Rank, cards[index-1].Rank
		if model.CardRanksValueMap[card] != model.CardRanksValueMap[preCard]-1 {
			return false
		}
	}
	return true
}

// threeOfAKind 三条
func threeOfAKind(cards model.Cards) (hit bool) {
	var counter = cards.Counter()
	var threeCount = 0
	for _, count := range counter {
		if count == 3 {
			threeCount++
		}
	}
	return threeCount == 1
}

// twoPairs 两对
func twoPairs(cards model.Cards) (hit bool) {
	var counter = cards.Counter()
	var pairCount = 0
	for _, count := range counter {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount == 2
}

// onePair 一对
func onePair(cards model.Cards) (hit bool) {
	var counter = cards.Counter()
	var pairCount = 0
	for _, count := range counter {
		if count == 2 {
			pairCount++
		}
	}
	return pairCount == 1
}

// highCard 高牌
func highCard(cards model.Cards) (hit bool) {
	return true
}
