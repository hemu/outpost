package logparse

import (
	"bufio"
	"fmt"
	// "io/ioutil"
	"github.com/hmuar/dominion-replay/card"
	"log"
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

var rxSupply, _ = regexp.Compile("^Supply cards:.*$")
var rxDraw, _ = regexp.Compile(".*draws.*$")
var rxPlay, _ = regexp.Compile(".*plays.*$")
var rxBuy, _ = regexp.Compile(".*buys.*$")
var rxGain, _ = regexp.Compile(".*gains.*$")
var rxDiscard, _ = regexp.Compile(".*discards.*$")
var rxTurn, _ = regexp.Compile(".*turn.*$")
var rxNumCards, _ = regexp.Compile("^.*[0-9] .*$")

type event struct {
	player string
	action string
	cards  []card.CardSet
}

type turn struct {
	player string
	events []event
}

type Game struct {
	logFile string
	players []string
	supply  []card.CardSet
	turns   []turn
	rating  string
	winner  string
}

// returns Game
// Game contains a slice of []turns
// each turn is  a slice of []event
// an event has a player, action, and cards []card.Card
func ParseLog(fileName string) Game {
	fileName = fileName
	log.Print("Parsing log: ", fileName)
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	game := Game{logFile: fileName}

	for scanner.Scan() {
		parseLine(scanner.Text(), game)
	}
	return game
}

func parseLine(text string, game Game) {
	switch {
	case rxSupply.MatchString(text):
		supplyCards := handleSupply(text)
		game.supply = supplyCards

		// player draws
		// event: action is 'draw'
	case rxDraw.MatchString(text):
		handleDraw(text)
	case rxPlay.MatchString(text):
		handlePlay(text)
	case rxBuy.MatchString(text):
		handleBuy(text)
	case rxGain.MatchString(text):
		handleGain(text)
	case rxDiscard.MatchString(text):
		handleDiscard(text)
	case rxTurn.MatchString(text):
		handleTurn(text)

	}
}

func handleSupply(text string) []card.CardSet {
	cards := parseCards(strings.Split(text, "cards:")[1])
	return cards
}

func handleDraw(text string) []card.CardSet {
	player, cardsText := parseAction(text, "draws")
	fmt.Printf("%v draws cards", player)
	cards := parseCards(cardsText)
	return cards
}

func handlePlay(text string) {
	player, cardsText := parseAction(text, "plays")
	fmt.Printf("%v plays cards", player)
	cardsWithNum := parseCards(cardsText)
	fmt.Println(cardsWithNum)
}

func handleBuy(text string) {

}

func handleGain(text string) {

}

func handleDiscard(text string) {

}

func handleTurn(text string) {

}

// returns slice of CardSets, can parse both
// Copper, Copper, Copper, Copper, Estate
// 2 Copper, 1 Gold, 1 Silver
func parseCards(text string) []card.CardSet {
	cardTextList := strings.Split(text, ",")
	cardGroups := []card.CardSet{}
	var cardGroup card.CardSet
	for _, cardText := range cardTextList {
		cardText := strings.TrimSpace(cardText)
		if rxNumCards.MatchString(cardText) {
			cardWithNum := strings.Split(cardText, " ")
			num, err := strconv.Atoi(cardWithNum[0])
			check(err)
			cardText = strings.TrimSpace(cardWithNum[1])
			cardGroup = card.CardSet{Num: num, Card: card.CardFactory[cardText]}
		} else {
			cardGroup = card.CardSet{Num: 1, Card: card.CardFactory[cardText]}
		}
		cardGroups = append(cardGroups, cardGroup)
	}
	return cardGroups
}

func parseAction(text, action string) (string, string) {
	playerWithCards := strings.Split(text, "- "+action)
	player := strings.TrimSpace(playerWithCards[0])
	cardsText := playerWithCards[1]
	return player, cardsText
}
