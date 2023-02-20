package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
)

// SeriallyRenameTask 通过自增编号批量重命名
// arg1 : path
// arg2: 开始文件的编号
func SeriallyRenameTask() {
	if len(os.Args) < 3 {
		fmt.Printf("The app reveived null arguments, it'll exit")
		os.Exit(1)
	}
	filePath := os.Args[1]
	if len(filePath) == 0 {
		fmt.Printf("Invalid input filePath, it'll exit")
		os.Exit(1)
	}
	startNameNumStr := os.Args[2]
	if len(startNameNumStr) == 0 {
		fmt.Printf("startNameNumStr is nil, system is exitig")
		os.Exit(1)
	}
	startNameNum, err := strconv.ParseInt(startNameNumStr, 10, 64)
	if err != nil {
		fmt.Printf("startNameNumStr is not a number, system is exitig")
		os.Exit(1)
	}
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		// 带扩展名的文件名
		fullFilename := f.Name()
		fileType := path.Ext(fullFilename)
		fmt.Println(fullFilename)
		oldFileName := filePath + "\\" + f.Name()
		newFileName := filePath + "\\" + fmt.Sprintf("RECP%d%s", startNameNum, fileType)
		err := os.Rename(oldFileName, newFileName)
		if err != nil {
			fmt.Printf("rename error : oldName[%s],newName[%s],err[%s]", oldFileName, newFileName, err.Error())
		}
		startNameNum++
	}
}
