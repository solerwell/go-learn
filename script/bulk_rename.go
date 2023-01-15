package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

// 批量重命名
func bulkNameTask() {
	if len(os.Args) <= 1 {
		fmt.Printf("The app reveived null arguments, it'll exit")
		os.Exit(1)
	}
	path := os.Args[1]
	if len(path) == 0 {
		fmt.Printf("Invalid input path, it'll exit")
		os.Exit(1)
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	fileNum := len(files)
	var fileMap = make(map[string]fs.FileInfo, fileNum)
	for _, file := range files {
		fileMap[file.Name()] = file
	}
	var count int
	for _, f := range fileMap {
		count++
		// 带扩展名的文件名
		fullFilename := f.Name()
		fmt.Println(fullFilename)
		err := os.Rename(path+"\\"+f.Name(), path+"\\"+fmt.Sprintf("%d-%s", count, fullFilename))
		if err != nil {
			fmt.Printf("rename error : %s", err.Error())
		}
	}
}
