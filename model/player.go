package model

type Player struct {
	ID         string     // 玩家ID
	HoldCards  Cards      // 手牌
	PokerHands PokerHands // 牌形
}

type Players []Player
