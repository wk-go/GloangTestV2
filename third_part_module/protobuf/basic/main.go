package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"protobuf_basic/models"
)

//go:generate protoc -I=proto --go_out=models  addressbook.proto

func main() {

	// Writing a Message
	addressBook := &models.AddressBook{
		People: []*models.Person{
			{
				Id:    1234,
				Name:  "John Doe",
				Email: "jdoe@example.com",
				Phones: []*models.Person_PhoneNumber{
					{Number: "555-4321", Type: models.PhoneType_PHONE_TYPE_HOME},
					{Number: "555-1234", Type: models.PhoneType_PHONE_TYPE_MOBILE},
				},
			},
			{
				Id:   4321,
				Name: "Kam",
				Phones: []*models.Person_PhoneNumber{
					{Number: "556-4321", Type: models.PhoneType_PHONE_TYPE_WORK},
				},
			},
			{
				Id:   4322,
				Name: "Susan",
				Phones: []*models.Person_PhoneNumber{
					{Number: "525-4321", Type: models.PhoneType_PHONE_TYPE_HOME},
					{Number: "566-1234", Type: models.PhoneType_PHONE_TYPE_MOBILE},
				},
			},
		},
	}

	fname := "address-book.bin"

	// Write the new address book back to disk.
	out, err := proto.Marshal(addressBook)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
	fmt.Println("Writing file", out)

	// Read the existing address book.
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	fmt.Println("Reading file", in)
	pIn := &models.AddressBook{}
	if err := proto.Unmarshal(in, pIn); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Printf("Person info:\n%#v\n", pIn)
	for _, p := range pIn.People {
		fmt.Println("------------------")
		fmt.Printf("Person ID: %d\n", p.Id)
		fmt.Printf("  Name: %s\n", p.Name)
		if p.Email != "" {
			fmt.Printf("  E-mail address: %s\n", p.Email)
		}
		for _, pn := range p.Phones {
			fmt.Printf("  Phone number (%s): %s\n", pn.Type.String(), pn.Number)
		}
		fmt.Println("------------------")
	}

}
