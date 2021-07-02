package ability_engine

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Ability struct {
	Name string
	Type int
}

const (
	abilityPath = "./config/abilities.yml"
)

const (
	DamageAbility = iota
	HealAbility
	Ultimate
)

func (ability Ability) generateValue() int {
	return getIntn(20) + 1
}

func retrieveAbility(abilities []Ability) Ability {
	abilitiesLen := len(abilities)
	index := getIntn(abilitiesLen)
	return abilities[index]
}

func loadAbilities() (abilities []Ability, err error) {
	abilities = make([]Ability, 0)

	fileByte, err := os.ReadFile(abilityPath)
	if err != nil {
		return abilities, err
	}

	err = yaml.Unmarshal(fileByte, &abilities)
	if err != nil {
		return abilities, err
	}

	return abilities, nil
}
