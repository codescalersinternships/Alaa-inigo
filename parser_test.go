package main

import (
	"fmt"
	"testing"
)

func TestParseIni(t *testing.T) {
	// parser := dataStructue{
	// 	map[string]map[string]string{
	// 		"owner": {
	// 			"name":         "John Doe",
	// 			"organization": "Acme Widgets Inc.",
	// 		},
	// 		"database": {
	// 			"file":   "payroll.dat",
	// 			"port":   "143",
	// 			"server": "129.0.2.62",
	// 		},
	// 	},
	// }

	text := "; last modified 1 April 2001 by John Doe\n" +
		"[owner]\n" + "name = John Doe\n" + "organization = Acme Widgets Inc.\n" +
		"\n" + "[database]\n" + "; use IP address in case network name resolution is not working\n" +
		"server = 192.0.2.62\n" + "port = 143\n" + "file = payroll.dat\n"

	actual, err := Parse_ini(text)

	if err != nil {
		t.Error(fmt.Sprintf("Error in parsing: '%v'", err))
	}
	expected := make(map[string]map[string]string)
	expected["owner"] = make(map[string]string)
	expected["owner"]["name"] = "John Doe"
	expected["owner"]["organization"] = "Acme Widgets Inc."
	expected["database"] = make(map[string]string)
	expected["database"]["server"] = "192.0.2.62"
	expected["database"]["port"] = "143"
	expected["database"]["file"] = "payroll.dat"
	if !compareMapMap(actual, expected) {
		t.Error(fmt.Sprintf("Expected : '%v', Actual : '%v'", expected, actual))
	}
}

func TestGetSections(t *testing.T) {
	parser := dataStructue{
		map[string]map[string]string{
			"owner": {
				"name":         "John Doe",
				"organization": "Acme Widgets Inc.",
			},
			"database": {
				"file":   "payroll.dat",
				"port":   "143",
				"server": "129.0.2.62",
			},
		},
	}
	actual := parser.GetSections()

	expected := map[string]map[string]string{
		"database": {
			"file":   "payroll.dat",
			"port":   "143",
			"server": "129.0.2.62",
		},
		"owner": {
			"name":         "John Doe",
			"organization": "Acme Widgets Inc.",
		},
	}

	if !compareMapMap(actual, expected) {
		t.Error(fmt.Sprintf("Expected '%v', Actual '%v'", expected, actual))
	}
}

func TestSet(t *testing.T) {
	parser := dataStructue{
		map[string]map[string]string{
			"owner": {
				"name":         "John Doe",
				"organization": "Acme Widgets Inc.",
			},
			"database": {
				"file":   "payroll.dat",
				"port":   "143",
				"server": "192.0.2.62",
			},
		},
	}

	expected := map[string]string{
		"file":   "payroll.dat",
		"port":   "500",
		"server": "192.0.2.62",
	}

	actual := parser.Set("database", "port", "500")

	if !compareMap(actual, expected) {
		t.Error(fmt.Sprintf("Expected '%v', Actual '%v'", expected, actual))
	}

}

func TestGetKeys(t *testing.T) {

	parser := dataStructue{
		map[string]map[string]string{
			"owner": {
				"name":         "John Doe",
				"organization": "Acme Widgets Inc.",
			},
			"database": {
				"file":   "payroll.dat",
				"port":   "143",
				"server": "192.0.2.62",
			},
		},
	}

	expected := map[string]string{
		"server": "192.0.2.62",
		"port":   "143",
		"file":   "payroll.dat",
	}

	err := parser.GetKeys("database")

	if err != nil {
		t.Error(fmt.Sprintf("Expected '%v' ", expected))
	}

}

func compareMapMap(a, b map[string]map[string]string) bool {

	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if !compareMap(value, b[key]) {
			return false
		}
	}
	for key, value := range b {
		if !compareMap(value, a[key]) {
			return false
		}
	}
	return true
}

func compareMap(a, b map[string]string) bool {

	if len(a) != len(b) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	for key, value := range b {
		if value != a[key] {
			return false
		}
	}
	return true
}
