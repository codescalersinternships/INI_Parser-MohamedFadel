package internal

import (
	"fmt"
	"os"
	"strings"
)

func (p INIParser) LoadFromString(data string) (MapOfMaps, error) {
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
	p.Data = parsedData

	return p.Data, nil
}

func (p INIParser) LoadFromFile(path string) (MapOfMaps, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file %w", err)
	}
	dataToString := string(data)

	return p.LoadFromString(dataToString)
}

func (p INIParser) GetSectionNames() ([]string, error) {
	p.SectionNames = []string{}

	if len(p.Data) == 0 {
		return nil, fmt.Errorf("the map is empty")
	}

	for section := range p.Data {
		p.SectionNames = append(p.SectionNames, section)
	}

	return p.SectionNames, nil
}

func (p INIParser) GetSections() MapOfMaps {
	return p.Data
}


