package model

type PokerHandsType string

const (
	RoyalFlush    PokerHandsType = "RoyalFlush"    // 皇家同花顺
	StraightFlush PokerHandsType = "StraightFlush" // 同花顺
	FourOfAKind   PokerHandsType = "FourOfAKind"   // 四条
	FullHouse     PokerHandsType = "FullHouse"     // 葫芦
	Flush         PokerHandsType = "Flush"         // 同花
	Straight      PokerHandsType = "Straight"      // 顺子
	ThreeOfAKind  PokerHandsType = "ThreeOfAKind"  // 三条
	TwoPairs      PokerHandsType = "TwoPairs"      // 两对
	OnePair       PokerHandsType = "OnePair"       // 一对
	HighCard      PokerHandsType = "HighCard"      // 高牌
)

var PokerHandsTypeList = []PokerHandsType{
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

var PokerHandsTypeWeightMap = map[PokerHandsType]int{
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

// PokerHands 手牌分析
type PokerHands struct {
	PokerHandsType PokerHandsType // 手牌类型
	Strengths      []int          // strengths 用于判定双方chop时大小
}

type CompareResult int

const (
	Greater CompareResult = 1
	Equal   CompareResult = 0
	Less    CompareResult = -1
)
