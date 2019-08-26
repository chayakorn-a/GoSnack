package main

import "fmt"

const DEUCE int = 4
const ADV int = 5
const WIN int = 6

var SCORE = [...]string{"Love", "Fifteen", "Thirty", "Forty", "Deuce", "Advantage for", "Wins"}

type Player struct {
	name  string
	point int
}

type Game struct {
	PlayerA, PlayerB Player
}

func (g Game) CurrentScore() {
	if g.PlayerA.point < 4 && g.PlayerB.point < 4 && (g.PlayerA.point < 3 || g.PlayerB.point < 3) {
		fmt.Printf("%s - %s\n", SCORE[g.PlayerA.point], SCORE[g.PlayerB.point])
	} else if g.PlayerA.point == g.PlayerB.point {
		fmt.Printf("%s\n", SCORE[DEUCE])
	} else if g.PlayerA.point-g.PlayerB.point == 1 {
		fmt.Printf("%s %s\n", SCORE[ADV], g.PlayerA.name)
	} else if g.PlayerB.point-g.PlayerA.point == 1 {
		fmt.Printf("%s %s\n", SCORE[ADV], g.PlayerB.name)
	} else { // someone win
		if g.PlayerA.point > g.PlayerB.point {
			fmt.Printf("%s %s\n", g.PlayerA.name, SCORE[WIN])
		} else {
			fmt.Printf("%s %s\n", g.PlayerB.name, SCORE[WIN])
		}
	}
}

func (g *Game) PlayerAGetPoint() {
	g.PlayerA.point++
}

func (g *Game) PlayerBGetPoint() {
	g.PlayerB.point++
}

func main() {
	g := Game{Player{"Aman", 0}, Player{"BB", 0}}

	g.CurrentScore()
	g.PlayerAGetPoint()
	g.PlayerBGetPoint()
	g.PlayerBGetPoint()
	g.PlayerBGetPoint()
	g.PlayerAGetPoint()
	g.PlayerAGetPoint()
	g.PlayerAGetPoint()
	g.PlayerBGetPoint()
	g.PlayerBGetPoint()
	g.PlayerBGetPoint()
	g.CurrentScore()
}
