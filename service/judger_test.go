package service_test

import (
	"github.com/Ginstein/TexasHoldemHackathon/model"
	"github.com/Ginstein/TexasHoldemHackathon/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	// players
	player1 = &model.Player{
		ID:        "001",
		HoldCards: model.Cards{{model.Ace, model.Diamond}, {model.King, model.Heart}},
	}
	player2 = &model.Player{
		ID:        "002",
		HoldCards: model.Cards{{model.Two, model.Spade}, {model.Two, model.Heart}},
	}
	player3 = &model.Player{
		ID:        "003",
		HoldCards: model.Cards{{model.Four, model.Club}, {model.Three, model.Club}},
	}
	player4 = &model.Player{
		ID:        "004",
		HoldCards: model.Cards{{model.Six, model.Club}, {model.Three, model.Heart}},
	}
	// public cards
	// player 1 顺子 player 2 两对 player 3 三条
	publicCards1 = model.Cards{{model.Queen, model.Diamond}, {model.Jack, model.Heart},
		{model.Three, model.Spade}, {model.Ten, model.Spade}, {model.Three, model.Diamond}}
)

var _ = Describe("Judger", func() {
	Context("Judger Error", func() {
		It("MinPlayerCountErr", func() {
			var players = model.Players{player1}
			var publicCards = publicCards1
			var _, err = service.NewTexasHoldemJudger(players, publicCards)
			Expect(err).To(Equal(model.MinPlayerCountErr))
		})

		It("HoldCardsCountErr", func() {
			var playerErr = &model.Player{
				ID:        "error",
				HoldCards: model.Cards{{model.Ace, model.Diamond}},
			}
			var players = model.Players{player1, playerErr}
			var publicCards = publicCards1
			var _, err = service.NewTexasHoldemJudger(players, publicCards)
			Expect(err).To(Equal(model.HoldCardsCountErr))
		})

		It("PublicCardsCountErr", func() {
			var players = model.Players{player1, player2, player3}
			var publicCardsErr = model.Cards{{model.Ace, model.Diamond}, {model.King, model.Heart}}
			var _, err = service.NewTexasHoldemJudger(players, publicCardsErr)
			Expect(err).To(Equal(model.PublicCardsCountErr))
		})

		It("CardsCheckErr", func() {
			var playerCardsErr = &model.Player{
				ID:        "error",
				HoldCards: model.Cards{{model.Queen, model.Diamond}, {model.King, model.Club}},
			}
			var players = model.Players{player1, player2, player3, playerCardsErr}
			var publicCards = publicCards1
			var _, err = service.NewTexasHoldemJudger(players, publicCards)
			Expect(err).To(Equal(model.CardsCheckErr))
		})
	})

	Context("Judger Success", func() {
		It("Judge test case - 001", func() {
			var players = model.Players{player1, player2, player3}
			var publicCards = publicCards1
			var judger, err = service.NewTexasHoldemJudger(players, publicCards)
			Expect(err).To(BeNil())
			var winners, _ = judger.Judge()
			Expect(len(winners)).To(Equal(1))
			Expect(winners[0]).To(Equal(player1.ID))
			Expect(players[0].PokerHands.PokerHandsType).To(Equal(model.Straight))
			Expect(players[1].PokerHands.PokerHandsType).To(Equal(model.ThreeOfAKind))
			Expect(players[2].PokerHands.PokerHandsType).To(Equal(model.TwoPairs))
		})

		It("Judge test case - 002", func() {
			var players = model.Players{player1, player2}
			var publicCards = publicCards1
			var judger, err = service.NewTexasHoldemJudger(players, publicCards)
			Expect(err).To(BeNil())
			var winners, _ = judger.Judge()
			Expect(len(winners)).To(Equal(1))
			Expect(winners[0]).To(Equal(player1.ID))
			Expect(players[0].PokerHands.PokerHandsType).To(Equal(model.Straight))
			Expect(players[1].PokerHands.PokerHandsType).To(Equal(model.TwoPairs))
		})

		It("Judge test case - 003", func() {
			var players = model.Players{player1, player3}
			var publicCards = publicCards1
			var judger, err = service.NewTexasHoldemJudger(players, publicCards)
			Expect(err).To(BeNil())
			var winners, _ = judger.Judge()
			Expect(len(winners)).To(Equal(1))
			Expect(winners[0]).To(Equal(player1.ID))
			Expect(players[0].PokerHands.PokerHandsType).To(Equal(model.Straight))
			Expect(players[1].PokerHands.PokerHandsType).To(Equal(model.ThreeOfAKind))
		})

		It("Judge test case - 004", func() {
			var players = model.Players{player2, player3}
			var publicCards = publicCards1
			var judger, err = service.NewTexasHoldemJudger(players, publicCards)
			Expect(err).To(BeNil())
			var winners, _ = judger.Judge()
			Expect(len(winners)).To(Equal(1))
			Expect(winners[0]).To(Equal(player3.ID))
			Expect(players[0].PokerHands.PokerHandsType).To(Equal(model.ThreeOfAKind))
			Expect(players[1].PokerHands.PokerHandsType).To(Equal(model.TwoPairs))
		})

		It("Judge test case - 005 chop", func() {
			var players = model.Players{player3, player4}
			var publicCards = publicCards1
			var judger, err = service.NewTexasHoldemJudger(players, publicCards)
			Expect(err).To(BeNil())
			var winners, _ = judger.Judge()
			Expect(len(winners)).To(Equal(2))
			Expect(players[0].PokerHands.PokerHandsType).To(Equal(model.ThreeOfAKind))
			Expect(players[1].PokerHands.PokerHandsType).To(Equal(model.ThreeOfAKind))
		})
	})
})
