package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestFileCreatedWithCurrentDate(t *testing.T) {
	Convey("File is created", t, func() {
		t := time.Now()
		t.String()
		fileName := t.Format("2006-01-02") + ".md"
		CreateFile(fileName)
		So(CreateFile(fileName), ShouldBeNil)
	})
}

func TestFileCreatedwithCustomName(t *testing.T) {
	Convey("File is creted with custom name", t, func() {

		fileName := "foobar.md"
		CreateFile(fileName)
		So(CreateFile(fileName), ShouldBeNil)
	})
}

func TestReadOpensLatestFile(t *testing.T) {
	Convey("File is readable", t, func() {
		content := readFile("")
		So(len(content), ShouldEqual, 0)
	})
}

func TestReadOldFile(t *testing.T) {
}

func ReadCustomFile(t *testing.T) {
	Convey("Read from custom file", t, func() {
		fileName := "foobar.md"
		content := readFile(fileName)
		So(len(content), ShouldEqual, 0)
	})
}
