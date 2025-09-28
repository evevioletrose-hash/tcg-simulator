package internal

import (
	"math/rand"
)

type DiceResult struct {
	Roll   int
	Band   string // "FAIL", "OK", "CRIT"
	Margin int
}

func OpposedD20(modA, modB int) (DiceResult, DiceResult) {
	rollA := rand.Intn(20) + 1 + modA
	rollB := rand.Intn(20) + 1 + modB
	margin := rollA - rollB
	return DiceResult{Roll: rollA, Band: band(rollA, rollB), Margin: margin},
		DiceResult{Roll: rollB, Band: band(rollB, rollA), Margin: -margin}
}

func band(roll, opp int) string {
	diff := roll - opp
	switch {
	case diff >= 2 && roll == 20:
		return "CRIT"
	case diff >= 2:
		return "OK"
	default:
		return "FAIL"
	}
}