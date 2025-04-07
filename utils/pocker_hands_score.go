package utils

import "github.com/Ginstein/TexasHoldemHackathon/model"

const pokerHandsBaseScore = 100000000000

// pokerHandsScore 计算手牌得分，用于最终结果比较
// |0|0000000000| 第一位为手牌类型，后十位为牌型内的大小
func pokerHandsScore(pokerHands model.PokerHands, cards model.Cards) (score int) {
	score = pokerHandsBaseScore*model.PokerHandsValueMap[pokerHands] + pokerHandsScoreFuncMap[pokerHands](cards)
	return
}

var pokerHandsScoreFuncMap = map[model.PokerHands]func(cards model.Cards) int{
	model.RoyalFlush:    royalFlushScore,
	model.StraightFlush: straightFlushScore,
	model.FourOfAKind:   fourOfAKindScore,
	model.FullHouse:     fullHouseScore,
	model.Flush:         flushScore,
	model.Straight:      straightScore,
	model.ThreeOfAKind:  threeOfAKindScore,
	model.TwoPairs:      twoPairsScore,
	model.OnePair:       onePairScore,
	model.HighCard:      highCardScore,
}

// royalFlushScore 皇家同花顺
// 不存在 chop
func royalFlushScore(cards model.Cards) (score int) {
	return
}

// straightFlushScore 同花顺
// 不存在 chop
func straightFlushScore(cards model.Cards) (score int) {
	return
}

// fourOfAKindScore 四条
// 00|00
// 1. 四张
// 2. 一张
func fourOfAKindScore(cards model.Cards) (score int) {
	var counter = cards.Counter()
	for index, count := range counter {
		if count == 4 {
			score += index * 100
		}
		if count == 1 {
			score += index
		}
	}
	return
}

// fullHouseScore 葫芦 3 + 2
// 00|00
// 1. 三张
// 2. 对子
func fullHouseScore(cards model.Cards) (score int) {
	var counter = cards.Counter()
	for index, count := range counter {
		if count == 3 {
			score += index * 100
		}
		if count == 2 {
			score += index
		}
	}
	return
}

// flushScore 同花
func flushScore(cards model.Cards) (score int) {
	// TODO
	return
}

// straightScore 顺子
func straightScore(cards model.Cards) (score int) {
	score = model.CardRanksValueMap[cards[0].Rank]
	return
}

// threeOfAKindScore 三条
// 00|00|00
// 1. 三条
// 2. 单张
// 3. 单张
func threeOfAKindScore(cards model.Cards) (score int) {
	var counter = cards.Counter()
	var highCards []int // 从小到大
	for index, count := range counter {
		if count == 3 {
			score += index * 1000000
		}
		if count == 1 {
			highCards = append(highCards, index)
		}
	}
	score += highCards[0] + highCards[1]*100
	return
}

// twoPairsScore 两对
// 00|00|00
// 1. 一对
// 2. 一对
// 3. 单张
func twoPairsScore(cards model.Cards) (score int) {
	var counter = cards.Counter()
	var onePairs []int
	for index, count := range counter {
		if count == 2 {
			onePairs = append(onePairs, index)
		}
		if count == 1 {
			score = index
		}
	}
	score += onePairs[0]*100 + onePairs[1]*10000
	return
}

// onePairScore 一对
// 00|00|00|00
func onePairScore(cards model.Cards) (score int) {
	var counter = cards.Counter()
	var highCards []int
	for index, count := range counter {
		if count == 2 {
			score += index * 1000000
		}
		if count == 1 {
			highCards = append(highCards, index)
		}
	}
	score += highCards[0] + highCards[1]*100 + highCards[2]*10000
	return
}

// highCardScore 高牌
// 00|00|00|00|00
func highCardScore(cards model.Cards) (score int) {
	for _, card := range cards {
		score = score*100 + model.CardRanksValueMap[card.Rank]
	}
	return
}
