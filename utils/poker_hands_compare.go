package utils

import "github.com/Ginstein/TexasHoldemHackathon/model"

// PokerHandsCompare 计算手牌得分，用于最终结果比较
// 先比较PokerHandsType再比较strengths
// CompareResult
func PokerHandsCompare(pokerHandsI, pokerHandsJ model.PokerHands) model.CompareResult {
	var pokerHandsIWeight = model.PokerHandsTypeWeightMap[pokerHandsI.PokerHandsType]
	var pokerHandsJWeight = model.PokerHandsTypeWeightMap[pokerHandsJ.PokerHandsType]
	if pokerHandsIWeight > pokerHandsJWeight {
		return model.Greater
	}
	if pokerHandsIWeight < pokerHandsJWeight {
		return model.Less
	}
	// 比较 strengths
	for index := 0; index < len(pokerHandsI.Strengths); index++ {
		if pokerHandsI.Strengths[index] > pokerHandsJ.Strengths[index] {
			return model.Greater
		}
		if pokerHandsI.Strengths[index] < pokerHandsJ.Strengths[index] {
			return model.Less
		}
	}
	return model.Equal
}
