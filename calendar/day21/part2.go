package adventofcode

import (
	util "adventofcode/util/common"
	"fmt"
)

type diracDice struct {
	f int
}

func (d diracDice) outcomes() []int {
	return []int{3, 4, 5, 6, 7, 8, 9}
}

func (d diracDice) probabilities() map[int]int {
	return map[int]int{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}
}

type gamestate struct {
	p1 player
	p2 player
}

type roundGamesMap map[int]map[gamestate]uint64

func Part2(p1s, p2s int) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	die := diracDice{}

	p1wins := uint64(0)
	p2wins := uint64(0)
	initialGS := gamestate{p1: player{num: 1, pos: p1s - 1, score: 0}, p2: player{num: 2, pos: p2s - 1, score: 0}}

	currentGames := map[gamestate]uint64{}
	currentGames[initialGS] = 1
	currentPlayerP1 := true
	for len(currentGames) > 0 {
		nextGames := map[gamestate]uint64{}
		for game, universes := range currentGames {
			util.IncreasedebugIndent()
			//util.Dprintf("%d universes entered gamestate %v\n", universes, game)
			for _, roll := range die.outcomes() {
				np1 := player{num: game.p1.num, pos: game.p1.pos, score: game.p1.score}
				np2 := player{num: game.p2.num, pos: game.p2.pos, score: game.p2.score}
				if currentPlayerP1 {
					np1.pos = (np1.pos + roll) % 10
					np1.score += np1.pos + 1
					//util.Dprintf("Player %d r: %d p: %d s: %d\n", np1.num, p1roll, np1.pos, np1.score)
				} else {
					np2.pos = (np2.pos + roll) % 10
					np2.score += np2.pos + 1
					//util.Dprintf("Player %d r: %d p: %d s: %d\n", np2.num, p1roll, np2.pos, np2.score)
				}

				nUniverses := universes * uint64(die.probabilities()[roll])

				if np1.score >= 21 {
					p1wins += nUniverses
				} else if np2.score >= 21 {
					p2wins += nUniverses
				} else {
					nGS := gamestate{p1: np1, p2: np2}
					nextGames[nGS] += nUniverses
				}
			}
			util.DecreasedebugIndent()
		}
		currentGames = nextGames
		currentPlayerP1 = !currentPlayerP1
	}

	util.Dprintf("p1 won %d p2 won %d\n", p1wins, p2wins)

	if p1wins > p2wins {
		return fmt.Sprintf("%d", p1wins)
	} else {
		return fmt.Sprintf("%d", p2wins)
	}
}
