package main

import "fmt"

func main() {
	/*
		//cara pertama buat struct
		alex := person{firstName: "Alex", lastName: "Anderson"}
	*/

	//cara kedua
	/*contact := contactInfo{email: "tes@ff.com", zip: 12345}

	var alex person
	alex.firstName = "Ales"
	alex.lastName = "Teles"
	alex.contact = contact
	*/

	alex := person{
		firstName: "Alex",
		lastName:  "Teles",
		contact: contactInfo{
			email: "Afds@tes.com",
			zip:   12345,
		},
	}

	rere := person{
		firstName: "RERE",
		lastName:  "HEGE",
		contact: contactInfo{
			email: "de@dslf.com",
			zip:   432145,
		},
	}

	//alexPointer := &alex

	//alex.updatePerson(rere)
	//alex.updateName("LEXA")
	//alex.updateNameNoPointer("TES")

	updatePerson(alex)

	alex.print()

	//Map

	personMap := make(map[string]person)

	personMap["alex"] = alex
	personMap["rere"] = rere

	// colors := map[string]string{
	// 	"red":  "00001",
	// 	"blue": "00002",
	// }

	//delete(personMap, "rere")

	//printMap(personMap)

	//fmt.Println(personMap["rere"].contact.zip)
}

func printMap(c map[string]person) {
	for name, person := range c {
		fmt.Println("name of person", name, "is", person.firstName)
	}

}
func updatePerson(per person) {
	per.firstName = "TES"
}

func (p *person) updateName(new string) {
	(*p).firstName = new
}

func (p person) updateNameNoPointer(new string) {
	p.firstName = new
}
func (p person) print() {
	fmt.Printf("%v", p)
}
