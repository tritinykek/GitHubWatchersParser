package main

import (
	"fmt"
	"github.com/opesun/goquery"
	"os"
)

func parseWatchers(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' && str[i] != '\n' {
			for j := i; j < len(str); j++ {
				if str[j] == ' ' || str[j] == '\n' {
					return str[i:j]
				}
			}
		}
	}
	return ""
}

func main() {
	url := ""
	fmt.Print("Print ur github project url: ")
	fmt.Fscan(os.Stdin,&url)
	x, err := goquery.ParseUrl(url)

	if err == nil {
		s := x.Find(".social-count").Text()
		s = parseWatchers(s)
		fmt.Print("U have ")
		fmt.Print(s)
		fmt.Print(" watchers")
	}
}
