
#  INI_Parser-MohamedFadel

  

Welcome to the INI_Parser-MohamedFadel repository! This project implements an INI file parser in Go, with functionality to load, modify, and save INI files. Each method in the `INIParser` provides specific functionality related to INI file manipulation.

  

##  Overview

  

The `INI_Parser` project provides functionality to parse, modify, and save INI files. It supports loading INI data from strings or files, manipulating sections and key-value pairs, and saving changes back to files.

  

##  Installation

  

To install the package, use:

  

```bash

go  get  https://github.com/codescalersinternships/INI_Parser-MohamedFadel/tree/development

```

  

##  Example Usage

  

###  Load INI Data from a String

  

```go

p  := parser.INIParser{}

err  :=  p.LoadFromString(`

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

  

###  Load INI Data from a File

```go

p  := parser.INIParser{}

err  :=  p.LoadFromFile("path/to/file.ini")

```

  

###  Get Section Names

```go

sectionNames :=  p.GetSectionNames()

```

  

###  Get Sections

```go

sections  :=  p.GetSections()

```

  

###  Get a Value

```go

value, _  :=  p.Get("owner", "name")

```

  

###  Set a Value

```go

p.Set("owner", "name", "WalterWhite")

```

  

###  Convert to String

```go

iniString :=  p.String()

```

  

###  Save to File

```go

err  :=  p.SaveToFile("path/to/output.ini")

```