# INI_Parser-MohamedFadel

Welcome to the INI_Parser-MohamedFadel repository! This project implements an INI file parser in Go, with functionality to load, modify, and save INI files. Each method in the `INIParser`  provides specific functionality related to INI file manipulation.

## Overview

The `INI_Parser` project provides functionality to parse, modify, and save INI files. It supports loading INI data from strings or files, manipulating sections and key-value pairs, and saving changes back to files.

## List of Methods

- **`LoadFromString(data string) (MapOfMaps, error)`**: Loads INI data from a string and parses it into `MapOfMaps`.
  - **Parameters**: 
    - `data` - INI formatted string.
  - **Returns**: 
    - `MapOfMaps` - Parsed data.
    - `error` - Error if parsing fails.

- **`LoadFromFile(path string) (MapOfMaps, error)`**: Loads INI data from a file and parses it.
  - **Parameters**: 
    - `path` - Path to the INI file.
  - **Returns**: 
    - `MapOfMaps` - Parsed data.
    - `error` - Error if file reading or parsing fails.

- **`GetSectionNames() ([]string, error)`**: Retrieves the names of all sections in the parsed data.
  - **Returns**: 
    - `[]string` - List of section names.
    - `error` - Error if the map is empty.

- **`GetSections() MapOfMaps`**: Retrieves all sections and their key-value pairs.
  - **Returns**: 
    - `MapOfMaps` - Sections and their key-value pairs.

- **`Get(section, key string) (string, error)`**: Retrieves the value for a given section and key.
  - **Parameters**: 
    - `section` - Section name.
    - `key` - Key name.
  - **Returns**: 
    - `string` - Value for the key.
    - `error` - Error if the key or section does not exist.

- **`Set(section, key, newValue string) (string, error)`**: Sets a new value for a given section and key.
  - **Parameters**: 
    - `section` - Section name.
    - `key` - Key name.
    - `newValue` - New value to set.
  - **Returns**: 
    - `string` - Result message ("added" or "not added").
    - `error` - Error if the section or key does not exist.

- **`ToString() (string, error)`**: Converts the parsed INI data back to a string format.
  - **Returns**: 
    - `string` - INI formatted string.
    - `error` - Error if the data cannot be converted.

- **`SaveToFile(path string) (string, error)`**: Saves the current INI data to a file.
  - **Parameters**: 
    - `path` - Path to the output file.
  - **Returns**: 
    - `string` - Result message ("saved" or "not saved").
    - `error` - Error if writing to the file fails.

## Example Usage

### Load INI Data from a String

```go
p := INIParser{}
data, err := p.LoadFromString(`
; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.

[database]
server = 192.0.2.62
port = 143
file = payroll.dat
`)
```

### Load INI Data from a File
```go
p := INIParser{}
data, err := p.LoadFromFile("path/to/file.ini")
```

### Get Section Names
```go
sectionNames, err := p.GetSectionNames()
```

### Get Sections
```go
sections := p.GetSections()
```

### Get a Value
```go
value, err := p.Get("owner", "name")
```

### Set a Value
```go
result, err := p.Set("owner", "name", "WalterWhite")
```

### Convert to String
```go
iniString, err := p.ToString()
```

### Save to File
```go
result, err := p.SaveToFile("path/to/output.ini")
```

## Testing
The package includes tests for each method to ensure correct functionality:

-   **TestLoadFromString**: Tests for valid and invalid INI data.
    
    -   **valid ini data**: Validates successful loading and parsing of a well-formed INI string.
    -   **invalid line format**: Tests error handling for improperly formatted INI lines.
    -   **key-value pair outside section**: Checks error handling for key-value pairs outside of any section.
-   **TestLoadFromFile**: Tests for valid and invalid file paths.
    
    -   **valid file path**: Validates successful loading of data from a correctly formatted INI file.
    -   **invalid file path**: Tests error handling for non-existent or inaccessible files.
-   **TestGetSectionNames**: Tests for non-empty and empty data.
    
    -   **non empty map**: Verifies retrieval of section names from non-empty data.
    -   **empty map**: Checks error handling when no data is available.
-   **TestGetSections**: Tests for retrieving sections.
    
    -   **retrieving sections**: Validates correct retrieval of all sections and their key-value pairs.
-   **TestGet**: Tests for existing and non-existing values.
    
    -   **value exists**: Ensures correct retrieval of existing values.
    -   **value does not exist**: Tests error handling for non-existent sections or keys.
-   **TestSet**: Tests for updating existing and non-existing sections/keys.
    
    -   **section and key do exist**: Validates successful update of existing values.
    -   **section or key does not exist**: Checks error handling for attempts to update non-existent sections or keys.
-   **TestToString**: Tests for non-empty and empty data.
    
    -   **non empty map**: Verifies correct conversion of data to INI formatted string.
    -   **empty map**: Tests error handling when no data is available to convert.
-   **TestSaveToFile**: Tests for valid, empty, and invalid file paths.
    
    -   **non empty map and valid file path**: Ensures successful saving of data to a valid file.
    -   **empty map**: Checks error handling for attempts to save empty data.
    -   **invalid file path**: Tests error handling for invalid file paths or permissions issues.