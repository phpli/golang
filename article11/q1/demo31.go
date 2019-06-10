package main

import "fmt"

type Pet interface {
			SetName(name string)
			Name() string
			Category() string
}

type Dog struct {
	name string
}

func (dog *Dog)SetName(name string)  {
		dog.name = name
}

func (dog Dog) Name() string  {
		return dog.name
}

func (dog Dog)Category()string  {
		return "dog"
}

func main()  {
		//示例１
		dog := Dog{"little pig"}
		_,ok := interface{}(dog).(Pet)
		fmt.Printf("Dog implements interface Pet:%v\n",ok)
		_,ok1 := interface{}(&dog).(Pet)
		fmt.Printf("*Dog implements interface Pet:%v\n",ok1)
		fmt.Println()

		//示例２
		var pet Pet = &dog
		fmt.Printf("The pet is a %s,the name is %q.\n",pet.Category(),pet.Name())
}


