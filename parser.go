package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
)

func parseInI(filename string) (map[string]map[string]string, error) {
	ini := make(map[string]map[string]string)
	var head string

	fh, err := os.Open(filename)
	if err != nil {
		return ini, fmt.Errorf("Could not open file '%v': %v", filename, err)
	}
	sectionHead := regexp.MustCompile(`^\[([^]]*)\]\s*$`)
	keyValue := regexp.MustCompile(`^(\w*)\s*=\s*(.*?)\s*$`)
	reader := bufio.NewReader(fh)
	for {
		line, _ := reader.ReadString('\n')
		result := sectionHead.FindStringSubmatch(line)
		if len(result) > 0 {
			head = result[1]
			ini[head] = make(map[string]string)
			continue
		}

		result = keyValue.FindStringSubmatch(line)
		if len(result) > 0 {
			key, value := result[1], result[2]
			ini[head][key] = value
			continue
		}

		if line == "" {
			break
		}
	}

	return ini, nil
}

func getSections(name string, MAP map[string]map[string]string) (map[string]string, error) {
	section := MAP[name]
	if section == nil {
		return nil, errors.New("No Section with that name")
	}

	return section, nil
}

func getSectionsName(Map4 map[string]map[string]string) {
	//section := Map4
	for sectionName := range Map4 {
		fmt.Println(sectionName)
	}
}

func set(name string, key string, value string, MAP2 map[string]map[string]string) map[string]string {
	section := MAP2[name]
	section[key] = value
	return section

}

func getKeys(sectionName string, MAP3 map[string]map[string]string) error {
	section := MAP3[sectionName]
	if section != nil {
		for key, value := range section {
			fmt.Println(key, value)
		}

	} else {
		return errors.New("No Section with that name !!")
	}

	return nil

}

func main() {
	var x = map[string]map[string]string{}
	fmt.Println(parseInI("PHP.ini")) //return all sections with their keys
	fmt.Println("\n")

	x, _ = parseInI("PHP.ini")
	fmt.Println(x)
	fmt.Println("\n")

	fmt.Println(getSections("database", x)) //return specific section with its keys
	fmt.Println("\n")

	getKeys("database", x)
	fmt.Println("\n")

	fmt.Println(set("database", "port", "500", x)) //update value of key
	fmt.Println("\n")

	//getSectionsName(x)
	//temp map[string]map[string]string
	//temp = parseIni("PHP.ini")
	//fmt.Println(getSection("owner", temp))
}
