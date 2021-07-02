package ability_engine

import (
	"math/rand"
	"time"
)

type AbilityEngine struct {
	Source rand.Source
}

type Ability struct {
	Value       int
	AbilityType int
	AbilityName string
	IsUltimate  bool
}

const (
	DamageAbility = iota
	HealAbility
)

func New() AbilityEngine {
	ts := time.Now().Unix()

	return AbilityEngine{
		Source: rand.NewSource(ts),
	}
}

func (ae AbilityEngine) GenerateDamage() (ability Ability, err error) {
	source := ae.Source
	isUltimate := isUltimateHappened()
	value := getRandIntn(source, 15) + 15

	ability.IsUltimate = isUltimate
	ability.Value = value
	ability.AbilityType = DamageAbility

	return ability, nil
}

func (a Ability) IsDamage() bool {
	return a.AbilityType == DamageAbility
}

func isUltimateHappened() bool {
	return false
}

func getRandIntn(source rand.Source, n int) int {
	random := rand.New(source)

	return random.Intn(n)
}
