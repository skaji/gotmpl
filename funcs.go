package main

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

type Values = map[string]interface{}

func Exec(command string) (string, error) {
	b, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		return "", err
	}
	return string(b), err
}

func ReadFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func FromYaml(str string) (Values, error) {
	out, err := FromYamlMulti(str)
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return Values{}, nil
	}
	return out[0], nil
}

func FromYamlMulti(str string) ([]Values, error) {
	out := []Values{}
	decoder := yaml.NewDecoder(strings.NewReader(str))
OUTER:
	for {
		o := Values{}
		switch err := decoder.Decode(&o); err {
		case nil:
			out = append(out, o)
		case io.EOF:
			break OUTER
		default:
			return nil, err
		}
	}
	return out, nil
}

func ToYaml(v interface{}) (string, error) {
	data, err := yaml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
