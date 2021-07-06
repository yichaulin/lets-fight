package combat

import (
	"lets-fight/combat/ability_engine"
)

type RoundResult struct {
	RoundNum       int                        `json:"roundNum"`
	Attacker       string                     `json:"attacker"`
	Defender       string                     `json:"defender"`
	CastAbility    ability_engine.CastAbility `json:"castAbility"`
	AttackerRestHP int                        `json:"attackerRestHP"`
	DefenderRestHP int                        `json:"defenderRestHP"`
}

type CombatResult struct {
	Fighters     [2]string     `json:"fighters"`
	Winner       string        `json:"winner"`
	RoundsCount  int           `json:"roundCounts"`
	RoundResults []RoundResult `json:"roundResults"`
}

const (
	fullHP = 100
)

func New(fighters [2]string) (combatResult CombatResult, err error) {

	firstAttackerIndex := ability_engine.GetIntn(2)
	firstDefenderIndex := 0
	if firstAttackerIndex == 0 {
		firstDefenderIndex = 1
	}

	firstAttacker := fighters[firstAttackerIndex]
	firstDefender := fighters[firstDefenderIndex]

	combatResult = CombatResult{
		Fighters: fighters,
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
	roundNum := 1
	for {
		roundResult := RoundResult{
			Attacker:       attacker,
			Defender:       defender,
			AttackerRestHP: attackerInitHP,
			RoundNum:       roundNum,
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
		roundNum++
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
