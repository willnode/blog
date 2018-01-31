package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"regexp"
)

var format = "---\r\ntitle: %s\r\nhidden: true\r\n---\r\n"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var invalids = regexp.MustCompile("[?@#\\/()]|\\(.+\\)")

func makeSlug(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(strings.Replace(s, " ", "-", -1), "'", "", -1)
	s = invalids.ReplaceAllString(s, "")
	return s
}

func main() {

	fmt.Print("New post. What's the title?\n")

	o, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	title := strings.Trim(o, " \r\n")
	date := time.Now().Format("2006-01-02")
	slang := date + "-" + makeSlug(title)

	fmt.Print("Language?\n")

	l, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	l = strings.Trim(l, " \r\n")

	// if not exist
	if _, err := os.Stat(l + "/_posts/" + slang + ".md"); err == nil {
		panic("Shit! File Exists!")
	}

	mem := []byte(fmt.Sprintf(format, title))
	check(ioutil.WriteFile(fmt.Sprint(l+"/_posts/", slang, ".md"), mem, 0644))

}

// This is a small utility i made to easily create new post
// To make exe:
// go build -o newPost.exe newPost.go
