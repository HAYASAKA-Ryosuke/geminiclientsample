package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/url"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
)

func main() {
	address := "<GEMINI_ADDRESS>"

	u, err := url.Parse(address)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to url parse: %v", err)
		os.Exit(1)
	}

	conn, err := tls.Dial("tcp", u.Host+":1965", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to dial: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(u.String() + "\r\n"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed connection: %v", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(conn)

	if scanner.Scan() {
		scanner.Text()
	}

	for scanner.Scan() {
		fmt.Println(string(markdown.Render(scanner.Text(), 128, 10)))
	}
}
