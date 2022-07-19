package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type Parser struct { //nested maps
	nested_map map[string]map[string]string
}

func (parser *Parser) LoadFromFile(path string) (err error) {
	content, err := ioutil.ReadFile(path)
	msg := string(content)
	if err != nil {
		return err
	}
	parser.nested_map, err = Parseini(msg)
	return err
}

func (parser *Parser) LoadFromText(txt string) (err error) {
	if len(txt) != 0 {
		parser.nested_map, err = Parseini(txt)
	} else {
		return err
	}

	return err

}

func Parseini(txt string) (map[string]map[string]string, error) {
	ini := make(map[string]map[string]string)
	var head string

	sectionHead := regexp.MustCompile(`^\[([^]]*)\]\s*$`)
	keyValue := regexp.MustCompile(`^(\w*)\s*=\s*(.*?)\s*$`)
	reader := bufio.NewReader(strings.NewReader(txt))

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
	fmt.Println(ini)
	return ini, nil

}

func (parser *Parser) GetSections() map[string]map[string]string {
	return parser.nested_map

}

func (parser *Parser) GetSectionsName() {

	for sectionName := range parser.nested_map {
		fmt.Println(sectionName)
	}
}

func (parser *Parser) Set(name string, key string, value string) map[string]string {
	section := parser.nested_map[name]
	section[key] = value
	fmt.Println(section)
	return section
}

func (parser *Parser) GetKeys(sectionName string) error {
	section := parser.nested_map[sectionName]
	if section != nil {
		for key, value := range section {
			fmt.Println(key, value)
		}

	} else {
		return errors.New("No Section with that name !!")
	}

	return nil

}

func (parser *Parser) LoadToFile() (err error) {
	file, err := os.Create("output.ini")

	if err != nil {
		log.Fatal(err)
	}

	data := ""
	for section, keys := range parser.nested_map {
		data += "[" + section + "]\n"
		for key, value := range keys {
			data += key + " = " + value + "\n"
		}

	}

	file.WriteString(data)
	file.Close()
	return err

}

func main() {

	parser := Parser{}

	parser.LoadFromFile("/home/aya/codescalers/parser_ini/PHP.ini")

	fmt.Println(parser.GetSections())

	parser.GetSectionsName()

	parser.Set("owner", "name", "alaa")

	parser.GetKeys("database")

	parser.LoadToFile()

}
