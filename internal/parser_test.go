package internal

import (
	"reflect"
	"testing"
)

func TestLoadFromString(t *testing.T) {
	t.Run("valid ini data", func(t *testing.T) {
		p := INIParser{}
		got, err := p.LoadFromString(
			`; last modified 1 April 2001 by John Doe
			[owner]
			name = John Doe
			organization = Acme Widgets Inc.

			[database]
			# use IP address in case network name resolution is not working
			server = 192.0.2.62
			port = 143
			file = payroll.dat`)

		want := MapOfMaps{
			"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
			"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
		}

		if err != nil {
			t.Fail()
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("invalid line format", func(t *testing.T) {
		p := INIParser{}
		_, err := p.LoadFromString(
			`; last modified 1 April 2001 by John Doe
			[owner]
			name =
			organization = Acme Widgets Inc.

			[database]
			# use IP address in case network name resolution is not working
			server = 192.0.2.62
			port = 143
			file = payroll.dat`)

		if err == nil {
			t.Errorf("expected error 'invalid line format', got: %v", err)
		}
	})

	t.Run("key-value pain outside section", func(t *testing.T) {
		p := INIParser{}
		_, err := p.LoadFromString(
			`; last modified 1 April 2001 by John Doe

			name = John Doe
			organization = Acme Widgets Inc.

			[database]
			# use IP address in case network name resolution is not working
			server = 192.0.2.62
			port = 143
			file = payroll.dat`)

		if err == nil {
			t.Errorf("expected error 'key-value pair found outside of a section', got: %v", err)
		}
	})

}

func TestLoadFromFile(t *testing.T) {
	t.Run("valid file path", func(t *testing.T) {
		p := INIParser{}
		got, err := p.LoadFromFile("./INI.txt")
		want := MapOfMaps{
			"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
			"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
		}

		if err != nil {
			t.Fail()
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("invalid file path", func(t *testing.T) {
		p := INIParser{}
		_, err := p.LoadFromFile("../../../INI.txt")
		if err == nil {
			t.Errorf("expected error 'error reading file', got: %v", err)
		}

	})
}

func TestGetSectionNames(t *testing.T) {
	t.Run("non empty map", func(t *testing.T) {
		p := INIParser{
			Data: MapOfMaps{
				"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}
		got, err := p.GetSectionNames()
		want := []string{"owner", "database"}

		if err != nil {
			t.Fail()
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v want: %v", got, want)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		p := INIParser{}
		_, err := p.GetSectionNames()

		if err == nil {
			t.Errorf("expected error 'the map is empty', got: %v", err)
		}
	})
}

func TestGetSections(t *testing.T) {
	p := INIParser{
		Data: MapOfMaps{
			"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
			"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
		},
	}
	got := p.GetSections()

	want := MapOfMaps{
		"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
		"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v want: %v", got, want)
	}

}

func TestGet(t *testing.T) {
	t.Run("value exists", func(t *testing.T) {
		p := INIParser{
			Data: MapOfMaps{
				"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}
		got, err := p.Get("owner", "name")
		want := "JohnDoe"

		if err != nil {
			t.Fail()
		}

		if got != want {
			t.Errorf("got: %v want: %v", got, want)
		}
	})

	t.Run("value does not exist", func(t *testing.T) {
		p := INIParser{
			Data: MapOfMaps{
				"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}
		_, err := p.Get("--", "--")

		if err == nil {
			t.Errorf("expected error 'value does not exist', got: %v", err)
		}
	})
}

func TestSet(t *testing.T) {
	t.Run("section and key do exist", func(t *testing.T) {
		p := INIParser{
			Data: MapOfMaps{
				"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}

		got, err := p.Set("owner", "name", "WalterWhite")
		want := "added"

		if err != nil {
			t.Fail()
		}

		if got != want {
			t.Errorf("got: %v want: %v", got, want)
		}
	})

	t.Run("section or key does not exist", func(t *testing.T) {
		p := INIParser{
			Data: MapOfMaps{
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}
		got, err := p.Set("owner", "name", "WalterWhite")
		want := "not added"

		if err == nil {
			t.Errorf("expected error 'value not added, section or key not found', got: %v", err)
		}

		if got != want {
			t.Errorf("got: %v want: %v", got, want)
		}

	})
}

func TestToString(t *testing.T) {
	t.Run("non empty map", func(t *testing.T) {
		p := INIParser{
			Data: MapOfMaps{
				"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}
		got, err := p.ToString()
		want := "[owner]\nname=JohnDoe\norganization=AcmeWidgetsInc.\n[database]\nfile=payroll.dat\nport=143\nserver=192.0.2.62\n"

		if err != nil {
			t.Fail()
		}

		if got != want {
			t.Errorf("got: %v\nwant: %v", got, want)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		p := INIParser{}
		_, err := p.ToString()

		if err == nil {
			t.Errorf("expected error 'there is no data to convert to string', got: %v", err)
		}
	})
}

func TestSaveToFile(t *testing.T) {
	t.Run("non empty map and valid file path", func(t *testing.T) {
		p := INIParser{
			Data: MapOfMaps{
				"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}

		got, err := p.SaveToFile("../../../../newINI.txt")
		want := "saved"

		if err != nil {
			t.Fail()
		}

		if got != want {
			t.Errorf("got: %v want: %v", got, want)
		}
	})

	t.Run("empty map", func(t *testing.T) {
		p := INIParser{}

		got, err := p.SaveToFile("../../../../newINI.txt")
		want := "not saved"

		if err == nil {
			t.Errorf("expected error 'there is no data to save to file', got: %v", err)
		}

		if got != want {
			t.Errorf("got: %v want: %v", got, want)
		}

	})

	t.Run("invalid file path", func(t *testing.T) {
		p := INIParser{
			Data: MapOfMaps{
				"owner":    {"name": "JohnDoe", "organization": "AcmeWidgetsInc."},
				"database": {"server": "192.0.2.62", "port": "143", "file": "payroll.dat"},
			},
		}

		got, err := p.SaveToFile("")
		want := "not saved"

		if err == nil {
			t.Errorf("expected error 'error writing to file', got: %v", err)
		}

		if got != want {
			t.Errorf("got: %v want: %v", got, want)
		}
	})

}
