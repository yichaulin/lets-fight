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
	Fighters     [2]string
	RoundResults []RoundResult
	Winner       string
	RoundsCount  int
}

const (
	fullHP = 100
)

func New(firstAttacker string, firstDefender string) (combatResult CombatResult, err error) {
	combatResult = CombatResult{
		Fighters: [2]string{firstAttacker, firstDefender},
	}
	abilityEngine, err := ability_engine.New()

	if err != nil {
		return combatResult, err
	}

	roundResults := make([]RoundResult, 0)
	attackerInitHP := fullHP
	defenderInitHP := fullHP
	attacker := firstAttacker
	defender := firstDefender
	for {

		roundResult := RoundResult{
			Attacker:       attacker,
			Defender:       defender,
			AttackerRestHP: attackerInitHP,
		}

		castAbility, err := abilityEngine.GenerateCastAbility()
		roundResult.CastAbility = castAbility

		if err != nil {
			return combatResult, err
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
	combatResult.RoundsCount = len(roundResults)
	return combatResult, nil
}

func (r *RoundResult) castAbilityHandler(defenderInitHP int, castAbility ability_engine.CastAbility) {
	if castAbility.IsDamage() {
		r.DefenderRestHP = defenderInitHP - castAbility.Value
	} else if castAbility.IsHeal() {
		r.DefenderRestHP = defenderInitHP + castAbility.Value
	}
}
