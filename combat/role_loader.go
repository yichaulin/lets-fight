package combat

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Role struct {
	Name      string   `yaml:"name"`
	Abilities []string `yaml:"abilities"`
	Ultimate  string   `yaml:"ultimate"`
}

const (
	rolesDir = "./roles"
)

func LoadRoles() (map[string]Role, error) {
	roles := make(map[string]Role, 0)

	files, err := os.ReadDir(rolesDir)
	if err != nil {
		return roles, err
	}

	for _, file := range files {
		filePath := fmt.Sprintf("%s/%s", rolesDir, file.Name())
		fileByte, err := os.ReadFile(filePath)
		if err != nil {
			return roles, err
		}

		var r Role
		err = yaml.Unmarshal(fileByte, &r)
		if err != nil {
			return roles, err
		}
		roles[r.Name] = r
	}

	return roles, nil
}
