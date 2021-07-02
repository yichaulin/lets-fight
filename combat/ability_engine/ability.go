package ability_engine

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Ability struct {
	Name   string `yaml:"name"`
	Effect int    `yaml:"effect"`
	Level  int    `yaml:"level"`
}

type Abilities []Ability

type AbilityPool struct {
	General  Abilities
	Skill    Abilities
	Ultimate Abilities
}

const (
	abilityDir = "./config/abilities"
)

const (
	DamageAbility = iota
	HealAbility
)

const (
	LowValue = iota
	MidValue
	HighValue
	UltimateValue
)

func (ability Ability) generateValue() int {
	var baseValue, floatValue int
	switch ability.Level {
	case LowValue:
		baseValue = 1
		floatValue = getIntn(4)
	case MidValue:
		baseValue = 5
		floatValue = getIntn(5)
	case HighValue:
		baseValue = 10
		floatValue = getIntn(10)
	case UltimateValue:
		baseValue = 20
		floatValue = getIntn(15)
	}

	return baseValue + floatValue
}

func (ability Ability) getReadableEffect() string {
	switch ability.Effect {
	case DamageAbility:
		return "Damage"
	case HealAbility:
		return "Heal"
	default:
		return ""
	}
}

func (abilityPool AbilityPool) retrieveAbility() Ability {
	roll := getIntn(100)
	if roll < 60 {
		return abilityPool.General.pickOne()
	} else if roll < 90 {
		return abilityPool.Skill.pickOne()
	} else {
		return abilityPool.Ultimate.pickOne()
	}
}

func (abilities Abilities) pickOne() (ability Ability) {
	abilitiesLen := len(abilities)
	if abilitiesLen == 0 {
		return ability
	}

	return abilities[getIntn(abilitiesLen)]
}

func prepareAbilityPool() (abilityPool AbilityPool, err error) {
	abilityTypes := map[string]*Abilities{
		"general":  &abilityPool.General,
		"skill":    &abilityPool.Skill,
		"ultimate": &abilityPool.Ultimate,
	}

	for key, ele := range abilityTypes {
		filePath := fmt.Sprintf("%s/%s.yml", abilityDir, key)

		fileByte, err := os.ReadFile(filePath)
		if err != nil {
			return abilityPool, err
		}

		err = yaml.Unmarshal(fileByte, &ele)
		if err != nil {
			return abilityPool, err
		}

	}

	return abilityPool, nil
}
