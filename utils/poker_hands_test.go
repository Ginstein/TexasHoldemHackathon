package utils

import (
	"github.com/Ginstein/TexasHoldemHackathon/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PokerHands", func() {
	Context("cards number is not 5", func() {
		It("should return error", func() {
			_, err := AnalyzePokerHands(model.Cards{})
			Expect(err).To(Equal(model.PickedCardsCountErr))
		})
	})

	Context("RoyalFlush", func() {
		It("should return RoyalFlush", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.Ace, Suit: model.Spade},
				{Rank: model.King, Suit: model.Spade},
				{Rank: model.Queen, Suit: model.Spade},
				{Rank: model.Jack, Suit: model.Spade},
				{Rank: model.Ten, Suit: model.Spade},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.RoyalFlush))
		})

		It("should return RoyalFlush", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.Ace, Suit: model.Heart},
				{Rank: model.King, Suit: model.Heart},
				{Rank: model.Queen, Suit: model.Heart},
				{Rank: model.Jack, Suit: model.Heart},
				{Rank: model.Ten, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.RoyalFlush))
			Expect(pokerHands.Strengths).To(Equal([]int{14}))
		})
	})

	Context("StraightFlush", func() {
		It("should return StraightFlush", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.King, Suit: model.Club},
				{Rank: model.Queen, Suit: model.Club},
				{Rank: model.Jack, Suit: model.Club},
				{Rank: model.Ten, Suit: model.Club},
				{Rank: model.Nine, Suit: model.Club},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.StraightFlush))
			Expect(pokerHands.Strengths).To(Equal([]int{13}))
		})
	})

	Context("FourOfAKind", func() {
		It("should return FourOfAKind", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Seven, Suit: model.Heart},
				{Rank: model.Seven, Suit: model.Spade},
				{Rank: model.Seven, Suit: model.Diamond},
				{Rank: model.Jack, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.FourOfAKind))
			Expect(pokerHands.Strengths).To(Equal([]int{7, 11}))
		})
	})

	Context("FullHouse", func() {
		It("should return FullHouse", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.Five, Suit: model.Club},
				{Rank: model.Five, Suit: model.Heart},
				{Rank: model.Five, Suit: model.Spade},
				{Rank: model.King, Suit: model.Diamond},
				{Rank: model.King, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.FullHouse))
			Expect(pokerHands.Strengths).To(Equal([]int{5, 13}))
		})
	})

	Context("Flush", func() {
		It("should return Flush", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.Two, Suit: model.Club},
				{Rank: model.Ace, Suit: model.Club},
				{Rank: model.Four, Suit: model.Club},
				{Rank: model.Queen, Suit: model.Club},
				{Rank: model.Six, Suit: model.Club},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.Flush))
			Expect(pokerHands.Strengths).To(Equal([]int{14, 12, 6, 4, 2}))
		})
	})

	Context("Straight", func() {
		It("should return Straight", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Eight, Suit: model.Heart},
				{Rank: model.Five, Suit: model.Spade},
				{Rank: model.Nine, Suit: model.Diamond},
				{Rank: model.Six, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.Straight))
			Expect(pokerHands.Strengths).To(Equal([]int{9}))
		})

		It("should return Straight", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.Ace, Suit: model.Club},
				{Rank: model.Five, Suit: model.Heart},
				{Rank: model.Two, Suit: model.Spade},
				{Rank: model.Four, Suit: model.Diamond},
				{Rank: model.Three, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.Straight))
			Expect(pokerHands.Strengths).To(Equal([]int{5}))
		})
	})

	Context("ThreeOfAKind", func() {
		It("should return ThreeOfAKind", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Seven, Suit: model.Heart},
				{Rank: model.Seven, Suit: model.Spade},
				{Rank: model.Ace, Suit: model.Diamond},
				{Rank: model.Six, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.ThreeOfAKind))
			Expect(pokerHands.Strengths).To(Equal([]int{7, 14, 6}))
		})
	})

	Context("TwoPairs", func() {
		It("should return TwoPairs", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Seven, Suit: model.Heart},
				{Rank: model.Ace, Suit: model.Spade},
				{Rank: model.Ace, Suit: model.Diamond},
				{Rank: model.Three, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.TwoPairs))
			Expect(pokerHands.Strengths).To(Equal([]int{14, 7, 3}))
		})
	})

	Context("OnePair", func() {
		It("should return OnePair", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Seven, Suit: model.Heart},
				{Rank: model.Queen, Suit: model.Spade},
				{Rank: model.Ace, Suit: model.Diamond},
				{Rank: model.Three, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.OnePair))
			Expect(pokerHands.Strengths).To(Equal([]int{7, 14, 12, 3}))
		})
	})

	Context("HighCard", func() {
		It("should return HighCard", func() {
			pokerHands, err := AnalyzePokerHands(model.Cards{
				{Rank: model.Seven, Suit: model.Club},
				{Rank: model.Queen, Suit: model.Heart},
				{Rank: model.Eight, Suit: model.Spade},
				{Rank: model.Ace, Suit: model.Diamond},
				{Rank: model.Three, Suit: model.Heart},
			})
			Expect(err).To(BeNil())
			Expect(pokerHands.PokerHandsType).To(Equal(model.HighCard))
			Expect(pokerHands.Strengths).To(Equal([]int{14, 12, 8, 7, 3}))
		})
	})
})
