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

func ParseLog(fileName string) {
	log.Print("Parsing log: ", fileName)
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		parseLine(scanner.Text())
	}
}

func parseLine(text string) {
	switch {
	case rxSupply.MatchString(text):
		handleSupply(text)
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

func handleSupply(text string) {
	fmt.Println("using supply cards")
	cards := parseCards(strings.Split(text, "cards:")[1])
	fmt.Println(cards)
}

func handleDraw(text string) {
	player, cardsText := parseAction(text, "draws")
	fmt.Printf("%v draws cards", player)
	cards := parseCards(cardsText)
	fmt.Println(cards)
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

// returns slice of CardGroups, can parse both
// Copper, Copper, Copper, Copper, Estate
// 2 Copper, 1 Gold, 1 Silver
func parseCards(text string) []card.CardGroup {
	cardTextList := strings.Split(text, ",")
	cardGroups := []card.CardGroup{}
	var cardGroup card.CardGroup
	for _, cardText := range cardTextList {
		cardText := strings.TrimSpace(cardText)
		if rxNumCards.MatchString(cardText) {
			cardWithNum := strings.Split(cardText, " ")
			num, err := strconv.Atoi(cardWithNum[0])
			check(err)
			cardText = strings.TrimSpace(cardWithNum[1])
			cardGroup = card.CardGroup{Num: num, Card: card.CardList[cardText]}
		} else {
			cardGroup = card.CardGroup{Num: 1, Card: card.CardList[cardText]}
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
