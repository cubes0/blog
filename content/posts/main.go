package main

import "fmt"

type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) Pages2() int {
	return b.Pages()
}

func main() {
	var b = Book{pages: 123}
	var p = &b
	var f1 = b.Pages //123
	var f2 = p.Pages //
	var g1 = p.Pages2
	var g2 = b.Pages2
	b.pages = 789
	fmt.Println(f1()) //123
	fmt.Println(f2()) //123
	fmt.Println(g1()) //
	fmt.Println(g2())
}
