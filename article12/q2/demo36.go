package main

type Dog struct {
		name string
}

func New(name string)Dog  {
		return Dog{name}
}

func (dog *Dog)SetName(name string)  {
		dog.name = name
}

func (dog Dog)Name()string  {
		return dog.name
}

func main()  {
		//示例１
		//New("little pig").SetName("monster")

		//示例２
		map[string]int{"The":0,"word":0,"counter": 0}["word"]++
		map1 := map[string]int{"the": 0, "word": 0, "counter": 0}
		map1["word"]++
}
