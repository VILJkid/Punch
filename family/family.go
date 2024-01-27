package family

import (
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// RunCommand handler of the family-tree commands
func RunCommand() error {
	people, err := parsePeopleFile()
	if err != nil {
		return err
	}
	args := os.Args[2:]
	switch os.Args[1] {
	case cmdAdd:
		return people.Add(args)
	case cmdConnect:
		return people.Connect(args)
	case cmdCount:
		return people.Count(args)
	case cmdFind:
		return people.Find(args)
	}
	return errors.New("unknown command")
}

// Add can add person and relationships
func (p *people) Add(args []string) error {
	switch args[0] {
	case addPerson:
		name := strings.Join(args[1:], " ")
		if name == "" {
			return errors.New("person name not provided")
		}
		person := person{}
		person.Id = len(p.Person) + 1
		person.Name = name
		p.Person = append(p.Person, person)
		if err := p.writeTopeopleFile(); err != nil {
			return err
		}
		log.Println("Person added")
	case addRelationship:
		log.Println("This functinality is built-in")
	default:
		return errors.New("unknown sub-command")
	}
	return nil
}

// Connect attaches the relationships of two people
func (p *people) Connect(args []string) error {
	r := regexp.MustCompile(`([a-zA-Z ]+) as (\w+) of ([a-zA-Z ]+)`)
	res := r.FindAllStringSubmatch(strings.Join(args, " "), -1)
	if len(res) != 1 || len(res[0]) != 4 {
		return errors.New("invalid syntax")
	}
	person1Name := res[0][1]
	var person1Details person
	for _, eachPerson := range p.Person {
		if eachPerson.Name == person1Name {
			person1Details = eachPerson
			break
		}
	}
	if person1Details.Id == 0 {
		return fmt.Errorf("person %s not found", person1Name)
	}
	person2Name := res[0][3]
	var person2Details person
	for _, eachPerson := range p.Person {
		if eachPerson.Name == person2Name {
			person2Details = eachPerson
			break
		}
	}
	if person2Details.Id == 0 {
		return fmt.Errorf("person %s not found", person2Name)
	}
	switch res[0][2] {
	case relationSon:
		person2Details.Relationships.Sons = append(person2Details.Relationships.Sons, person1Details.Id)
	case relationDaughter:
		person2Details.Relationships.Daughters = append(person2Details.Relationships.Daughters, person1Details.Id)
	case relationWife:
		person2Details.Relationships.Wives = append(person2Details.Relationships.Wives, person1Details.Id)
	case relationFather:
		if person2Details.Relationships.Father != 0 {
			return errors.New("relation already exists")
		}
		person2Details.Relationships.Father = person1Details.Id
	default:
		return errors.New("undefined relation")
	}
	p.Person[person2Details.Id-1] = person2Details
	return p.writeTopeopleFile()
}

// Count return the number of relationships of a person
func (p *people) Count(args []string) error {
	r := regexp.MustCompile(`(\w+) of ([a-zA-Z ]+)`)
	res := r.FindAllStringSubmatch(strings.Join(args, " "), -1)
	if len(res) != 1 || len(res[0]) != 3 {
		return errors.New("invalid syntax")
	}
	personName := res[0][2]
	var personDetails person
	for _, eachPerson := range p.Person {
		if eachPerson.Name == personName {
			personDetails = eachPerson
			break
		}
	}
	if personDetails.Id == 0 {
		return fmt.Errorf("person %s not found", personName)
	}
	switch res[0][1] {
	case relationSons:
		log.Printf("%s has %d %s\n", personName, len(personDetails.Relationships.Sons), res[0][1])
	case relationDaughters:
		log.Printf("%s has %d %s\n", personName, len(personDetails.Relationships.Daughters), res[0][1])
	case relationWives:
		log.Printf("%s has %d %s\n", personName, len(personDetails.Relationships.Wives), res[0][1])
	default:
		return errors.New("undefined relation")
	}
	return nil
}

// Find returns the names of people in a relationship
func (p *people) Find(args []string) error {
	r := regexp.MustCompile(`(\w+) of ([a-zA-Z ]+)`)
	res := r.FindAllStringSubmatch(strings.Join(args, " "), -1)
	if len(res) != 1 || len(res[0]) != 3 {
		return errors.New("invalid syntax")
	}
	personName := res[0][2]
	var personDetails person
	for _, eachPerson := range p.Person {
		if eachPerson.Name == personName {
			personDetails = eachPerson
			break
		}
	}
	if personDetails.Id == 0 {
		return fmt.Errorf("person %s not found", personName)
	}
	switch res[0][1] {
	case relationSons:
		if len(personDetails.Relationships.Sons) == 0 {
			return fmt.Errorf("person %s doesn't have any %s", personName, res[0][1])
		}
		sons := []string{}
		for _, eachSon := range personDetails.Relationships.Sons {
			sons = append(sons, p.Person[eachSon-1].Name)
		}
		log.Printf("%s of %s are %s", res[0][1], personName, strings.Join(sons, ","))
	case relationDaughters:
		if len(personDetails.Relationships.Daughters) == 0 {
			return fmt.Errorf("person %s doesn't have any %s", personName, res[0][1])
		}
		daughters := []string{}
		for _, eachSon := range personDetails.Relationships.Daughters {
			daughters = append(daughters, p.Person[eachSon-1].Name)
		}
		log.Printf("%s of %s are %s", res[0][1], personName, strings.Join(daughters, ","))
	case relationWives:
		if len(personDetails.Relationships.Wives) == 0 {
			return fmt.Errorf("person %s doesn't have any %s", personName, res[0][1])
		}
		wives := []string{}
		for _, eachSon := range personDetails.Relationships.Wives {
			wives = append(wives, p.Person[eachSon-1].Name)
		}
		log.Printf("%s of %s are %s", res[0][1], personName, strings.Join(wives, ","))
	case relationFather:
		if personDetails.Relationships.Father == 0 {
			return fmt.Errorf("person %s doesn't have a %s", personName, res[0][1])
		}
		log.Printf("%s is the %s of %s", p.Person[personDetails.Relationships.Father-1].Name, res[0][1], personName)
	default:
		return errors.New("undefined relation")
	}
	return nil
}

// parsePeopleFile parses the people file
func parsePeopleFile() (*people, error) {
	people := &people{}
	peopleFileBytes, err := os.ReadFile(peopleFileName)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return nil, err
	}
	if err = yaml.Unmarshal(peopleFileBytes, people); err != nil {
		return nil, err
	}
	return people, nil
}

// writeTopeopleFile writes the people contents to people file
func (p *people) writeTopeopleFile() error {
	peopleFileBytes, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return os.WriteFile(peopleFileName, peopleFileBytes, os.ModePerm)
}
