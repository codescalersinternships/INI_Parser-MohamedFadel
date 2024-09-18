package parser

import (
	"errors"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadFromString(t *testing.T) {
	t.Run("valid ini data", func(t *testing.T) {
		p := INIParser{}
		err := p.LoadFromString(`
			; last modified 1 April 2001 by John Doe
			[owner]
			name = John Doe
			organization = Acme Widgets Inc.

			[database]
			# use IP address in case network name resolution is not working
			server = 192.0.2.62
			port = 143
			file = payroll.dat`)

		want := MapOfMaps{
			"owner":    {"name": "John Doe", "organization": "Acme Widgets Inc."},
			"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
		}

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(p.data, want) {
			t.Errorf("got %v, want %v", p.data, want)
		}
	})

	t.Run("invalid section header", func(t *testing.T) {
		p := INIParser{}
		err := p.LoadFromString(`
			; last modified 1 April 2001 by John Doe
			[owner]
			name = John Doe

			[database
			server = 192.0.2.62
			port = 143`)

		if !errors.Is(err, ErrInvalidSectionHeader) {
			t.Errorf("expected error %v, got: %v", ErrInvalidSectionHeader, err)
		}
	})

	t.Run("section already exists", func(t *testing.T) {
		p := INIParser{}
		err := p.LoadFromString(`
			; last modified 1 April 2001 by John Doe
			[owner]
			name = John Doe

			[owner]
			organization = Acme Widgets Inc.`)

		if !errors.Is(err, ErrSectionAlreadyExists) {
			t.Errorf("expected error %v, got: %v", ErrSectionAlreadyExists, err)
		}
	})

	t.Run("invalid line format", func(t *testing.T) {
		p := INIParser{}
		err := p.LoadFromString(`
			; last modified 1 April 2001 by John Doe
			[owner]
			name = John Doe
			organization = Acme Widgets Inc.

			[database]
			# use IP address in case network name resolution is not working
			server = 192.0.2.62
			port = 143
			file`)

		if !errors.Is(err, ErrInvalidLineFormat) {
			t.Errorf("expected error %v, got: %v", ErrInvalidLineFormat, err)
		}
	})

	t.Run("key-value pair outside section", func(t *testing.T) {
		p := INIParser{}
		err := p.LoadFromString(`
			; last modified 1 April 2001 by John Doe

			name = John Doe
			organization = Acme Widgets Inc.

			[database]
			# use IP address in case network name resolution is not working
			server = 192.0.2.62
			port = 143
			file = payroll.dat`)

		if !errors.Is(err, ErrKeyValuePairOutsideSection) {
			t.Errorf("expected error %v, got: %v", ErrKeyValuePairOutsideSection, err)
		}
	})

}

func TestLoadFromFile(t *testing.T) {
	t.Run("valid file path", func(t *testing.T) {
		p := INIParser{}
		err := p.LoadFromFile("../../testdata/INI.txt")

		want := MapOfMaps{
			"owner":    {"name": "John Doe", "organization": "Acme Widgets Inc."},
			"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
		}

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !reflect.DeepEqual(p.data, want) {
			t.Errorf("got %v, want %v", p.data, want)
		}
	})

	t.Run("invalid file path", func(t *testing.T) {
		p := INIParser{}
		err := p.LoadFromFile(t.TempDir())

		if !errors.Is(err, ErrFileReadError) {
			t.Errorf("expected error %v, got: %v", ErrFileReadError, err)
		}
	})
}

func TestGetSectionNames(t *testing.T) {
	t.Run("non-empty map", func(t *testing.T) {
		p := INIParser{
			data: MapOfMaps{
				"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}
		got := p.GetSectionNames()
		want := []string{"owner", "database"}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		p := INIParser{
			data: MapOfMaps{},
		}
		got := p.GetSectionNames()

		if len(got) != 0 {
			t.Errorf("expected empty list, got: %v", got)
		}
	})
}

func TestGetSections(t *testing.T) {
	p := INIParser{
		data: MapOfMaps{
			"owner":    {"name": "John Doe", "organization": "Acme Widgets Inc."},
			"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
		},
	}
	got := p.GetSections()

	want := MapOfMaps{
		"owner":    {"name": "John Doe", "organization": "Acme Widgets Inc."},
		"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v want: %v", got, want)
	}

}

func TestGet(t *testing.T) {
	t.Run("value exists", func(t *testing.T) {
		p := INIParser{
			data: MapOfMaps{
				"owner":    {"name": "John Doe", "organization": "Acme Widgets Inc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}
		got, exists := p.Get("owner", "name")
		want := "John Doe"

		if !exists {
			t.Errorf("expected key to exist, but it does not")
		}

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("value does not exist", func(t *testing.T) {
		p := INIParser{
			data: MapOfMaps{
				"owner":    {"name": "John Doe", "organization": "Acme Widgets Inc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}
		_, exists := p.Get("owner", "age")

		if exists {
			t.Errorf("expected key not to exist, but it does")
		}
	})
}

func TestSet(t *testing.T) {
	t.Run("section and key do exist", func(t *testing.T) {
		p := INIParser{
			data: MapOfMaps{
				"owner":    {"name": "John Doe", "organization": "Acme Widgets Inc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}

		p.Set("owner", "name", "Walter White")
		got := p.data["owner"]["name"]
		want := "Walter White"

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("section does not exist", func(t *testing.T) {
		p := INIParser{
			data: MapOfMaps{
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}

		p.Set("owner", "name", "Walter White")
		got := p.data["owner"]["name"]
		want := "Walter White"

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("key does not exist in existing section", func(t *testing.T) {
		p := INIParser{
			data: MapOfMaps{
				"database": {"server": "192.0.2.62", "port": "143"},
			},
		}

		p.Set("database", "file", "payroll.dat")
		got := p.data["database"]["file"]
		want := "payroll.dat"

		if got != want {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}

func TestString(t *testing.T) {
	t.Run("non-empty map", func(t *testing.T) {
		p := INIParser{
			data: MapOfMaps{
				"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}
		got := p.String()
		want := `[owner]
name=JohnDoe
organization=AcmeWidgetsInc.
[database]
file=payroll.dat
port=143
server=192.0.2.62
`

		if got != want {
			t.Errorf("got: %v\nwant: %v", got, want)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		p := INIParser{}
		got := p.String()
		want := "there is no data to convert to string"

		if got != want {
			t.Errorf("got: %v\nwant: %v", got, want)
		}
	})
}

func TestSaveToFile(t *testing.T) {
	t.Run("valid file path", func(t *testing.T) {
		p := INIParser{
			data: MapOfMaps{
				"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}

		testFilePath := filepath.Join(t.TempDir(), "newINI.txt")
		err := p.SaveToFile(testFilePath)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

	})

	t.Run("invalid file path", func(t *testing.T) {
		p := INIParser{
			data: MapOfMaps{
				"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}

		err := p.SaveToFile(t.TempDir())
		if !errors.Is(err, ErrFileWriteError) {
			t.Errorf("expected error %v, got: %v", ErrFileWriteError, err)
		}
	})
}
