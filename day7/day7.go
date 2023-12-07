package day7

import (
	"fmt"
	"github.com/the-medo/go-advent-2023/utils"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand          string
	handStrength  int
	cardStrengths []int
}

type HandWithBid struct {
	hand     *Hand
	bid      int
	strength int
}

func Solve(input string) {
	handRows := utils.SplitRows(input)
	hands := make([]HandWithBid, len(handRows))
	for i, h := range handRows {
		hSplit := strings.Split(h, " ")
		bid, _ := strconv.Atoi(hSplit[1])

		hands[i] = HandWithBid{
			hand: getHandStruct(hSplit[0]),
			bid:  bid,
		}
	}

	sort.Slice(hands, func(i, j int) bool {
		return compareHandStrength(*hands[i].hand, *hands[j].hand) > 0
	})

	part1 := 0
	for i, j := range hands {
		part1 += j.bid * (len(hands) - i)
		fmt.Println(j, j.hand)
	}
	fmt.Println("Part 1: ", part1)
}

func getHandStruct(handString string) *Hand {
	handStrengths := map[string]int{
		"FiveOfAKind":  7,
		"FourOfAKind":  6,
		"FullHouse":    5,
		"ThreeOfAKind": 4,
		"TwoPairs":     3,
		"OnePair":      2,
		"HighCard":     1,
	}

	cards := make([]int, 5)
	availableCards := make(map[rune]int)

	for i, x := range handString {
		cards[i] = getCardStrength(x)
		_, exists := availableCards[x]
		if exists {
			availableCards[x]++
		} else {
			availableCards[x] = 1
		}
	}

	result := "HighCard"
	if len(availableCards) == 1 {
		result = "FiveOfAKind"
	} else if len(availableCards) == 2 { // 4 o k, full house
		result = "FullHouse"
		for _, v := range availableCards {
			if v == 4 {
				result = "FourOfAKind"
			}
		}
	} else if len(availableCards) == 3 { // 3 o k, two pairs
		result = "TwoPairs"
		for _, v := range availableCards {
			if v == 3 {
				result = "ThreeOfAKind"
			}
		}

	} else if len(availableCards) == 4 {
		result = "OnePair"
	}

	fmt.Println("Hand string: ", handString, result, availableCards)

	return &Hand{
		hand:          handString,
		handStrength:  handStrengths[result],
		cardStrengths: cards,
	}
}

func compareHandStrength(hand1 Hand, hand2 Hand) int {
	if hand1.handStrength == hand2.handStrength {

		for i, _ := range hand1.cardStrengths {
			diff := hand1.cardStrengths[i] - hand2.cardStrengths[i]
			if diff != 0 {
				return diff
			}
		}
	}
	return hand1.handStrength - hand2.handStrength
}

func compareCardStrength(card1 rune, card2 rune) int {
	return getCardStrength(card1) - getCardStrength(card2)
}

func getCardStrength(card rune) int {

	cardStrengths := map[rune]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'J': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
	}

	return cardStrengths[card]
}
