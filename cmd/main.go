package main

import (
	"fmt"
	"log"

	"github.com/codescalersinternships/INI_Parser-MohamedFadel/internal"
)

func main() {
	parser := internal.INIParser{}
	parser2 := internal.INIParser{}

	parser.LoadFromString(`; last modified 1 April 2001 by John Doe
[owner]
name = John Doe
organization = Acme Widgets Inc.

[database]
# use IP address in case network name resolution is not working
server = 192.0.2.62
port = 143
file = payroll.dat
`)

	parser2.LoadFromFile("../../../../INI.txt")

	sections, err := parser.GetSectionNames()
	if err != nil {
		log.Fatal(err)
	}

	for _, section := range sections {
		fmt.Println(section)
	}

	data := parser2.GetSections()
	for section, keyValue := range data {
		fmt.Println("[", section, "]")

		for key, value := range keyValue {
			fmt.Println(key, "=", value)
		}
	}

	parser.Set("owner", "name", "WalterWhite")
	value, err := parser.Get("owner", "name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)

	stringData, err := parser.ToString()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stringData)

	_, err = parser2.SaveToFile("../../../../newINI.txt")
	if err != nil {
		log.Fatal(err)
	}

}
