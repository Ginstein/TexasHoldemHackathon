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

// AnalyzePokerHands 手牌类型判定
func AnalyzePokerHands(cards model.Cards) (pokerHands model.PokerHands, err error) {
	if len(cards) != model.PickedCardsCountLimit {
		err = model.PickedCardsCountErr
		return
	}
	if !cards.Check() {
		err = model.CardsCheckErr
		return
	}

	cards.Sort()
	var hit bool
	var strengths []int
	for _, pokerHandsType := range model.PokerHandsTypeList {
		if hit, strengths = pokerHandsTypeJudgeFuncMap[pokerHandsType](cards); hit {
			pokerHands.PokerHandsType = pokerHandsType
			pokerHands.Strengths = strengths
			return
		}
	}
	return
}

var pokerHandsTypeJudgeFuncMap = map[model.PokerHandsType]func(cards model.Cards) (bool, []int){
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
func royalFlush(cards model.Cards) (hit bool, strengths []int) {
	var straightFlushHit bool
	straightFlushHit, strengths = straightFlush(cards)
	hit = straightFlushHit && cards[0].Rank == model.Ace
	return
}

// straightFlush 同花顺
func straightFlush(cards model.Cards) (hit bool, strengths []int) {
	var straightHit, flushHit bool
	straightHit, strengths = straight(cards)
	flushHit, _ = flush(cards)
	hit = straightHit && flushHit
	return
}

// fourOfAKind 四条
func fourOfAKind(cards model.Cards) (hit bool, strengths []int) {
	var counter = cards.Counter()
	var oneCardIndex int
	for index, count := range counter {
		if count == 4 {
			hit = true
			strengths = append(strengths, index)
		}
		if count == 1 {
			oneCardIndex = index
		}
	}
	strengths = append(strengths, oneCardIndex)
	return
}

// fullHouse 葫芦
func fullHouse(cards model.Cards) (hit bool, strengths []int) {
	var counter = cards.Counter()
	var threeCardHit, pairCardHit bool
	var pairCardIndex int
	for index, count := range counter {
		if count == 3 {
			threeCardHit = true
			strengths = append(strengths, index)
		}
		if count == 2 {
			pairCardHit = true
			pairCardIndex = index
		}
	}
	hit = threeCardHit && pairCardHit
	strengths = append(strengths, pairCardIndex)
	return
}

// flush 同花
func flush(cards model.Cards) (hit bool, strengths []int) {
	hit = true
	strengths = append(strengths, model.CardRanksWeightMap[cards[0].Rank])
	for index := 1; index < len(cards); index++ {
		strengths = append(strengths, model.CardRanksWeightMap[cards[index].Rank])
		if cards[index].Suit != cards[0].Suit {
			hit = false
		}
	}
	return
}

// straight 顺子
func straight(cards model.Cards) (hit bool, strengths []int) {
	hit = true
	// from large to small
	for index := 1; index < len(cards); index++ {
		var card, preCard = cards[index].Rank, cards[index-1].Rank
		var cardWeight = model.CardRanksWeightMap[card]
		var preCardWeight = model.CardRanksWeightMap[preCard]
		if cardWeight != preCardWeight-1 {
			hit = false
		}
		if index == 1 {
			strengths = append(strengths, preCardWeight)
		}
	}
	// 特判 A 5 4 3 2
	if cards[0].Rank == model.Ace &&
		cards[1].Rank == model.Five &&
		cards[2].Rank == model.Four &&
		cards[3].Rank == model.Three &&
		cards[4].Rank == model.Two {
		hit = true
		strengths = []int{5}
	}
	return
}

// threeOfAKind 三条
func threeOfAKind(cards model.Cards) (hit bool, strengths []int) {
	var counter = cards.Counter()
	var oneCardIndexs []int
	for index, count := range counter {
		if count == 3 {
			hit = true
			strengths = append(strengths, index)
		}
		if count == 1 {
			oneCardIndexs = append(oneCardIndexs, index)
		}
	}
	for index := len(oneCardIndexs) - 1; index >= 0; index-- {
		strengths = append(strengths, oneCardIndexs[index])
	}
	return
}

// twoPairs 两对
func twoPairs(cards model.Cards) (hit bool, strengths []int) {
	var counter = cards.Counter()
	var pairCount = 0
	var twoPairIndexs []int
	var oneCardIndex int
	for index, count := range counter {
		if count == 2 {
			pairCount++
			twoPairIndexs = append(twoPairIndexs, index)
		}
		if count == 1 {
			oneCardIndex = index
		}
	}
	hit = pairCount == 2
	for index := len(twoPairIndexs) - 1; index >= 0; index-- {
		strengths = append(strengths, twoPairIndexs[index])
	}
	strengths = append(strengths, oneCardIndex)
	return
}

// onePair 一对
func onePair(cards model.Cards) (hit bool, strengths []int) {
	var counter = cards.Counter()
	var oneCardIndexs []int
	for index, count := range counter {
		if count == 2 {
			hit = true
			strengths = append(strengths, index)
		}
		if count == 1 {
			oneCardIndexs = append(oneCardIndexs, index)
		}
	}
	for index := len(oneCardIndexs) - 1; index >= 0; index-- {
		strengths = append(strengths, oneCardIndexs[index])
	}
	return
}

// highCard 高牌
func highCard(cards model.Cards) (hit bool, strengths []int) {
	hit = true
	for _, card := range cards {
		strengths = append(strengths, model.CardRanksWeightMap[card.Rank])
	}
	return
}
