package internal

/*
MapOfMaps represents the structure of the parsed INI data,
where the outer map's keys are section names, and the inner maps store key-value pairs for each section.
*/
type MapOfMaps map[string]map[string]string

// INIParser is the main structure that holds the parsed INI data and section names.
type INIParser struct {
	// Data stores the parsed INI data as a MapOfMaps.
	Data MapOfMaps

	// SectionNames is a slice containing the names of the sections parsed from the INI data.
	SectionNames []string
}
