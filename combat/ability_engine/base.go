package ability_engine

type AbilityEngine struct {
	Abilities []Ability
}

type CastAbility struct {
	Value int
	Type  int
	Name  string
}

func New() (abilityEngine AbilityEngine, err error) {
	abilities, err := loadAbilities()

	if err != nil {
		return abilityEngine, err
	}

	abilityEngine.Abilities = abilities

	return abilityEngine, nil
}

func (ae AbilityEngine) GenerateCastAbility() (castAbility CastAbility, err error) {
	ability := retrieveAbility(ae.Abilities)
	castAbility.Name = ability.Name
	castAbility.Type = ability.Type
	castAbility.Value = ability.generateValue()

	return castAbility, nil
}

func (c CastAbility) IsDamage() bool {
	return c.Type == DamageAbility
}

func (c CastAbility) IsHeal() bool {
	return c.Type == HealAbility
}
