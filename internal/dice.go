package internal

import (
	"math/rand"
)

// DiceResult represents the outcome of a single d20 roll in a contested action.
// The result includes both the raw roll value and contextual information
// about success level and margin of victory.
type DiceResult struct {
	Roll   int    // The final roll result (d20 + modifiers)
	Band   string // Success level: "FAIL", "OK", or "CRIT"
	Margin int    // Difference between this roll and opponent's roll (can be negative)
}

// OpposedD20 performs a contested d20 roll between two participants.
// This is the core resolution mechanic for combat, contests, and skill checks.
//
// Resolution Bands:
//   - FAIL: Margin < 2 (no significant effect)
//   - OK: Margin ≥ 2 (basic success, can trigger standard effects)
//   - CRIT: Margin ≥ 2 AND natural 20 (enhanced success, powerful effects)
//
// The margin system ensures that narrow victories (1 point) don't trigger
// major effects, requiring decisive wins for significant impact.
//
// Parameters:
//   - modA: Combat modifier for participant A (from archetype, room, cards, etc.)
//   - modB: Combat modifier for participant B
//
// Returns:
//   - DiceResult for participant A
//   - DiceResult for participant B (margin will be negative of A's margin)
//
// Example Usage:
//   // Soldier (+2) vs Mage (+1) in contested room
//   resultA, resultB := OpposedD20(2, 1)
//   if resultA.Margin >= 2 {
//       // Soldier wins decisively - can remove OS=1 enemy unit
//       if resultA.Band == "CRIT" {
//           // Natural 20 with margin - enhanced effect
//       }
//   }
func OpposedD20(modA, modB int) (DiceResult, DiceResult) {
	naturalA := rand.Intn(20) + 1    // Natural d20 roll for A
	naturalB := rand.Intn(20) + 1    // Natural d20 roll for B
	rollA := naturalA + modA         // Final modified roll for A  
	rollB := naturalB + modB         // Final modified roll for B
	margin := rollA - rollB          // Positive = A wins, negative = B wins

	return DiceResult{Roll: rollA, Band: band(rollA, rollB, naturalA), Margin: margin},
		DiceResult{Roll: rollB, Band: band(rollB, rollA, naturalB), Margin: -margin}
}

// band determines the success level based on roll comparison and natural die result.
// The band system creates three distinct outcome categories:
//
//   - FAIL: Margin < 2 - No significant advantage gained
//   - OK: Margin ≥ 2 - Standard success, basic effects trigger
//   - CRIT: Margin ≥ 2 + Natural 20 - Enhanced success, powerful effects
//
// Critical hits require both a decisive margin AND maximum die roll,
// making them rare but impactful moments.
//
// Parameters:
//   - roll: Final modified roll result
//   - opp: Opponent's final modified roll result  
//   - natural: The natural d20 result (before modifiers)
//
// Returns success band as string: "FAIL", "OK", or "CRIT"
func band(roll, opp, natural int) string {
	diff := roll - opp
	switch {
	case diff >= 2 && natural == 20:
		return "CRIT" // Decisive win with natural 20
	case diff >= 2:
		return "OK" // Decisive win
	default:
		return "FAIL" // Narrow victory or loss
	}
}

// ResolveCombat handles combat between two cards using opposed d20 resolution.
// Applies archetype bonuses, room effects, and determines outcomes based on margin.
//
// Combat Flow:
//   1. Calculate modifiers for each participant
//   2. Roll opposed d20s
//   3. Apply effects based on margin and bands
//   4. Handle unit removal, positioning, etc.
//
// Parameters:
//   - attacker: The initiating card
//   - defender: The target card
//   - room: The room where combat occurs (affects modifiers)
//
// Returns:
//   - DiceResult for attacker
//   - DiceResult for defender
//   - bool indicating if defender was eliminated
//
// Future enhancements will include:
//   - Equipment and temporary effect modifiers
//   - Environmental hazards and room-specific rules
//   - Chain reactions and triggered abilities
func ResolveCombat(attacker, defender *Card, room *Room) (DiceResult, DiceResult, bool) {
	// Calculate combat modifiers
	attackerMod := attacker.GetCombatModifier(defender, room)
	defenderMod := defender.GetCombatModifier(attacker, room)

	// Perform opposed roll
	attackResult, defendResult := OpposedD20(attackerMod, defenderMod)

	// Determine if defender is eliminated
	eliminated := false
	if attackResult.Margin >= 2 && defender.OS <= 1 {
		// Decisive victory against OS=1 unit results in elimination
		eliminated = true
	}

	return attackResult, defendResult, eliminated
}