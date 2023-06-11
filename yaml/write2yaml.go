package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"yaml/datamodels"

	"gopkg.in/yaml.v2"
)

func readYaml2Struct() {
	f := &datamodels.Developer{}
	source, err := ioutil.ReadFile(".//files/developer.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(source), &f)
	if err != nil {
		log.Fatal(err)
	}

	for lang, hrs := range f.Languages {
		fmt.Printf("%s has %.2f hrs\n", lang, hrs)
	}
}

func writeStruct2Yaml() {
	d := &datamodels.Developer{
		"0126517407",
		3,
		90000.00,
		[]string{"feature/task123", "feature/314", "feature/567"},
		[]int{7, 9, 6, 8, 7},
		map[string]float32{"Golang": 9.1, "Python": 7.9},
	}

	y, er1 := yaml.Marshal(d) // convert the Developer struct to YAML.
	if er1 != nil {
		er1 = fmt.Errorf("%s", er1)
	}

	er2 := ioutil.WriteFile(".\\files\\developer.yaml", y, 0644)
	if er2 != nil {
		er2 = fmt.Errorf("%s", er2)
	}
}

func writeList2Yaml() {
	c1 := &datamodels.Company{"513", true, []string{"David", "Fred"}, 900000, map[string]float32{"Golang": 3, "C++": 4}}
	c2 := &datamodels.Company{"514", true, []string{"Alex", "Yap"}, 758000, map[string]float32{"Python": 5, "Firebase": 5}}

	list := &datamodels.CompanyGroup{}
	list.Companies = append(list.Companies, *c1, *c2)

	y, er1 := yaml.Marshal(list)
	if er1 != nil {
		er1 = fmt.Errorf("%v", er1)
	}

	er2 := ioutil.WriteFile(".\\files\\companies.yaml", y, 0644)
	if er2 != nil {
		er2 = fmt.Errorf("%v", er2)
	}
}
func main() {
	writeStruct2Yaml()
	writeList2Yaml()
	readYaml2Struct()
}
