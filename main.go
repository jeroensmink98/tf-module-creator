package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	moduleName := flag.String("name", "", "The name of your Terraform module")
	flag.Parse()

	if err := createDir(*moduleName); err != nil {
		log.Fatalf("Error: could not create directory %v", err)
	}

	if err := createFile(*moduleName, "main.tf"); err != nil {
		log.Fatalf("Error: could not create file %s %s", "main.tf", err)
	}

	if err := createFile(*moduleName, "variables.tf"); err != nil {
		log.Fatalf("Error: could not create file %s %s", "variables.tf", err)
	}

	if err := createFile(*moduleName, "outputs.tf"); err != nil {
		log.Fatalf("Error: could not create file %s %s", "outputs.tf", err)
	}

	if err := createFile(*moduleName, "README.md"); err != nil {
		log.Fatalf("Error: could not create file %s %s", "README.md", err)
	}

	fmt.Println("Files created successfully!")
}

func createDir(dirName string) error {
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func createFile(dirName string, fileName string) error {
	cwd, _ := os.Getwd()
	path := filepath.Join(cwd, dirName, fileName)
	dst, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dst.Close()
	return nil
}
