package internal

type MapOfMaps map[string]map[string]string

type INIParser struct {
	Data MapOfMaps
	SectionNames []string
}