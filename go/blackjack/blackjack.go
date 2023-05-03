package blackjack

const (
	hit   = "H"
	split = "P"
	stand = "S"
	win   = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// isBlackjack returns true if the player has a blackjack, false otherwise.
func isBlackjack(card1, card2 string) bool {
	return ParseCard(card1)+ParseCard(card2) == 21
}

// largeHand implements the decision tree for hand scores larger than 20 points.
func largeHand(isBlackjack bool, dealerScore int) string {
	if !isBlackjack {
		return split
	}
	switch dealerScore {
	case 10, 11:
		return stand
	default:
		return win
	}
}

// smallHand implements the decision tree for hand scores with less than 21 points.
func smallHand(handScore, dealerScore int) string {
	switch {
	case handScore > 16:
		return stand
	case handScore < 12:
		return hit
	default:
		if dealerScore > 6 {
			return hit
		}
		return stand
	}
}
func FirstTurn(card1, card2, dealerCard string) string {
	handScore := ParseCard(card1) + ParseCard(card2)
	switch {
	case handScore < 21:
		return smallHand(handScore, ParseCard(dealerCard))
	default:
		return largeHand(isBlackjack(card1, card2), ParseCard(dealerCard))
	}
}
