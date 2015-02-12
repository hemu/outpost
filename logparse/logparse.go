package logparse

import (
	"bufio"
	"errors"
	"fmt"
	mCard "github.com/hmuar/dominion-replay/card"
	mHistory "github.com/hmuar/dominion-replay/history"
	// "log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var rxGameSetup, _ = regexp.Compile(".*Game Setup.*$")
var rxSupply, _ = regexp.Compile("^Supply cards:.*$")
var rxDraw, _ = regexp.Compile(".*draws.*$")
var rxPlay, _ = regexp.Compile(".*plays.*$")
var rxBuy, _ = regexp.Compile(".*buys.*$")
var rxGain, _ = regexp.Compile(".*gains.*$")
var rxDiscard, _ = regexp.Compile(".*discards.*$")
var rxTrash, _ = regexp.Compile(".*trashes.*$")
var rxShuffle, _ = regexp.Compile(".*shuffles.*$")
var rxPlaceOnDeck, _ = regexp.Compile(".*places.*on top of deck.*$")
var rxLookAt, _ = regexp.Compile(".*looks.*$")
var rxTurn, _ = regexp.Compile(".*turn.*$")
var rxNumCards, _ = regexp.Compile("^.*[0-9] .*$")

// returns Game
// Game contains a slice of []turns
// each turn is  a slice of []event
// an event has a player, action, and cards []mCard.Card
func ParseLog(fileName string) mHistory.History {
	fileName = fileName
	// log.Println("Parsing log: ", fileName)
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	gBuilder := mHistory.NewHistoryBuilder()

	for scanner.Scan() {
		parseLine(scanner.Text(), &gBuilder)
	}
	return gBuilder.History
}

func parseLine(text string, gBuilder *mHistory.HistoryBuilder) {
	switch {
	case rxSupply.MatchString(text):
		supplyCards := handleSupply(text)
		gBuilder.SetSupply(supplyCards)

	// player event: start turn
	case rxTurn.MatchString(text):
		player, turnNum := handleTurn(text)
		gBuilder.StartPlayerTurn(player, turnNum)

	// player draws -- event: action 'draw'
	case rxDraw.MatchString(text):
		player, cards := handleDraw(text)
		gBuilder.AddEvent(player, mHistory.ACTION_DRAW, cards)

	// player plays -- event: action 'play'
	case rxPlay.MatchString(text):
		player, cards := handlePlay(text)
		gBuilder.AddEvent(player, mHistory.ACTION_PLAY, cards)

	// player buys -- event: action 'buy'
	case rxBuy.MatchString(text):
		player, cards := handleBuy(text)
		gBuilder.AddEvent(player, mHistory.ACTION_BUY, cards)

	// player gains -- event: action 'gain'
	case rxGain.MatchString(text):
		player, cards := handleGain(text)
		gBuilder.AddEvent(player, mHistory.ACTION_GAIN, cards)

	// player discards -- event: action 'discard'
	case rxDiscard.MatchString(text):
		player, cards := handleDiscard(text)
		gBuilder.AddEvent(player, mHistory.ACTION_DISCARD, cards)

	// player places cards on top of dec - event: action 'place-on-deck'
	case rxPlaceOnDeck.MatchString(text):
		player, cards := handlePlaceOnDeck(text)
		gBuilder.AddEvent(player, mHistory.ACTION_PLACE_ON_DECK, cards)

	// player looks at cards from deck - event: action 'look-at'
	case rxLookAt.MatchString(text):
		player, cards := handleLookAt(text)
		gBuilder.AddEvent(player, mHistory.ACTION_LOOK_AT, cards)

	// player looks at cards from deck - event: action 'trash'
	case rxShuffle.MatchString(text):
		player, _, err := parsePlayerWithAction(text, "shuffles")
		if err != nil {
			check(err)
		}
		gBuilder.AddEvent(player, mHistory.ACTION_SHUFFLE, []mCard.CardSet{})

	// player discards -- event: action is 'discard'
	case rxTrash.MatchString(text):
		player, cards := handleTrash(text)
		gBuilder.AddEvent(player, mHistory.ACTION_TRASH, cards)

	case rxGameSetup.MatchString(text):
		gBuilder.RegisterGameSetup()

	}

}

func handleSupply(text string) []mCard.CardSet {
	cards := parseCards(strings.Split(text, "cards:")[1])
	return cards
}

func handleTurn(text string) (string, int) {
	stripped := strings.Replace(text, "-", "", -1)
	trimmed := strings.TrimSpace(stripped)
	playerWithNum := strings.Split(trimmed, ": turn ")
	turnNum, _ := strconv.Atoi(playerWithNum[1])
	return playerWithNum[0], turnNum
}

func handleActionWithCards(text string, action string) (string, []mCard.CardSet) {
	player, cardsText, err := parseActionWithCards(text, action)
	if err != nil {
		return "", []mCard.CardSet{}
	}
	cards := parseCards(cardsText)
	return player, cards
}

func handleDraw(text string) (string, []mCard.CardSet) {
	return handleActionWithCards(text, "draws")
}

func handlePlay(text string) (string, []mCard.CardSet) {
	return handleActionWithCards(text, "plays")
}

func handleBuy(text string) (string, []mCard.CardSet) {
	return handleActionWithCards(text, "buys")
}

func handleGain(text string) (string, []mCard.CardSet) {
	return handleActionWithCards(text, "gains")
}

func handleDiscard(text string) (string, []mCard.CardSet) {
	return handleActionWithCards(text, "discards")
}

func handlePlaceOnDeck(text string) (string, []mCard.CardSet) {
	player, actionText, err := parsePlayerWithAction(text, "places")
	check(err)
	cardName := strings.TrimSpace(strings.Split(actionText, " ")[0])
	cardSet := mCard.CardSet{Num: 1, Card: mCard.CardFactory[cardName]}
	cardSets := []mCard.CardSet{}
	cardSets = append(cardSets, cardSet)
	return player, cardSets
}

func handleLookAt(text string) (string, []mCard.CardSet) {
	// player, actionText, err := parsePlayerWithAction(text, "looks at")
	player, cardsText, err := parsePlayerWithAction(text, "looks at")
	check(err)
	cardTextList := strings.Split(cardsText, ",")
	cardSets := []mCard.CardSet{}
	for _, cardText := range cardTextList {
		cardSet := mCard.CardSet{Num: 1,
			Card: mCard.CardFactory[strings.TrimSpace(cardText)]}
		cardSets = append(cardSets, cardSet)
	}
	return player, cardSets
}

func handleTrash(text string) (string, []mCard.CardSet) {
	return handleActionWithCards(text, "trashes")
}

// returns slice of CardSets, can parse both
// Copper, Copper, Copper, Copper, Estate
// 2 Copper, 1 Gold, 1 Silver
func parseCards(text string) []mCard.CardSet {

	var cardGroup mCard.CardSet
	var num int
	var err error
	var cardName string
	cardGroups := []mCard.CardSet{}
	cardCount := make(map[string]int)
	cardOrder := []string{}

	cardTextList := strings.Split(text, ",")

	for _, cardText := range cardTextList {
		cardText := strings.TrimSpace(cardText)
		// if this is a number followed by card
		// e.g. 3 Copper
		if rxNumCards.MatchString(cardText) {
			cardWithNum := strings.Split(cardText, " ")
			num, err = strconv.Atoi(cardWithNum[0])
			check(err)
			cardName = strings.TrimSpace(cardWithNum[1])
		} else {
			num = 1
			cardName = cardText
		}
		val, exists := cardCount[cardName]
		if exists {
			cardCount[cardName] = val + num
		} else {
			cardCount[cardName] = num
			cardOrder = append(cardOrder, cardName)
		}
	}

	for _, cardName := range cardOrder {
		cardGroup = mCard.CardSet{Num: cardCount[cardName], Card: mCard.CardFactory[cardName]}
		cardGroups = append(cardGroups, cardGroup)
	}

	return cardGroups
}

func parseActionWithCards(text, action string) (string, string, error) {
	player, actionText, err := parsePlayerWithAction(text, action)
	if err != nil {
		return "", "", err
	}
	cardsText := actionText
	return player, cardsText, nil
}

func parsePlayerWithAction(text, action string) (string, string, error) {
	playerWithActionText := strings.Split(text, "- "+action)
	if len(playerWithActionText) != 2 {
		return "", "", errors.New(fmt.Sprintf("PARSE ERROR: Could not find action %v in '%v'",
			action,
			text))
	}
	player := strings.TrimSpace(playerWithActionText[0])
	actionText := strings.TrimSpace(playerWithActionText[1])
	return player, actionText, nil
}
