package main

import "fmt"

type Article struct {
	Title  string
	Author string
}

func (a Article) method() string {
	return fmt.Sprintf("The %q article was written by %s", a.Title, a.Author)
}

func main() {
	a := Article{
		Title:  "Beliver",
		Author: "Yuvak",
	}
	fmt.Println(a.method())

}
