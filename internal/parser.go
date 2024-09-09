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

func (p INIParser) Get(section, key string) (string, error) {
	value, exists := p.Data[section][key]

	if !exists {
		return "", fmt.Errorf("value does not exist")
	}

	return value, nil
}

func (p INIParser) Set(section, key, newValue string) (string, error) {
	state := ""
	_, exists := p.Data[section][key]

	if !exists {
		state = "not added"
		return state, fmt.Errorf("value not added, section or key not found")
	}

	state = "added"
	return state, nil

}

func (p INIParser) ToString() (string, error) {
	if len(p.Data) == 0 {
		return "", fmt.Errorf("there is no data to convert to string")
	}

	output := ""
	for section, keyValue := range p.Data {
		output += "[" + section + "]" + "\n"

		for key, value := range keyValue {
			output += key + "=" + value + "\n"
		}
	}

	return output, nil
}

func (p INIParser) SaveToFile(path string) (string, error) {
	state := "not saved"
	if len(p.Data) == 0 {
		return state, fmt.Errorf("there is no data to save to file")
	}

	data, _ := p.ToString()
	dataToBytes := []byte(data)
	if err := os.WriteFile(path, dataToBytes, 0644); err != nil {
		return state, fmt.Errorf("error writing to file %w", err)
	}

	state = "saved"
	return state, nil

}
