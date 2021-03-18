package main

import "fmt"

type Book struct {
	pages int
}

func (b Book) Pages() int  {
	return b.pages
}

func (b *Book) SetPages(pages int)  {
	b.pages = pages
}

func main()  {
	var a = 6
	var book Book
	//a = 6
	fmt.Println(1)
	fmt.Println(*(&a))
	fmt.Println(a)
	fmt.Printf("%T \n", book.Pages)
	fmt.Printf("%T \n", (&book).SetPages)
	// &book值有一个隐式方法Pages。
	fmt.Printf("%T \n", (&book).Pages)
	// 调用这三个方法。
	(&book).SetPages(13)
	book.SetPages(13) //等价上一行
	fmt.Println(book.Pages())
	fmt.Println((&book).Pages())
}
