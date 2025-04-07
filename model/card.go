package model

type CardRank string

const (
	Two   CardRank = "2"
	Three CardRank = "3"
	Four  CardRank = "4"
	Five  CardRank = "5"
	Six   CardRank = "6"
	Seven CardRank = "7"
	Eight CardRank = "8"
	Nine  CardRank = "9"
	Ten   CardRank = "10"
	Jack  CardRank = "J"
	Queen CardRank = "Q"
	King  CardRank = "K"
	Ace   CardRank = "A"
)

var CardRanks = []CardRank{Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}

var CardRanksValueMap = map[CardRank]int{
	Two:   2,
	Three: 3,
	Four:  4,
	Five:  5,
	Six:   6,
	Seven: 7,
	Eight: 8,
	Nine:  9,
	Ten:   10,
	Jack:  11,
	Queen: 12,
	King:  13,
	Ace:   14,
}

type CardSuit string

const (
	Spade   CardSuit = "♠"
	Heart   CardSuit = "♥"
	Club    CardSuit = "♣"
	Diamond CardSuit = "♦"
)

var CardSuits = []CardSuit{Club, Diamond, Heart, Spade}

var CardSuitsValueMap = map[CardSuit]int{
	Spade:   4,
	Heart:   3,
	Club:    2,
	Diamond: 1,
}

type Card struct {
	Rank CardRank
	Suit CardSuit
}

type Cards []Card
