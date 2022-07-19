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

func checkLine(line string) (string, error) {
	if line[0] == '[' && line[len(line)-1] == ']' {
		return "section", nil
	} else if line[0] == ';' {
		return "comment", nil
	} else if line[0] == '\n' {
		return "empty line", nil
	} else if strings.Contains(line, "=") && strings.Count(line, "=") == 1 {
		return "key line", nil
	} else {
		return "", errors.New("Syntax Error !")
	}

}

func (parser *Parser) LoadFromFile(path string) (err error) {
	content, err := ioutil.ReadFile(path)
	msg := string(content)
	if err != nil {
		return err
	}
	parser.nested_map, err = Parse(msg)
	return err
}

func (parser *Parser) LoadFromText(txt string) (err error) {
	if len(txt) != 0 {
		parser.nested_map, err = Parse(txt)
	} else {
		return err
	}

	return err

}

func Parse(txt string) (map[string]map[string]string, error) {
	ini := make(map[string]map[string]string)
	var head string

	sectionHead := regexp.MustCompile(`^\[([^]]*)\]\s*$`)
	keyValue := regexp.MustCompile(`^(\w+)\s*=\s*(.*?)\s*$`)
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
	return ini, nil
}

func (parser *Parser) GetSections() map[string]map[string]string {
	return parser.nested_map
}

func (parser *Parser) GetSectionsName() []string {
	names := []string{}
	for name, _ := range parser.nested_map {
		names = append(names, name)
	}
	return names
}

func (parser *Parser) Set(name string, key string, value string) map[string]string { //ToDo
	section := parser.nested_map[name]
	section[key] = value
	fmt.Println(section)
	return section
}

func (parser *Parser) GetKeys(sectionName string) ([]string, error) {
	keys := []string{}
	section := parser.nested_map[sectionName]
	if section != nil {
		for key, value := range section {
			keys = append(keys, key+":"+value)
		}
	} else {
		return nil, errors.New("No Section with that name !!")
	}
	return keys, nil
}

func (parser *Parser) Get(sectionName string, key string) string {
	section := parser.nested_map[sectionName][key]
	return section

}

func (parser *Parser) SaveToFile() (err error) {
	file, err := os.Create("output.ini")
	file.Close()
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
	return err
}

func main() {
	parser := Parser{}
	parser.LoadFromFile("/home/aya/codescalers/parser_ini/PHP.ini")
	fmt.Println(parser.GetSections())
	fmt.Println(parser.GetSectionsName())
	fmt.Println(parser.GetKeys("database"))
	fmt.Println(parser.Get("owner", "name"))
	parser.SaveToFile()
}
