package state

import (
	"errors"
	"fmt"
	mCard "github.com/hmuar/dominion-replay/card"
)

type Player struct {
	Action   int
	Buy      int
	Coin     int
	Victory  int
	Hand     []mCard.Card
	Draw     []mCard.Card
	Discard  []mCard.Card
	Duration []mCard.Card
}

type Board struct {
	Trash  []mCard.Card
	Supply map[string]*mCard.CardSet
	Play   []mCard.Card
}

type State struct {
	Players    map[string]*Player
	Board      Board
	TurnPlayer string
	TurnNum    int
}

func (s *State) Print() {
	fmt.Println("------------------")
	fmt.Println("[State]")
	fmt.Printf("  [TurnNum]: %d\n", s.TurnNum)
	fmt.Printf("  [TurnPlayer]: %v\n", s.TurnPlayer)
	fmt.Printf("  [Board]\n")
	fmt.Printf("    [Supply]\n")
	for _, v := range s.Board.Supply {
		cardSet := *v
		fmt.Printf("      [%d] %v\n", cardSet.Num, cardSet.Card.Name)
	}
	fmt.Printf("    [Trash]: %v\n", s.Board.Trash)
	fmt.Printf("    [Play]: %v\n", s.Board.Play)
	for name, player := range s.Players {
		fmt.Printf("  [Player %v]\n", name)
		fmt.Printf("    [A,B,C,V]: [%d,%d,%d,%d]\n",
			player.Action,
			player.Buy,
			player.Coin,
			player.Victory)
		fmt.Printf("    [Hand]: %v\n", player.Hand)
		fmt.Printf("    [Draw]: %v\n", player.Draw)
		fmt.Printf("    [Discard]: %v\n", player.Discard)
		fmt.Printf("    [Duration]: %v\n", player.Duration)
	}
	fmt.Println("------------------")
}

func (s *State) PrintPlayers() {
	fmt.Println("-------------------")
	fmt.Printf("[TurnNum]: %d\n", s.TurnNum)
	fmt.Printf("[TurnPlayer]: %v\n", s.TurnPlayer)
	for name, player := range s.Players {
		fmt.Printf("[Player %v]\n", name)
		fmt.Printf("  [A,B,C,V]: [%d,%d,%d,%d]\n",
			player.Action,
			player.Buy,
			player.Coin,
			player.Victory)
		fmt.Printf("  [Hand]: %v\n", player.Hand)
		fmt.Printf("  [Draw]: %v\n", player.Draw)
		fmt.Printf("  [Discard]: %v\n", player.Discard)
		fmt.Printf("  [Duration]: %v\n", player.Duration)
	}
}

func removeFromCards(removeCard mCard.Card, cards []mCard.Card) ([]mCard.Card, error) {
	for i, card := range cards {
		if card.Name == removeCard.Name {
			prunedCards := append(cards[:i], cards[i+1:]...)
			return prunedCards, nil
		}
	}
	return cards, errors.New(fmt.Sprintf("Could not find %v in cards", removeCard.Name))
}

func (s *State) SetPlayers(players []string) {
	s.Players = make(map[string]*Player)
	for _, player := range players {
		s.Players[player] = &Player{}
	}
}

func (s *State) SetTurnPlayer(player string) {
	s.TurnPlayer = player
}

func (s *State) getPlayer(player string) (*Player, bool) {
	val, exists := s.Players[player]
	return val, exists
}

func (s *State) SetSupply(cardSets []mCard.CardSet) {
	s.Board.Supply = make(map[string]*mCard.CardSet)
	for i := 0; i < len(cardSets); i++ {
		cardSet := cardSets[i]
		s.Board.Supply[cardSet.Card.Name] = &cardSet
	}
}

func (s *State) RemoveSupplyCard(cardName string) (mCard.Card, error) {
	cardSet, exists := s.Board.Supply[cardName]
	if exists {
		if cardSet.Num > 0 {
			cardSet.Num = cardSet.Num - 1
			return cardSet.Card, nil
		}
	}
	return mCard.NewCard("UNKNOWN"), errors.New("Could not get card from supply")
}

func (s *State) SetHand(player string, cards []mCard.Card) {
	p, exists := s.getPlayer(player)
	if exists {
		p.Hand = cards
	} else {
		fmt.Println("Could not find player %v", player)
	}
}

func (s *State) GetHand(player string) []mCard.Card {
	p, _ := s.getPlayer(player)
	return p.Hand
}

func (s *State) RemoveFromDraw(player string, cards []mCard.Card) {
	p, _ := s.getPlayer(player)
	remainCards := p.Draw
	for _, card := range cards {
		remainCards, _ = removeFromCards(card, remainCards)
	}
	p.Draw = remainCards
}

func (s *State) SetDraw(player string, cards []mCard.Card) {
	p, exists := s.getPlayer(player)
	if exists {
		p.Draw = cards
	}
}

func (s *State) GetDraw(player string) []mCard.Card {
	p, _ := s.getPlayer(player)
	return p.Draw
}

func (s *State) AddToDraw(player string, cards []mCard.Card) {
	p, _ := s.getPlayer(player)
	for _, card := range cards {
		p.Draw = append(p.Draw, card)
	}
}

func (s *State) SetDiscard(player string, cards []mCard.Card) {
	p, _ := s.getPlayer(player)
	p.Discard = cards
}

func (s *State) GetDiscard(player string) []mCard.Card {
	p, _ := s.getPlayer(player)
	return p.Discard
}

func (s *State) AddDiscardCard(player string, card mCard.Card) {
	p, exists := s.getPlayer(player)
	if exists {
		p.Discard = append(p.Discard, card)
	}
}

func (s *State) GetPlay() []mCard.Card {
	return s.Board.Play
}

func (s *State) AddPlay(card mCard.Card) {
	s.Board.Play = append(s.Board.Play, card)
}

func (s *State) SetPlay(cards []mCard.Card) {
	s.Board.Play = cards
}

func (s *State) RemoveFromHand(card mCard.Card) error {
	p, exists := s.getPlayer(s.TurnPlayer)
	if exists {
		remainCards, err := removeFromCards(card, p.Hand)
		if err == nil {
			p.Hand = remainCards
		}
		return err
	} else {
		return errors.New("Could not find current player")
	}
}

func (s *State) ClearPlayerStats(player string) {
	p, exists := s.getPlayer(player)
	if exists {
		p.Action = 0
		p.Buy = 0
		p.Coin = 0
	}
}

func (s *State) ResetPlayerStats(player string) {
	p, exists := s.getPlayer(player)
	if exists {
		p.Action = 1
		p.Buy = 1
	}
}

func (s *State) GetPlayerStats(player string) [4]int {
	p, exists := s.getPlayer(player)
	if !exists {
		return [4]int{}
	}
	return [4]int{p.Action, p.Buy, p.Coin, p.Victory}
}

func (s *State) AddPlayerStats(player string, stats [4]int) {
	s.AddPlayerActionCount(player, stats[0])
	s.AddPlayerBuyCount(player, stats[1])
	s.AddPlayerCoinCount(player, stats[2])
	s.AddPlayerVictoryCount(player, stats[3])
}

func (s *State) AddPlayerActionCount(player string, val int) {
	p, exists := s.getPlayer(player)
	if exists {
		p.Action = p.Action + val
	}
}

func (s *State) AddPlayerBuyCount(player string, val int) {
	p, exists := s.getPlayer(player)
	if exists {
		p.Buy = p.Buy + val
	}
}

func (s *State) AddPlayerCoinCount(player string, val int) {
	p, exists := s.getPlayer(player)
	if exists {
		p.Coin = p.Coin + val
	}
}

func (s *State) AddPlayerVictoryCount(player string, val int) {
	p, exists := s.getPlayer(player)
	if exists {
		p.Victory = p.Victory + val
	}
}
