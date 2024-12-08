package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

var version string = "dev"

func main() {
	os.Exit(run(os.Stdout, os.Stderr, os.Args))
}

func run(stdout io.Writer, stderr io.Writer, args []string) int {
	if len(args) == 2 {
		if args[1] == "-h" || args[1] == "--help" {
			fmt.Fprintf(stdout, "Usage:\n")
			fmt.Fprintf(stdout, "  ❯ %s FILE\n", args[0])
			fmt.Fprintf(stdout, "  ❯ cat FILE | %s\n", args[0])
			return 1
		}
		if args[1] == "-v" || args[1] == "--version" {
			fmt.Fprintln(stdout, version)
			return 0
		}
	}

	file := ""
	if len(args) > 1 {
		file = args[1]
	}
	var content []byte
	var err error
	if file == "-" || file == "" {
		if content, err = io.ReadAll(os.Stdin); err != nil {
			fmt.Fprintf(stderr, "os.Stdin: %s\n", err)
			return 1
		}
	} else {
		if content, err = os.ReadFile(file); err != nil {
			fmt.Fprintln(stderr, err)
			return 1
		}
	}

	out, err := process(string(content))
	if err != nil {
		fmt.Fprintf(stderr, "process %s: %s\n", file, err)
		return 1
	}

	fmt.Fprint(stdout, out)
	return 0
}

func process(content string) (string, error) {
	funcMap := sprig.TxtFuncMap()
	funcMap["exec"] = Exec
	funcMap["fromYaml"] = FromYaml
	funcMap["fromYamlMulti"] = FromYamlMulti
	funcMap["readFile"] = ReadFile
	funcMap["toYaml"] = ToYaml

	tmpl, err := template.New("stringTemplate").Funcs(funcMap).Parse(content)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, struct{}{}); err != nil {
		return "", err
	}

	return buf.String(), nil
}
