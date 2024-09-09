package internal

import (
	"fmt"
	"os"
	"strings"
)

func LoadFromString(data string) (MapOfMaps, error) {
	lines := strings.Split(data, "\n")
	cleanLines := make([]string, 0)
	parsedData := make(MapOfMaps)
	var currentSection string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		if !strings.HasPrefix(line, "[") {
			line = strings.ReplaceAll(line, " ", "")
		}

		cleanLines = append(cleanLines, line)

	}

	for _, line := range cleanLines {
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = strings.Trim(line, "[]")

			parsedData[currentSection] = make(map[string]string)

		} else if currentSection != "" {
			trimmedLine := strings.Split(line, "=")

			if len(trimmedLine) != 2 || trimmedLine[0] == "" || trimmedLine[1] == "" {
				return nil, fmt.Errorf("invalid line format")
			}

			key := trimmedLine[0]
			value := trimmedLine[1]

			parsedData[currentSection][key] = value

		} else {
			return nil, fmt.Errorf("key-value pair found outside of a section")
		}

	}

	return parsedData, nil
}

func LoadFromFile(path string) (MapOfMaps, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file")
	}
	dataToString := string(data)

	return LoadFromString(dataToString)
}

func GetSectionNames(data MapOfMaps) ([]string, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("the map is empty")
	}

	sectionNames := make([]string, 0)
	for section := range data {
		sectionNames = append(sectionNames, section)
	}

	return sectionNames, nil
}

func GetSections(data MapOfMaps) MapOfMaps {
	return data
}
