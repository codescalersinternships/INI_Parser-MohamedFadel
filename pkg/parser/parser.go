/*
Package parser provides an INI parser with methods to load, manipulate, a
nd save configuration data from INI formatted strings and files.
*/
package parser

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
LoadFromString parses an INI formatted string and loads the data into the parser.
Returns the parsed data as a MapOfMaps or an error if the format is invalid.
*/
func (p *INIParser) LoadFromString(data string) (MapOfMaps, error) {
	lines := strings.Split(data, "\n")
	cleanLines := make([]string, 0)
	parsedData := make(MapOfMaps)
	var currentSection string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		// Remove spaces if the line is not a section header
		if !strings.HasPrefix(line, "[") {
			line = strings.ReplaceAll(line, " ", "")
		}

		cleanLines = append(cleanLines, line)

	}

	for _, line := range cleanLines {
		// Detect section headers
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = strings.Trim(line, "[]")

			parsedData[currentSection] = make(map[string]string)

		} else if currentSection != "" {
			trimmedLine := strings.Split(line, "=")

			// Check for valid key-value pair
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

/*
LoadFromFile reads an INI file from the specified path and loads the data into the parser.
Returns the parsed data as a MapOfMaps or an error if the file cannot be read or parsed.
*/
func (p *INIParser) LoadFromFile(path string) (MapOfMaps, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file %w", err)
	}
	dataToString := string(data)

	return p.LoadFromString(dataToString)
}

/*
GetSectionNames returns a list of section names from the loaded INI data.
Returns an error if no data has been loaded.
*/
func (p *INIParser) GetSectionNames() ([]string, error) {
	p.SectionNames = []string{}

	if len(p.Data) == 0 {
		return nil, fmt.Errorf("the map is empty")
	}

	for section := range p.Data {
		p.SectionNames = append(p.SectionNames, section)
	}

	return p.SectionNames, nil
}

// GetSections returns the entire parsed data as a MapOfMaps.
func (p *INIParser) GetSections() MapOfMaps {
	return p.Data
}

/*
Get retrieves the value associated with the specified key in the given section.
Returns the value or an error if the key does not exist.
*/
func (p *INIParser) Get(section, key string) (string, error) {
	value, exists := p.Data[section][key]

	if !exists {
		return "", fmt.Errorf("value does not exist")
	}

	return value, nil
}

/*
Set updates the value of a specified key in a given section.
Returns the status of the operation or an error if the section or key is not found.
*/
func (p *INIParser) Set(section, key, newValue string) (string, error) {
	state := ""
	_, exists := p.Data[section][key]

	if !exists {
		state = "not added"
		return state, fmt.Errorf("value not added, section or key not found")
	}

	p.Data[section][key] = newValue

	state = "added"
	return state, nil

}

/*
ToString converts the loaded INI data into a formatted string representation.
Returns the formatted string or an error if no data has been loaded.
*/
func (p *INIParser) ToString() (string, error) {
	if len(p.Data) == 0 {
		return "", fmt.Errorf("there is no data to convert to string")
	}

	output := ""
	for section, keyValue := range p.Data {
		output += "[" + section + "]" + "\n"

		// Sort keys
		keys := make([]string, 0, len(keyValue))
		for key := range keyValue {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			output += key + "=" + keyValue[key] + "\n"
		}
	}

	return output, nil
}

func (p *INIParser) String() string {
	data, err := p.ToString()
	if err != nil {
		return "error: no data loaded"
	}

	return data
}

/*
SaveToFile writes the loaded INI data to the specified file path.
Returns the status of the save operation or an error if the file cannot be written.
*/
func (p *INIParser) SaveToFile(path string) (string, error) {
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
