package gomain

import (
	"github.com/kandros/goutil/editorutil"
	"github.com/kandros/goutil/fileutil"
)

func createMainFile() {
	fileContent := `package main
import "fmt"

func main() {
	fmt.Println("hello")
}`
	fileutil.SafeWriteFile("main.go", []byte(fileContent))
}

func createMainTestFile() {
	fileContent := `package main

import "testing"

func TestMain(t *testing.T) {
	got := "hello"
	want := "world"

	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
	`
	fileutil.SafeWriteFile("main_test.go", []byte(fileContent))
}

func Run(openInEditor bool) {
	createMainFile()
	createMainTestFile()
	if openInEditor {
		editorutil.OpenFileInEditor("./main.go")
	}
}
