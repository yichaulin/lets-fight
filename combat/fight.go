package combat

import "math/rand"

type RoundResult struct {
	Attacker       string
	Defender       string
	Damage         int
	Ability        string
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

		err := roundResult.generateDamage(attacker.Abilities)

		if err != nil {
			return err
		}

		defenderRestHP := defenderInitHP - roundResult.Damage
		roundResult.DefenderRestHP = defenderRestHP
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

func (c *RoundResult) generateDamage(abilities []string) (err error) {
	damage := rand.Intn(15) + 15
	ability := abilities[0]

	c.Damage = damage
	c.Ability = ability
	return err
}
