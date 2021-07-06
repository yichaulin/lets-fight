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
	Type   int    `yaml:"type"`
}

type AbilityPool struct {
	General  []Ability
	Skill    []Ability
	Ultimate []Ability
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
	LightValue
	MidValue
	HighValue
	UltimateValue
)

const (
	GeneralAbility = iota
	SkillAbility
	UltimateAbility
)

func (ability Ability) generateValue() int {
	var baseValue, floatValue int
	switch ability.Level {
	case LowValue:
		baseValue = 1
		floatValue = GetIntn(3)
	case LightValue:
		baseValue = 5
		floatValue = GetIntn(4)
	case MidValue:
		baseValue = 10
		floatValue = GetIntn(4)
	case HighValue:
		baseValue = 15
		floatValue = GetIntn(4)
	case UltimateValue:
		baseValue = 20
		floatValue = GetIntn(10)
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

func (ability Ability) getReadableType() string {
	switch ability.Type {
	case GeneralAbility:
		return "General"
	case SkillAbility:
		return "Skill"
	case UltimateAbility:
		return "Ultimate"
	default:
		return ""
	}
}

func (abilityPool AbilityPool) retrieveAbility() Ability {
	roll := GetIntn(100)
	var targetAbilities []Ability
	if roll < 40 {
		targetAbilities = abilityPool.General
	} else if roll < 80 {
		targetAbilities = abilityPool.Skill
	} else {
		targetAbilities = abilityPool.Ultimate
	}

	abilitiesLen := len(targetAbilities)
	if abilitiesLen == 0 {
		return Ability{}
	}

	return targetAbilities[GetIntn(abilitiesLen)]

}

func prepareAbilityPool() (abilityPool AbilityPool, err error) {
	files, err := os.ReadDir(abilityDir)

	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", abilityDir, file.Name())
		fileByte, err := os.ReadFile(filePath)
		if err != nil {
			return abilityPool, err
		}

		var abs []Ability
		err = yaml.Unmarshal(fileByte, &abs)
		if err != nil {
			return abilityPool, err
		}

		for _, a := range abs {
			switch a.Type {
			case GeneralAbility:
				abilityPool.General = append(abilityPool.General, a)
			case SkillAbility:
				abilityPool.Skill = append(abilityPool.Skill, a)
			case UltimateAbility:
				abilityPool.Ultimate = append(abilityPool.Ultimate, a)
			}
		}

	}

	return abilityPool, nil
}
