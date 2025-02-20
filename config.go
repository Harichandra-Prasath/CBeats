package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Dirs []string
}

func ParseConfig(path string) (*Config, error) {

	parts := strings.Split(path, ".")

	if parts[len(parts)-1] != "conf" {
		return nil, fmt.Errorf("parsing config: invalid config file; expected .conf")
	}

	fp, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("parsing config: %s", err)
	}

	data, err := io.ReadAll(fp)
	if err != nil {
		return nil, fmt.Errorf("parsing config: %s", err)
	}

	_data := string(data)

	// Remove all the blank lines
	lines := strings.Split(_data, "\n")

	var valids []string

	for _, line := range lines {

		// strip whitespaces
		line = strings.TrimSpace(line)

		// Skip Comment Lines
		if strings.HasPrefix(line, "#") || len(line) == 0 {
			continue
		}

		valids = append(valids, line)
	}

	var config Config

	for _, line := range valids {

		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			return nil, fmt.Errorf("parsing config: invalid line structure ")
		}

		if ok := isIncluded(parts[0], ALLOWED_DIRECTIVES); !ok {
			return nil, fmt.Errorf("parsing config: invalid directive mentioned %s", parts[0])
		}

		switch parts[0] {
		case "DIRS":
			dirs := strings.Split(parts[1], ",")
			config.Dirs = dirs

		case "PORT":
			SNK_PORT = parts[1]
		case "READ_TIME":
			t, _ := strconv.Atoi(parts[1])
			READ_TIME = t
		case "DOCKER":
			if parts[1] == "True" || parts[1] == "true" {
				CONTAINERS = true
			}
		case "DOCKER_SOCK":
			DOCKER_SOCK = parts[1]
		case "CONTAINERS_LIST":

			CONTAINERS_LIST = parts[1]

		}

	}

	return &config, nil

}
