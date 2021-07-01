package main

import (
	"encoding/json"
	"fmt"
	"lets-fight/combat"
)

func main() {
	combatResult, err := combat.Fight("saber", "archer")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", combatResult)
	PrettyPrint(combatResult)
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}
