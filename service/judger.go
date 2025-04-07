package service

import (
	"github.com/Ginstein/TexasHoldemHackathon/model"
	"github.com/Ginstein/TexasHoldemHackathon/utils"
	"sort"
)

var masks []int

func init() {
	// 生成所有可能的 7 选 5 的掩码
	var maxMask = (1 << model.AllCardsCountLimit) - 1 // 111111
	for mask := 1; mask < maxMask; mask++ {
		var numOfBit = 0
		for index := 0; index < model.AllCardsCountLimit; index++ {
			if mask&(1<<index) != 0 {
				numOfBit++
			}
		}
		if numOfBit == model.PickedCardsCountLimit {
			masks = append(masks, mask)
		}
	}
}

// TexasHoldemJudger 德州扑克判定器
type TexasHoldemJudger struct {
	players     model.Players // 玩家列表
	publicCards model.Cards
	winners     []string // 可能有平局
}

// NewTexasHoldemJudger 创建一个新的德州扑克判定器
// players 玩家列表
// publicCards 公共牌
func NewTexasHoldemJudger(players model.Players, publicCards model.Cards) (judge *TexasHoldemJudger, err error) {
	if len(players) < model.MinPlayerCountLimit {
		err = model.MinPlayerCountErr
		return
	}
	var cards = publicCards
	// check players
	for _, player := range players {
		cards = append(cards, player.HoldCards...)
		if len(player.HoldCards) != model.HoldCardsCountLimit {
			err = model.HoldCardsCountErr
			return
		}
	}
	// check public cards
	if len(publicCards) != model.PickedCardsCountLimit {
		err = model.PublicCardsCountErr
		return
	}
	// check cards
	if !cards.Check() {
		err = model.CardsCheckErr
		return
	}
	judge = &TexasHoldemJudger{
		players:     players,
		publicCards: publicCards,
	}
	return
}

// Judge 判定牌型
func (p *TexasHoldemJudger) Judge() (err error) {
	for _, player := range p.players {
		var cards = append(p.publicCards, player.HoldCards...)
		player.PokerHands, player.Score, err = judgePlayer(cards)
	}
	sort.Slice(p.players, func(i, j int) bool {
		var playerI, playerJ = p.players[i], p.players[j]
		return playerI.Score > playerJ.Score
	})
	p.winners = []string{p.players[0].ID}
	// 多名获胜者
	var winnerScore = p.players[0].Score
	for index := 1; index < len(p.players); index++ {
		if p.players[index].Score != winnerScore {
			break
		}
		p.winners = append(p.winners, p.players[index].ID)
	}
	return
}

// judgePlayer 判定玩家牌型
func judgePlayer(cards model.Cards) (pokerHands model.PokerHands, score int, err error) {
	// Pick five cards out of seven
	for _, mask := range masks {
		var pickedCards model.Cards
		for index, card := range cards {
			if mask&(1<<index) != 0 {
				pickedCards = append(pickedCards, card)
			}
		}
		var curPokerHands model.PokerHands
		var curScore int
		if curPokerHands, curScore, err = utils.PokerHandsJudge(pickedCards); err != nil {
			return
		}
		if curScore > score {
			pokerHands = curPokerHands
			score = curScore
		}
	}
	return
}
