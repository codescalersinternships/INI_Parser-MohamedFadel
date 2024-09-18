package parser

import "errors"

var (
	// ErrInvalidLineFormat is returned when a line in the INI file has an invalid format
	ErrInvalidLineFormat = errors.New("invalid line format")

	// ErrKeyValuePairOutsideSection is returned when a key-value pair is found outside of a section
	ErrKeyValuePairOutsideSection = errors.New("key-value pair found outside of a section")

	// ErrInvalidSectionHeader is returned when a section header is invalid
	ErrInvalidSectionHeader = errors.New("invalid section header")

	// ErrSectionAlreadyExists is returned when attempting to create a section that already exists
	ErrSectionAlreadyExists = errors.New("section already exists")

	// ErrFileReadError is returned when there's an error reading the INI file
	ErrFileReadError = errors.New("error reading file")

	// ErrFileWriteError is returned when there's an error writing to the INI file
	ErrFileWriteError = errors.New("error writing to file")
)
