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

// Prefix is the prefix of a markov chain.
type Prefix []string

// String returns the Prefix as a string.
func (p Prefix) String() string {
	return strings.Join(p, " ")
}

// Shift adds the given word to the prefix, shifting existing words over
// and removing the first word so that the Prefix is the same number of
// words.
func (p Prefix) Shift(word string) {
	copy(p, p[1:])
	p[len(p)-1] = word
}

// Chain is a markov chain that implements trumpet.Generator.
type Chain struct {
	chain     map[string][]string
	prefixLen int
	mutex     *sync.Mutex
}

// NewChain returns a markov chain that implements trumpet.Generator.
func NewChain(prefixLen int) *Chain {
	return &Chain{
		chain:     make(map[string][]string),
		prefixLen: prefixLen,
		mutex:     new(sync.Mutex),
	}
}

// Train adds the given string to the markov chain.
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

// Generate returns a string created from the markov chain that is at
// most maxLength characters long.
func (c *Chain) Generate(maxLength int) string {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for {
		// Divide maxLength by 6 to guess the number of words
		// that will satisfy our maxLength. 6 is just an
		// estimated average word length, it might be more
		// useful to derive that value from the corpus.
		numWords := maxLength / 6
		words := c.generateWords(numWords)
		text := strings.Join(words, " ")
		if len(text) <= maxLength {
			return text
		}
	}
}
