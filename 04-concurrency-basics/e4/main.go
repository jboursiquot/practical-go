package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type proverb struct {
	line  string
	chars map[rune]int
}

func (p *proverb) charCount() map[rune]int {
	if p.chars != nil {
		return p.chars
	}

	m := make(map[rune]int, 0)
	for _, c := range p.line {
		m[c] = m[c] + 1
	}
	p.chars = m
	return p.chars
}

func main() {
	path := pathFromFlag()
	if path == "" {
		path = pathFromEnv()
	}

	if path == "" {
		fmt.Println("You must specify one the file path with -f or as FILE environment variable.")
		os.Exit(1)
	}

	proverbs, err := loadProverbs(path)
	if err != nil {
		fmt.Printf("Failed to load proverbs: %s", err)
		os.Exit(1)
	}

	ch1 := make(chan *proverb)
	ch2 := make(chan *proverb)
	go printProverbs(ch1, ch2)
	for i, p := range proverbs {
		if i%2 == 0 {
			ch1 <- p
		} else {
			ch2 <- p
		}
	}
	close(ch1)
	close(ch2)
}

func pathFromFlag() string {
	path := flag.String("f", "", "file flag")
	flag.Parse()
	return *path
}

func pathFromEnv() string {
	return os.Getenv("FILE")
}

func printProverbs(pc1, pc2 chan *proverb) {
	// functions are first class in Go
	// they can be assigned to a variable and referenced later
	pp := func(p *proverb, whichChan string) {
		fmt.Printf("%s -- %s\n", whichChan, p.line)
		for k, v := range p.charCount() {
			fmt.Printf("'%c'=%d, ", k, v)
		}
		fmt.Print("\n\n")
	}

	for {
		select {
		case p := <-pc1:
			pp(p, "FIRST")
		case p := <-pc2:
			pp(p, "SECOND")
		}
	}
}

func loadProverbs(path string) ([]*proverb, error) {
	var proverbs []*proverb

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bs), "\n")
	for _, line := range lines {
		p := &proverb{line: line}
		proverbs = append(proverbs, p)
	}

	return proverbs, nil
}
