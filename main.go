package main

import (
	"fmt"
	"lets-fight/combat"
)

func main() {
	roles, err := combat.LoadRoles()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v", roles)
}
