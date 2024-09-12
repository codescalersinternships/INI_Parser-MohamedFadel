package main

func main() {
	/* just testing the functionality

		// Initialize INIParser instances
		p1 := parser.INIParser{}
		p2 := parser.INIParser{}

		// Load INI data from a string
		iniData := `; last modified 1 April 2001 by John Doe
	[owner]
	name = John Doe
	organization = Acme Widgets Inc.

	[database]
	# use IP address in case network name resolution is not working
	server = 192.0.2.62
	port = 143
	file = payroll.dat
	`
		_, err := p1.LoadFromString(iniData)
		if err != nil {
			log.Fatal(err)
		}

		// Load INI data from a file
		filePath := filepath.Join("pkg", "parser", "INI.txt")
		_, err = p2.LoadFromFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		// Get section names from the first parser
		sections, err := p1.GetSectionNames()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Section Names:")
		for _, section := range sections {
			fmt.Println(section)
		}

		// Print all sections and their key-value pairs from the second parser
		data := p2.GetSections()
		fmt.Println("\nINI Data:")
		for section, keyValue := range data {
			fmt.Println("[", section, "]")
			for key, value := range keyValue {
				fmt.Println(key, "=", value)
			}
		}

		// Update a value in the first parser
		_, err = p1.Set("owner", "name", "WalterWhite")
		if err != nil {
			log.Fatal(err)
		}

		// Retrieve and print the updated value
		value, err := p1.Get("owner", "name")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nUpdated Value:")
		fmt.Println(value)

		// Convert the INI data to a string and print it
		stringData, err := p1.ToString()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("\nINI Data as String:")
		fmt.Println(stringData)

		// Save the updated INI data to a new file
		newFilePath := filepath.Join("pkg", "parser", "newINI.txt")
		_, err = p1.SaveToFile(newFilePath)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("\nINI data saved to", newFilePath)

	*/
}
