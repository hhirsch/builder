package main

import (
	"strings"
)

type Package struct{}

func (this *Package) GetDpkgData(data string) string {
	lines := strings.Split(data, "\n")

	metadata := make(map[string]string)

	for _, line := range lines {
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			metadata[key] = value
		}
	}

	//return metadata
	return "fixme"
}
