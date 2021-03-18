package main

import "fmt"

type Books struct {
	pages int
}

type BooksNew []Books

func (bookNew BooksNew) modify()  {
	bookNew[0].pages = 500
	bookNew = append(bookNew, Books{789})
}

func (bookNew BooksNew) change()  {
	//这两处修改都不能反映到Modify方法之外的原因是append函数调用将开辟一块新的内存来存储它返回的结果切片的元素。
	// 而此结果切片的前两个元素是属主参数切片的元素的副本。对此副本所做的修改不会反映到Modify方法之外。
	bookNew = append(bookNew, Books{789})
	bookNew[0].pages = 600
}

func (bookNew *BooksNew) charge ()  {
	*bookNew = append(*bookNew, Books{654})
	(*bookNew)[0].pages = 700
}

func (b Books) SetPages (pages int)  {
	b.pages = pages
}

func main()  {
	var b Books
	b.SetPages(123)//和普通参数传参一样，属主参数的传参也是一个值复制过程。 所以，在方法体内对属主参数的直接部分的修改将不会反映到方法体外。
	fmt.Println(b.pages)

	var newBook = BooksNew{{123},{456}}
	newBook.modify()
	fmt.Println(newBook)
	var newBooks = BooksNew{{124},{456}}
	newBooks.change()
	fmt.Println(newBooks)
	var bookCharge = BooksNew{{111},{777}}
	bookCharge.charge()
	fmt.Println(bookCharge)
}
