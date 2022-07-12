// Shelby Simpson
// Module 6 Activity: Black Jack
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type BlackJackGame struct {
	Deck    []Card
	dealer  Player
	Players []Player
}

type Player struct {
	Total int
	Hand  []Card
}

type Card struct {
	ValueNum int
	ValueStr string
	Suit     string
}

func (g *BlackJackGame) GenerateDeck() {
	g.Deck = []Card{
		{1, "Ace", "Spades"}, {2, "Two", "Spades"},
		{3, "Three", "Spades"}, {4, "Four", "Spades"},
		{5, "Five", "Spades"}, {6, "Six", "Spades"},
		{7, "Seven", "Spades"}, {8, "Eight", "Spades"},
		{9, "Nine", "Spades"}, {10, "Ten", "Spades"},
		{10, "Jack", "Spades"}, {10, "Queen", "Spades"},
		{10, "King", "Spades"}, {1, "Ace", "Hearts"},
		{2, "Two", "Hearts"}, {3, "Three", "Hearts"},
		{4, "Four", "Hearts"}, {5, "Five", "Hearts"},
		{6, "Six", "Hearts"}, {7, "Seven", "Hearts"},
		{8, "Eight", "Hearts"}, {9, "Nine", "Hearts"},
		{10, "Ten", "Hearts"}, {10, "Jack", "Hearts"},
		{10, "Queen", "Hearts"}, {10, "King", "Hearts"},
		{1, "Ace", "Diamonds"}, {2, "Two", "Diamonds"},
		{3, "Three", "Diamonds"}, {4, "Four", "Diamonds"},
		{5, "Five", "Diamonds"}, {6, "Six", "Diamonds"},
		{7, "Seven", "Diamonds"}, {8, "Eight", "Diamonds"},
		{9, "Nine", "Diamonds"}, {10, "Ten", "Diamonds"},
		{10, "Jack", "Diamonds"}, {10, "Queen", "Diamonds"},
		{10, "King", "Diamonds"}, {1, "Ace", "Clubs"},
		{2, "Two", "Clubs"}, {3, "Three", "Clubs"},
		{4, "Four", "Clubs"}, {5, "Five", "Clubs"},
		{6, "Six", "Clubs"}, {7, "Seven", "Clubs"},
		{8, "Eight", "Clubs"}, {9, "Nine", "Clubs"},
		{10, "Ten", "Clubs"}, {10, "Jack", "Clubs"},
		{10, "Queen", "Clubs"}, {10, "King", "Clubs"},
	}
}

func (g *BlackJackGame) ShuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(g.Deck), func(i, j int) { g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i] })
}

func (g *BlackJackGame) Deal() {
	for i := 0; i < 2; i++ {
		g.handoutCard(&g.dealer)
		for j := range g.Players {
			g.handoutCard(&g.Players[j])
		}
	}
}

func (g *BlackJackGame) handoutCard(p *Player) {
	cardNum := rand.Intn(len(g.Deck))
	card := g.Deck[cardNum]
	p.Hand = append(p.Hand, card)
	p.Total += card.ValueNum
	g.Deck = append(g.Deck[0:cardNum], g.Deck[cardNum+1:]...)
}

// The dealer hits until his total is 17 or more
func (g *BlackJackGame) dealersTurn() {
	for g.dealer.Total < 17 {
		g.handoutCard(&g.dealer)
	}
}

func main() {
	game := BlackJackGame{
		Deck:    []Card{},
		dealer:  Player{Total: 0, Hand: []Card{}},
		Players: []Player{},
	}
	game.GenerateDeck()
	game.ShuffleDeck()
	player1 := Player{
		Total: 0,
		Hand:  []Card{},
	}
	game.Players = append(game.Players, player1)
	game.Deal()
	fmt.Println(game.dealer.Total)
	fmt.Println(game.Players[0].Total)
	game.dealersTurn()
	fmt.Println(game.dealer.Total)
}
