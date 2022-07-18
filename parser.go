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

type dataStructue struct { //nested maps
	nested_map map[string]map[string]string
}

func (ds *dataStructue) LoadFromFile(path string) (err error) {
	content, err := ioutil.ReadFile(path)
	msg := string(content)
	if err != nil {
		panic(err)
	}
	ds.nested_map, err = Parse_ini(msg)
	return err
}

func (ds *dataStructue) LoadFromText(txt string) (err error) {
	if len(txt) != 0 {
		ds.nested_map, err = Parse_ini(txt)
	} else {
		panic(err)
	}

	return err

}

func Parse_ini(txt string) (map[string]map[string]string, error) {
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

func (ds *dataStructue) GetSections() map[string]map[string]string {
	return ds.nested_map

}

func (ds *dataStructue) GetSectionsName() {

	for sectionName := range ds.nested_map {
		fmt.Println(sectionName)
	}
}

func (ds *dataStructue) Set(name string, key string, value string) map[string]string {
	section := ds.nested_map[name]
	section[key] = value
	fmt.Println(section)
	return section
}

func (ds *dataStructue) GetKeys(sectionName string) error {
	section := ds.nested_map[sectionName]
	if section != nil {
		for key, value := range section {
			fmt.Println(key, value)
		}

	} else {
		return errors.New("No Section with that name !!")
	}

	return nil

}

func (ds *dataStructue) LoadToFile() (err error) {
	file, err := os.Create("output.ini")

	if err != nil {
		log.Fatal(err)
	}

	data := ""
	for section, keys := range ds.nested_map {
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

	parser := dataStructue{}

	parser.LoadFromFile("/home/aya/codescalers/parser_ini/PHP.ini")

	fmt.Println(parser.GetSections())

	parser.GetSectionsName()

	parser.Set("owner", "name", "alaa")

	parser.GetKeys("database")

	parser.LoadToFile()

}
