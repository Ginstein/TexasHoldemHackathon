package model

import (
	"errors"
	"fmt"
)

var (
	MinPlayerCountErr   = errors.New(fmt.Sprintf("player count must be greater than %d", MinPlayerCountLimit))
	HoldCardsCountErr   = errors.New(fmt.Sprintf("hold cards count must be %d", HoldCardsCountLimit))
	PickedCardsCountErr = errors.New(fmt.Sprintf("picked cards count must be %d", PickedCardsCountLimit))
	PublicCardsCountErr = errors.New(fmt.Sprintf("public cards count must be %d", PickedCardsCountLimit))
	CardsCheckErr       = errors.New("cards illegal")
)
