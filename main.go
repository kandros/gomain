package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func createMainFile() {
	fileContent := `package main
import "fmt"

func main() {
	fmt.Println("hello")
}`
	safeWriteFile("main.go", fileContent)
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
	safeWriteFile("main_test.go", fileContent)
}

func safeWriteFile(filename, fileContent string) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := ioutil.WriteFile(filename, []byte(fileContent), 0777)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("file %s already exists\n", filename)
		os.Exit(0)
	}

}

func openFileInEditor(filename string) {
	editor := os.Getenv("EDITOR")

	var cmd *exec.Cmd
	cmd = exec.Command(editor, filename)
	cmd.Start()
}

func main() {
	createMainFile()
	createMainTestFile()
	openFileInEditor("./main.go")
}
