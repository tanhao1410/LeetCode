package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	java, rs, goCount := CountNum2("./")
	fmt.Println("java:", java, ",rs:", rs, ",go:", goCount, ",总:", java+rs+goCount)
	fmt.Scanln(&java)
}

//计算目录的行数
func CountNum2(dirName string) (int, int, int) {
	dir, err := os.Open(dirName)
	if err != nil {
		fmt.Println("打开文件夹出错", err)
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("读取文件夹出错", err)
	}
	javaCount, rsCount, goCount := 0, 0, 0
	for _, file := range files {
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), ".java") {
				javaCount += FileLineCount2(dirName + file.Name())
			}
			if strings.HasSuffix(file.Name(), ".rs") {
				rsCount += FileLineCount2(dirName + file.Name())
			}
			if strings.HasSuffix(file.Name(), ".go") {
				goCount += FileLineCount2(dirName + file.Name())
			}
		} else {
			javaCount1, rsCount1, goCount1 := CountNum2(dirName + file.Name() + "/")
			javaCount += javaCount1
			rsCount += rsCount1
			goCount += goCount1
		}
	}
	return javaCount, rsCount, goCount
}

//计算文件的行数
func FileLineCount2(dir string) int {
	openFile, err := os.Open(dir)
	if err != nil {
		fmt.Println("打开文件出错", err)
	}
	defer openFile.Close()
	reader := bufio.NewReader(openFile)
	count := 0
	for line, _, err := reader.ReadLine(); err != io.EOF; line, _, err = reader.ReadLine() {
		if err != nil {
			fmt.Println("读取文件出错", err)
			return 0
		}
		if len(strings.TrimSpace(string(line))) > 0 {
			count++
		}
	}
	return count
}
