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
It returns an error if the format is invalid.
*/
func (p *INIParser) LoadFromString(data string) error {
	lines := strings.Split(data, "\n")
	var currentSection string
	parsedData := make(MapOfMaps)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}

		// Detect section headers (e.g., [section])
		if strings.HasPrefix(line, "[") || strings.HasSuffix(line, "]") {
			// Validate if section header is well-formed (e.g., [section])
			if len(line) < 3 || strings.Count(line, "[") != 1 || strings.Count(line, "]") != 1 {
				return fmt.Errorf("%w: %s", ErrInvalidSectionHeader, line)
			}

			// Extract section name
			currentSection = strings.Trim(line, "[]")

			if _, exists := parsedData[currentSection]; exists {
				return fmt.Errorf("%w: %s", ErrSectionAlreadyExists, currentSection)
			}

			parsedData[currentSection] = make(map[string]string)

		} else if currentSection != "" {
			trimmedLine := strings.SplitN(line, "=", 2)

			// Check if it's a valid key-value pair (the key must be non-empty)
			if len(trimmedLine) != 2 || strings.TrimSpace(trimmedLine[0]) == "" {
				return fmt.Errorf("%w: %s", ErrInvalidLineFormat, line)
			}

			// Trim spaces around key and value
			key := strings.TrimSpace(trimmedLine[0])
			value := strings.TrimSpace(trimmedLine[1])

			parsedData[currentSection][key] = value
		} else {
			return fmt.Errorf("%w: %s", ErrKeyValuePairOutsideSection, line)
		}
	}

	p.data = parsedData
	return nil
}

/*
LoadFromFile reads an INI file from the specified path and loads the data into the parser.
Returns an error if the file cannot be read or parsed.
*/
func (p *INIParser) LoadFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrFileReadError, err)
	}
	dataToString := string(data)

	return p.LoadFromString(dataToString)
}

/*
GetSectionNames returns a list of section names from the loaded INI data.
If no data is present, it will return an empty list.
*/
func (p *INIParser) GetSectionNames() []string {
	sectionNames := make([]string, 0, len(p.data))

	for section := range p.data {
		sectionNames = append(sectionNames, section)
	}

	return sectionNames
}

// GetSections returns the entire parsed data as a MapOfMaps.
func (p *INIParser) GetSections() MapOfMaps {
	return p.data
}

/*
Get retrieves the value associated with the specified key in the given section.
Returns the value and a boolean indicating whether the key exists in the section.
If the key exists, the boolean will be true; otherwise, it will be false.
*/

func (p *INIParser) Get(section, key string) (string, bool) {
	value, exists := p.data[section][key]
	return value, exists
}

/*
Set updates the value of a specified key in a given section.
If the section does not exist, it will be created.
If the key does not exist in the section, it will be added.
*/
func (p *INIParser) Set(section, key, newValue string) {
	if _, sectionExists := p.data[section]; !sectionExists {
		p.data[section] = make(map[string]string)
	}

	p.data[section][key] = newValue
}

/*
String implements the Stringer interface for INIParser.
It converts the loaded INI data into a formatted string representation.
*/
func (p *INIParser) String() string {
	if len(p.data) == 0 {
		return "there is no data to convert to string"
	}

	var builder strings.Builder
	for section, keyValue := range p.data {
		builder.WriteString(fmt.Sprintf("[%s]\n", section))

		keys := make([]string, 0, len(keyValue))
		for key := range keyValue {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			builder.WriteString(fmt.Sprintf("%s=%s\n", key, keyValue[key]))
		}
	}

	return builder.String()
}

/*
SaveToFile writes the loaded INI data to the specified file path.
Returns an error if the file cannot be written.
*/
func (p *INIParser) SaveToFile(path string) error {
	data := p.String()

	if err := os.WriteFile(path, []byte(data), 0644); err != nil {
		return fmt.Errorf("%w: %v", ErrFileWriteError, err)
	}

	return nil
}
