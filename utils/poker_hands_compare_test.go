package utils_test

import (
	"github.com/Ginstein/TexasHoldemHackathon/model"
	"github.com/Ginstein/TexasHoldemHackathon/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PokerHandsCompare", func() {
	Context("PokerHandsCompare", func() {
		It("result should be greater", func() {
			var pokerHandsI = model.PokerHands{
				PokerHandsType: model.RoyalFlush,
				Strengths:      []int{14},
			}
			var pokerHandsJ = model.PokerHands{
				PokerHandsType: model.OnePair,
				Strengths:      []int{13, 7, 5, 2},
			}
			var result = utils.PokerHandsCompare(pokerHandsI, pokerHandsJ)
			Expect(result).To(Equal(model.Greater))
		})

		It("result should be greater", func() {
			var pokerHandsI = model.PokerHands{
				PokerHandsType: model.TwoPairs,
				Strengths:      []int{13, 11, 5},
			}
			var pokerHandsJ = model.PokerHands{
				PokerHandsType: model.TwoPairs,
				Strengths:      []int{13, 10, 5},
			}
			var result = utils.PokerHandsCompare(pokerHandsI, pokerHandsJ)
			Expect(result).To(Equal(model.Greater))
		})

		It("result should be equal", func() {
			var pokerHandsI = model.PokerHands{
				PokerHandsType: model.Straight,
				Strengths:      []int{7},
			}
			var pokerHandsJ = model.PokerHands{
				PokerHandsType: model.Straight,
				Strengths:      []int{7},
			}
			var result = utils.PokerHandsCompare(pokerHandsI, pokerHandsJ)
			Expect(result).To(Equal(model.Equal))
		})

		It("result should be less", func() {
			var pokerHandsI = model.PokerHands{
				PokerHandsType: model.TwoPairs,
				Strengths:      []int{11, 6, 5},
			}
			var pokerHandsJ = model.PokerHands{
				PokerHandsType: model.Straight,
				Strengths:      []int{7},
			}
			var result = utils.PokerHandsCompare(pokerHandsI, pokerHandsJ)
			Expect(result).To(Equal(model.Less))
		})

		It("result should be less", func() {
			var pokerHandsI = model.PokerHands{
				PokerHandsType: model.ThreeOfAKind,
				Strengths:      []int{9, 5, 4},
			}
			var pokerHandsJ = model.PokerHands{
				PokerHandsType: model.ThreeOfAKind,
				Strengths:      []int{9, 14, 13},
			}
			var result = utils.PokerHandsCompare(pokerHandsI, pokerHandsJ)
			Expect(result).To(Equal(model.Less))
		})
	})
})
