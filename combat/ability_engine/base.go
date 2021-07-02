package ability_engine

type AbilityEngine struct {
	AbilityPool AbilityPool
}

type CastAbility struct {
	Value        int
	Effect       int
	ReadableType string
	Name         string
}

func New() (abilityEngine AbilityEngine, err error) {
	abilityEngine.AbilityPool, err = prepareAbilityPool()
	if err != nil {
		return abilityEngine, err
	}

	return abilityEngine, nil
}

func (ae AbilityEngine) GenerateCastAbility() (castAbility CastAbility, err error) {
	ability := ae.AbilityPool.retrieveAbility()
	castAbility.Name = ability.Name
	castAbility.Effect = ability.Effect
	castAbility.ReadableType = ability.getReadableEffect()
	castAbility.Value = ability.generateValue()

	return castAbility, nil
}

func (c CastAbility) IsDamage() bool {
	return c.Effect == DamageAbility
}

func (c CastAbility) IsHeal() bool {
	return c.Effect == HealAbility
}
