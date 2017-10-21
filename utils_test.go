package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"testing"
	"time"
)

func TestFileCreatedWithCurrentDate(t *testing.T) {
	Convey("File is created", t, func() {
		t := time.Now()
		t.String()
		CreateFile()
		currentDate := t.Format("2006-01-02")
		_, err := ioutil.ReadFile(currentDate + ".txt")
		So(err, ShouldEqual, nil)
	})
}

func TestFileCreatedwithCustomName(t *testing.T) {

}

func TestReadOpensLatestFile(t *testing.T) {

}

func TestReadOldFile(t *testing.T) {

}

func ReadCustomFile(t *testing.T) {

}
