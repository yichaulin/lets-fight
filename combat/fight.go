package combat

import (
	"math/rand"

	"lets-fight/combat/ability_engine"
)

type RoundResult struct {
	Attacker       string
	Defender       string
	Ability        ability_engine.Ability
	AttackerRestHP int
	DefenderRestHP int
}

type CombatResult struct {
	RoleA        Role
	RoleB        Role
	RoundResults []RoundResult
	Winner       string
}

const (
	fullHP = 100
)

func Fight(A string, B string) (combatResult CombatResult, err error) {
	roles, err := LoadRoles()
	if err != nil {
		return combatResult, err
	}

	combatResult = CombatResult{
		RoleA: roles[A],
		RoleB: roles[B],
	}

	err = combatResult.generateRounds()
	if err != nil {
		return combatResult, err
	}

	return combatResult, nil
}

func (combatResult *CombatResult) generateRounds() error {
	abilityEngine := ability_engine.New()
	roundResults := make([]RoundResult, 0)

	fighters := [2]Role{combatResult.RoleA, combatResult.RoleB}
	attackerIndex := rand.Intn(2)
	defenderIndex := 0
	if attackerIndex == 0 {
		defenderIndex = 1
	}

	attackerInitHP := fullHP
	defenderInitHP := fullHP
	for {
		attacker := fighters[attackerIndex]
		defender := fighters[defenderIndex]

		roundResult := RoundResult{
			Attacker:       attacker.Name,
			Defender:       defender.Name,
			AttackerRestHP: attackerInitHP,
		}

		ability, err := abilityEngine.GenerateDamage()
		roundResult.Ability = ability

		if err != nil {
			return err
		}

		err = roundResult.abilityHandler(defenderInitHP, ability)
		defenderRestHP := roundResult.DefenderRestHP
		roundResults = append(roundResults, roundResult)

		if defenderRestHP <= 0 {
			combatResult.Winner = attacker.Name
			break
		}

		attackerInitHP, defenderInitHP = defenderRestHP, attackerInitHP
	}

	combatResult.RoundResults = roundResults
	return nil
}

func (r *RoundResult) abilityHandler(defenderInitHP int, ability ability_engine.Ability) (err error) {
	if ability.IsDamage() {
		r.DefenderRestHP = defenderInitHP - ability.Value
	}

	return nil
}
