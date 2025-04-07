package utils

import (
	"github.com/Ginstein/TexasHoldemHackathon/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PokerHands", func() {
	Context("cards number is not 5", func() {
		It("should return error", func() {
			_, _, err := PokerHandsJudge([]model.Card{})
			Expect(err).To(Equal(model.PickedCardsCountErr))
		})
	})

	Context("RoyalFlush", func() {
		It("should return RoyalFlush", func() {
			pokerHands, _, err := PokerHandsJudge([]model.Card{
				{Rank: model.Ace, Suit: model.Spade},
				{Rank: model.King, Suit: model.Spade},
				{Rank: model.Queen, Suit: model.Spade},
				{Rank: model.Jack, Suit: model.Spade},
				{Rank: model.Ten, Suit: model.Spade},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands).To(Equal(model.RoyalFlush))
		})

		It("should return RoyalFlush", func() {
			pokerHands, _, err := PokerHandsJudge([]model.Card{
				{Rank: model.Ace, Suit: model.Heart},
				{Rank: model.King, Suit: model.Heart},
				{Rank: model.Queen, Suit: model.Heart},
				{Rank: model.Jack, Suit: model.Heart},
				{Rank: model.Ten, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands).To(Equal(model.RoyalFlush))
		})
	})

	Context("StraightFlush", func() {
		It("should return StraightFlush", func() {
			pokerHands, _, err := PokerHandsJudge([]model.Card{
				{Rank: model.King, Suit: model.Club},
				{Rank: model.Queen, Suit: model.Club},
				{Rank: model.Jack, Suit: model.Club},
				{Rank: model.Ten, Suit: model.Club},
				{Rank: model.Nine, Suit: model.Club},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands).To(Equal(model.StraightFlush))
		})
	})

	Context("FourOfAKind", func() {
		It("should return FourOfAKind", func() {
			pokerHands, _, err := PokerHandsJudge([]model.Card{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Seven, Suit: model.Heart},
				{Rank: model.Seven, Suit: model.Spade},
				{Rank: model.Seven, Suit: model.Diamond},
				{Rank: model.Three, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands).To(Equal(model.FourOfAKind))
		})
	})

	Context("FullHouse", func() {
		It("should return FullHouse", func() {
			pokerHands, _, err := PokerHandsJudge([]model.Card{
				{Rank: model.Five, Suit: model.Club},
				{Rank: model.Five, Suit: model.Heart},
				{Rank: model.Five, Suit: model.Spade},
				{Rank: model.Ace, Suit: model.Diamond},
				{Rank: model.Ace, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands).To(Equal(model.FullHouse))
		})
	})

	Context("Flush", func() {
		It("should return Flush", func() {
			pokerHands, _, err := PokerHandsJudge([]model.Card{
				{Rank: model.Two, Suit: model.Club},
				{Rank: model.Ace, Suit: model.Club},
				{Rank: model.Four, Suit: model.Club},
				{Rank: model.Queen, Suit: model.Club},
				{Rank: model.Six, Suit: model.Club},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands).To(Equal(model.Flush))
		})
	})

	Context("Straight", func() {
		It("should return Straight", func() {
			pokerHands, _, err := PokerHandsJudge([]model.Card{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Eight, Suit: model.Heart},
				{Rank: model.Five, Suit: model.Spade},
				{Rank: model.Nine, Suit: model.Diamond},
				{Rank: model.Six, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands).To(Equal(model.Straight))
		})
	})

	Context("ThreeOfAKind", func() {
		It("should return ThreeOfAKind", func() {
			pokerHands, _, err := PokerHandsJudge([]model.Card{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Seven, Suit: model.Heart},
				{Rank: model.Seven, Suit: model.Spade},
				{Rank: model.Ace, Suit: model.Diamond},
				{Rank: model.Three, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands).To(Equal(model.ThreeOfAKind))
		})
	})

	Context("TwoPairs", func() {
		It("should return TwoPairs", func() {
			pokerHands, _, err := PokerHandsJudge([]model.Card{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Seven, Suit: model.Heart},
				{Rank: model.Ace, Suit: model.Spade},
				{Rank: model.Ace, Suit: model.Diamond},
				{Rank: model.Three, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands).To(Equal(model.TwoPairs))
		})
	})

	Context("OnePair", func() {
		It("should return OnePair", func() {
			pokerHands, _, err := PokerHandsJudge([]model.Card{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Seven, Suit: model.Heart},
				{Rank: model.Queen, Suit: model.Spade},
				{Rank: model.Ace, Suit: model.Diamond},
				{Rank: model.Three, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands).To(Equal(model.OnePair))
		})
	})

	Context("HighCard", func() {
		It("should return HighCard", func() {
			pokerHands, _, err := PokerHandsJudge([]model.Card{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Queen, Suit: model.Heart},
				{Rank: model.Eight, Suit: model.Spade},
				{Rank: model.Ace, Suit: model.Diamond},
				{Rank: model.Three, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands).To(Equal(model.HighCard))
		})
	})
})
