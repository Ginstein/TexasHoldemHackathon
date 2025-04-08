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
func NewTexasHoldemJudger(players model.Players, publicCards model.Cards) (judger *TexasHoldemJudger, err error) {
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
	judger = &TexasHoldemJudger{
		players:     players,
		publicCards: publicCards,
	}
	return
}

// Judge 判定牌型
func (p *TexasHoldemJudger) Judge() (winners []string, err error) {
	for _, player := range p.players {
		var cards = append(p.publicCards, player.HoldCards...)
		player.PokerHands, err = playerBestPokerHands(cards)
	}
	sort.Slice(p.players, func(i, j int) bool {
		var playerI, playerJ = p.players[i], p.players[j]
		return utils.PokerHandsCompare(playerI.PokerHands, playerJ.PokerHands) == model.Greater
	})
	p.winners = []string{p.players[0].ID}
	// 多名获胜者
	for index := 1; index < len(p.players); index++ {
		if utils.PokerHandsCompare(p.players[0].PokerHands, p.players[index].PokerHands) != model.Equal {
			break
		}
		p.winners = append(p.winners, p.players[index].ID)
	}
	winners = p.winners
	return
}

// playerBestPokerHands 选取玩家最优牌形
func playerBestPokerHands(cards model.Cards) (pokerHands model.PokerHands, err error) {
	// init
	pokerHands.PokerHandsType = model.HighCard
	pokerHands.Strengths = []int{-1}
	// Pick five cards out of seven
	for _, mask := range masks {
		var pickedCards model.Cards
		for index, card := range cards {
			if mask&(1<<index) != 0 {
				pickedCards = append(pickedCards, card)
			}
		}
		var curPokerHands model.PokerHands
		if curPokerHands, err = utils.AnalyzePokerHands(pickedCards); err != nil {
			return
		}
		// 选取最优解
		if utils.PokerHandsCompare(curPokerHands, pokerHands) == model.Greater {
			pokerHands = curPokerHands
		}
	}
	return
}
