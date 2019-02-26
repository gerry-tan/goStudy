package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"goStudy/protobuf/person"
	"io/ioutil"
)

func writeProto(fileName string) error {
	var contactBook person.ContactBook
	for i := 0; i < 64; i++ {
		p := &person.Person{
			Id:   int32(i),
			Name: fmt.Sprintf("Jame%d", i),
		}

		phone := &person.Phone{
			Type:   person.PhoneType_HOME,
			Number: "15874057227",
		}

		p.Phones = append(p.Phones, phone)

		contactBook.Persons = append(contactBook.Persons, p)
	}

	data, err := proto.Marshal(&contactBook)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fileName, data, 0755)
	if err != nil {
		return err
	}
	return nil
}

func readProto(fileName string) error {
	var contactBook person.ContactBook
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = contactBook.XXX_Unmarshal(data)
	if err != nil {
		return err
	}
	fmt.Printf("proto Unmarshal: %v\n", contactBook)

	/*err = proto.Unmarshal(data, &contactBook)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", contactBook)*/
	return nil
}

func main() {

	fileName := "/Users/tanqian/go/src/goStudy/protobuf/ctbk.dat"
	err := writeProto(fileName)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("write finished.")

	err = readProto(fileName)
	if err != nil {
		panic(err)
		return
	}

}
