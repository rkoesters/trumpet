// Package markov provides a trumpet.Generator that uses markov chains.
package markov

// modified from https://golang.org/doc/codewalk/markov/

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"sync"
)

type Prefix []string

func (p Prefix) String() string {
	return strings.Join(p, " ")
}

func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

type Chain struct {
	chain     map[string][]string
	prefixLen int
	mutex     *sync.Mutex
}

func NewChain(prefixLen int) *Chain {
	return &Chain{
		chain:     make(map[string][]string),
		prefixLen: prefixLen,
		mutex:     new(sync.Mutex),
	}
}

func (c *Chain) Train(s string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.build(strings.NewReader(s))
}

func (c *Chain) build(r io.Reader) {
	br := bufio.NewReader(r)
	p := make(Prefix, c.prefixLen)
	for {
		var s string

		_, err := fmt.Fscan(br, &s)
		if err != nil {
			break
		}

		key := p.String()
		c.chain[key] = append(c.chain[key], s)
		p.Shift(s)
	}
}

func (c *Chain) generateWords(n int) []string {
	var words []string

	p := make(Prefix, c.prefixLen)

	for i := 0; i < n; i++ {
		choices := c.chain[p.String()]
		if len(choices) == 0 {
			break
		}
		next := choices[rand.Intn(len(choices))]
		words = append(words, next)
		p.Shift(next)
	}
	return words
}

func (c *Chain) Generate(maxLength int) string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for {
		numWords := maxLength / 6
		words := c.generateWords(numWords)
		text := strings.Join(words, " ")
		if len(text) <= maxLength {
			return text
		}
	}
}
