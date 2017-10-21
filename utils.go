package main

import (
	"io/ioutil"
	"time"
)

func CreateFile() {
	t := time.Now()
	t.String()
	_ = ioutil.WriteFile(t.Format("2006-01-02")+".txt", nil, 0644)
}
