package main

// import (
// 	"os"
// 	"os/exec"
// 	"testing"
// )

// func TestSafeWriteFile(t *testing.T) {
// 	cmd := exec.Command("go", "./main.go")
// 	err := cmd.Run()
// 	if err != nil {
// 		panic(err)
// 	}

// 	folderName := "testing"
// 	err = os.MkdirAll(folderName, 0777)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = os.Chdir("./" + folderName)
// 	if err != nil {
// 		panic(err)
// 	}

// 	files := []string{"main.go", "main_test.go"}

// 	defer func() {
// 		err := os.RemoveAll(folderName)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}()

// 	for _, filename := range files {
// 		if _, err := os.Stat(filename); os.IsNotExist(err) {
// 			t.Errorf("expected to find %s but it doesn't exist", filename)
// 		}
// 	}
// }

// // Testing stdo is hard 🤮
// // func TestMain(t *testing.T) {
// // 	dir, err := os.Getwd()
// // 	fp := filepath.Join(dir, "main.go")
// // 	cmd := exec.Command("/usr/local/bin/go " + fp)

// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	cmd.Dir = dir

// // 	var stdBuffer bytes.Buffer
// // 	mw := io.MultiWriter(os.Stdout, &stdBuffer)

// // 	cmd.Stdout = mw
// // 	cmd.Stderr = mw

// // 	// Execute the command
// // 	if err := cmd.Run(); err != nil {
// // 		log.Panic(err)
// // 	}

// // 	log.Println(stdBuffer.String())

// // }
