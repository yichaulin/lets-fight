package combat

import (
	"lets-fight/combat/ability_engine"
)

type RoundResult struct {
	Attacker       string
	Defender       string
	CastAbility    ability_engine.CastAbility
	AttackerRestHP int
	DefenderRestHP int
}

type CombatResult struct {
	RedSide      string
	BlueSide     string
	RoundResults []RoundResult
	Winner       string
}

const (
	fullHP = 100
)

func Fight(blue string, red string) (combatResult CombatResult, err error) {
	combatResult = CombatResult{
		BlueSide: blue,
		RedSide:  red,
	}

	err = combatResult.generateRounds()
	if err != nil {
		return combatResult, err
	}

	return combatResult, nil
}

func (combatResult *CombatResult) generateRounds() error {
	abilityEngine, err := ability_engine.New()
	if err != nil {
		return err
	}

	fighters := []string{combatResult.RedSide, combatResult.BlueSide}
	roundResults := make([]RoundResult, 0)
	attackerIndex, defenderIndex := 0, 1

	attackerInitHP := fullHP
	defenderInitHP := fullHP
	attacker := fighters[attackerIndex]
	defender := fighters[defenderIndex]
	for {

		roundResult := RoundResult{
			Attacker:       attacker,
			Defender:       defender,
			AttackerRestHP: attackerInitHP,
		}

		castAbility, err := abilityEngine.GenerateCastAbility()
		roundResult.CastAbility = castAbility

		if err != nil {
			return err
		}

		roundResult.castAbilityHandler(defenderInitHP, castAbility)
		roundResults = append(roundResults, roundResult)

		if roundResult.DefenderRestHP <= 0 {
			combatResult.Winner = attacker
			break
		}

		attackerInitHP, defenderInitHP = roundResult.DefenderRestHP, attackerInitHP
		attacker, defender = defender, attacker
	}

	combatResult.RoundResults = roundResults
	return nil
}

func (r *RoundResult) castAbilityHandler(defenderInitHP int, castAbility ability_engine.CastAbility) {
	if castAbility.IsDamage() {
		r.DefenderRestHP = defenderInitHP - castAbility.Value
	} else if castAbility.IsHeal() {
		r.DefenderRestHP = defenderInitHP + castAbility.Value
	}
}
