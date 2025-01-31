package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"log"
	person "person"
)

func main() {

	originalPerson := &person.Person{
		Name:  "Rie",
		Age:   31,
		Email: []string{"rie@example.com", "rie@gmail.com"},
	}

	serializedPerson, err := proto.Marshal(originalPerson)
	if err != nil {
		log.Fatal("Marshalling error", err)
	}

	deSerializedPerson := &person.Person{}
	err = proto.Unmarshal(serializedPerson, deSerializedPerson)
	if err != nil {
		log.Fatal("Unmarshalling error", err)
	}

	fmt.Println("Original person: ", originalPerson)
	fmt.Println("Serialized person: ", serializedPerson)
	fmt.Println("Buff person: ", deSerializedPerson)

}
