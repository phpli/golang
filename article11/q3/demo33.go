package main

import (
	"fmt"
	"reflect"
)

type Pet interface {
		Name() string
		Category() string
}

type Dog struct {
		name string
}

func (dog *Dog)setName(name string)  {
		dog.name = name
}

func (dog Dog)Name()string  {
		return dog.name
}

func (dog Dog)Category()string  {
		return "dog"
}



func main()  {
		//示例１
		var dog1 *Dog
		fmt.Printf("The first dog is nil.\n")
		dog2 := dog1
		fmt.Printf("The second dog is nil.\n")
		var pet Pet = dog2
		if pet == nil {
			fmt.Printf("The pet is nil.\n")
		}else {
			fmt.Printf("The pet is not nil.\n")
		}
		fmt.Printf("The type of pet is %T.\n",pet)
		fmt.Printf("The type of pet is %s.\n",reflect.TypeOf(pet).String())
		fmt.Printf("The type of second dog is %T.\n",dog2)
		fmt.Println()
		//示例２
		wrap := func(dog *Dog) Pet{
			if dog == nil{
				return nil
			}
			return dog
		}

		pet = wrap(dog2)

		if pet == nil {
			fmt.Println("The pet is nil.")
		}else {
			fmt.Println("The pet is not nil.")
		}
}
