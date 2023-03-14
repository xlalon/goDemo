package conf

import (
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/xlalon/golee/pkg/os"
)

func fromFile(cnfPath string) (*Config, error) {
	cnf := new(Config)
	*cnf = *Conf

	data, err := os.ReadFromFile(cnfPath, 10240)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(data, cnf); err != nil {
		return nil, fmt.Errorf("unmarshal YAML error: %s", err)
	}

	return cnf, nil
}
