package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	s := getArrayOfFileNamesInPath("./folder")
	t := getArrayOfRowsInFile("code.txt")

	sm := make(map[string]struct{}, len(s))
	for _, n := range s {
		sm[n] = struct{}{}
	}

	var d []string
	for _, n := range t {
		if _, ok := sm[n]; !ok {
			d = append(d, n)
		}
	}
	fmt.Println("Названия недостающих файлов")
	fmt.Println(d)
}

func getArrayOfFileNamesInPath(path string) (array []string) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		array = append(array, fileNameWithoutExtention(file.Name()))
	}
	return
}

func getArrayOfRowsInFile(filePath string) (array []string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var i int
	for scanner.Scan() {
		if scanner.Text() != "" {
			array = append(array, scanner.Text())
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return
}
func fileNameWithoutExtention(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}
