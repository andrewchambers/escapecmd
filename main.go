package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/mattn/go-shellwords"
)

func main() {
	jsonFormat := flag.Bool("json", false, "use json formatting")
	printNewLine := flag.Bool("n", false, "print a new line at the end of the expansion")
	flag.Parse()

	inArgs := flag.Args()
	var buf bytes.Buffer

	if *jsonFormat {
		encBytes, err := json.Marshal(inArgs)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		_, err = buf.Write(encBytes)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	} else {
		outArgs := make([]string, 0, len(inArgs))

		for _, s := range inArgs {
			parsed, err := shellwords.Parse(s)
			if err != nil || len(parsed) != 1 {
				s = strconv.Quote(s)
			}
			outArgs = append(outArgs, s)
		}

		for i, s := range outArgs {
			_, err := buf.WriteString(s)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			if i != len(outArgs)-1 {
				_, err = buf.WriteString(" ")
				if err != nil {
					fmt.Fprintln(os.Stderr, err.Error())
					os.Exit(1)
				}
			}
		}
	}

	eol := ""
	if *printNewLine {
		eol = "\n"
	}

	_, err := fmt.Printf("%s%s", strconv.Quote(buf.String()), eol)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
