package ability_engine

type AbilityEngine struct {
	AbilityPool AbilityPool
}

type CastAbility struct {
	Name           string `json:"name"`
	Effect         int    `json:"-"`
	ReadableEffect string `json:"effect"`
	Type           int    `json:"-"`
	ReadableType   string `json:"type"`
	Value          int    `json:"value"`
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
	castAbility.ReadableEffect = ability.getReadableEffect()
	castAbility.Type = ability.Type
	castAbility.ReadableType = ability.getReadableType()
	castAbility.Value = ability.generateValue()

	return castAbility, nil
}

func (c CastAbility) IsDamage() bool {
	return c.Effect == DamageAbility
}

func (c CastAbility) IsHeal() bool {
	return c.Effect == HealAbility
}
