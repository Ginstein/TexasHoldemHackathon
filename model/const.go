package model

const (
	MinPlayerCountLimit   = 2
	HoldCardsCountLimit   = 2
	PickedCardsCountLimit = 5
	AllCardsCountLimit    = 7
)

type PokerHands string

const (
	RoyalFlush    PokerHands = "RoyalFlush"    // 皇家同花顺
	StraightFlush PokerHands = "StraightFlush" // 同花顺
	FourOfAKind   PokerHands = "FourOfAKind"   // 四条
	FullHouse     PokerHands = "FullHouse"     // 葫芦
	Flush         PokerHands = "Flush"         // 同花
	Straight      PokerHands = "Straight"      // 顺子
	ThreeOfAKind  PokerHands = "ThreeOfAKind"  // 三条
	TwoPairs      PokerHands = "TwoPairs"      // 两对
	OnePair       PokerHands = "OnePair"       // 一对
	HighCard      PokerHands = "HighCard"      // 高牌
)

var PokerHandsList = []PokerHands{
	RoyalFlush,
	StraightFlush,
	FourOfAKind,
	FullHouse,
	Flush,
	Straight,
	ThreeOfAKind,
	TwoPairs,
	OnePair,
	HighCard,
}

var PokerHandsValueMap = map[PokerHands]int{
	RoyalFlush:    9,
	StraightFlush: 8,
	FourOfAKind:   7,
	FullHouse:     6,
	Flush:         5,
	Straight:      4,
	ThreeOfAKind:  3,
	TwoPairs:      2,
	OnePair:       1,
	HighCard:      0,
}
