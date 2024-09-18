package parser

/*
MapOfMaps represents the structure of the parsed INI data,
where the outer map's keys are section names, and the inner maps store key-value pairs for each section.
*/
type MapOfMaps map[string]map[string]string

// INIParser is the main structure that holds the parsed INI data.
type INIParser struct {
	// Data stores the parsed INI data as a MapOfMaps.
	data MapOfMaps
}
