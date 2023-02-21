package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

/*
 * This is a simple example that shows how to open a temporary file, let the
 * user edit it, and then output the data into STDOUT.
 *
 * In real world, there probably would be some file or data in a database that
 * is read, then modified by the user, and something meaningful happens with
 * the input. Alternative usage would be for semi-automated scraping of the
 * web. The Go program would download some data, the user would check them, and
 * edit if necessary.
 */
func main() {
	// Create a temporary file
	f, err1 := ioutil.TempFile("", "tmpfile-")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer f.Close()
	defer os.Remove(f.Name())
	
	// Fill the termporary file with some data
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

	// Open the temporary file in an external editor
	err2 := cmd.Start()
	if err2 != nil {
		log.Printf("2")
		log.Fatal(err2)
	}
	
	// Wait for the user to do their thing...
	err2 = cmd.Wait()

	// Check if they didn't mess up everything
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
