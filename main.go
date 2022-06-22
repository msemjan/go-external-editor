package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	// Create a temporary file
	f, err1 := ioutil.TempFile("", "tmpfile-")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer f.Close()
	defer os.Remove(f.Name())

	f.WriteString("{\n")
	f.WriteString("\t\"name\": \"\",\n")
	f.WriteString("\t\"surname\": \"\",")
	f.WriteString("\n\t\"age\": ,\n")
	f.WriteString("\t\"profession\": \"\"\n")
	f.WriteString("}")

	// Open the temporary file in Vim for editing
	cmd := exec.Command("vim", f.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err2 := cmd.Start()
	if err2 != nil {
		log.Printf("2")
		log.Fatal(err2)
	}

	err2 = cmd.Wait()
	if err2 != nil {
		log.Printf("Error while editing. Error: %v\n", err2)
	} else {
		log.Printf("Successfully edited.")
	}

	// Output the content of the file line by line
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		log.Println(fileScanner.Text())
	}

}
