package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

var format = "---\r\ntitle: %s\r\nhidden: true\r\n---\r\n"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func makeSlug(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "-", -1)
	if len(s) > 20 {
		s = s[0:20]
		if len(s) == 20 && s[19] == '-' {
			s = s[0:19]
		}
	}
	return s
}

func main() {

	fmt.Print("New post. What's the title?\n")

	o, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	title := strings.Trim(o, " \r\n")
	date := time.Now().Format("2006-01-02")
	slang := date + "-" + makeSlug(title)

	// if not exist
	if _, err := os.Stat("en/_posts/" + slang + ".md"); err == nil {
		panic("Shit! File Exists!")
	}

	mem := []byte(fmt.Sprintf(format, title))
	check(ioutil.WriteFile(fmt.Sprint("en/_posts/", slang, ".md"), mem, 0644))
	check(ioutil.WriteFile(fmt.Sprint("id/_posts/", slang, ".md"), mem, 0644))

}

// This is a small utility to easily create new post
// To make exe:
// go build -o newPost.exe newPost.go
