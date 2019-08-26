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

func (p *Player) Get() {
	p.point++
}

type Game struct {
	PlayerA, PlayerB Player
}

func TennisGame(n1, n2 string) (func(), func(), func()) {
	g := Game{Player{n1, 0}, Player{n2, 0}}

	return func() { // CurrentScore()
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

		}, func() { // PlayerAGetPoint()
			g.PlayerA.Get()
		}, func() { // PlayerBGetPoint()
			g.PlayerB.Get()
		}
}

func main() {
	CurrentScore, PlayerAGetPoint, PlayerBGetPoint := TennisGame("Aman", "BB")

	CurrentScore()
	PlayerAGetPoint()
	PlayerBGetPoint()
	PlayerBGetPoint()
	PlayerBGetPoint()
	PlayerAGetPoint()
	PlayerAGetPoint()
	PlayerAGetPoint()
	PlayerBGetPoint()
	PlayerBGetPoint()
	PlayerBGetPoint()
	CurrentScore()
}
