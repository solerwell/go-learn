package main

import (
	"fmt"
	"github.com/solerwell/go-learn/logger"
	"github.com/solerwell/go-learn/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

var log = logger.GetStdLogger()
var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	testFindConfig()
}

func getRootPath() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return path
}
func testFindConfig() {
	rootPath := util.GetProjectRootPath()
	if len(rootPath) == 0 {
		fmt.Println("find root path error ")
	}
	sep := os.PathSeparator
	path := fmt.Sprintf("%s%cyanyue_spider%cdata%cyanyue.sql", rootPath, sep, sep, sep)
	fmt.Println("config path: ", path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("read config file error: %s", err.Error())
	}
	log.Infof("the sql content: %s", content)
}
