package main

import (
	"io/ioutil"
	"time"
)

func CreateFile(fileName string) error {
	t := time.Now()
	t.String()
	if fileName == "" {
		err := ioutil.WriteFile(t.Format("2006-01-02")+".md", nil, 0644)
		return err
	} else {
		err := ioutil.WriteFile(fileName, nil, 0644)
		return err
	}
}

func readFile(fileName string) []byte {
	t := time.Now()
	t.String()
	if fileName == "" {
		content, _ := ioutil.ReadFile(t.Format("2006-01-02") + ".md")
		return content
	} else {
		content, _ := ioutil.ReadFile(fileName)
		return content
	}
}
