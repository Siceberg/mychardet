package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"

	"gitlab.com/siceberg/chardet"
)

func main() {
	// fileName := "/Users/leo/Downloads/gb2312test.csv"
	// fileName := "/Users/leo/Downloads/test.csv"
	// fileName := "/Users/leo/Downloads/utf8test.csv"
	fileName := "/Users/leo/Downloads/651770-Dataupdate.csv"
	// 打开csv文件
	if csvFileByte, err := ioutil.ReadFile(fileName); err != nil && err != io.EOF {
		log.Printf("Open CSV File fail: %+v", err)
	} else {
		log.Println(chardet.Mostlikein(csvFileByte, []string{"utf8", "gbk", "gb18030"}))

		csvFile, err := os.Open(fileName)
		if err != nil {
			log.Printf("Open CSV File fail: %+v", err)
		} else {
			// TODO 判断字符集后决定是否使用chardet.NewReader
			charsetReader, err := chardet.NewReader(bufio.NewReader(csvFile), "gbk", csvFileByte)
			if err != nil {
				log.Println("test!!!")
			} else {
				reader := csv.NewReader(charsetReader)
				for {
					line, error := reader.Read()
					if error == io.EOF {
						break
					} else if error != nil {
						log.Fatal(error)
					}
					log.Println(line)
				}
			}
		}
	}

	// csvFile, err := os.Open("/Users/leo/Downloads/gb2312test.csv")
	// csvFile, err := os.Open("/Users/leo/Downloads/test.csv")
	// csvFile, err := os.Open("/Users/leo/Downloads/utf8test.csv")
	// if err != nil {
	// 	log.Printf("Open CSV File fail: %+v", err)
	// } else {
	// 	reader := csv.NewReader(bufio.NewReader(csvFile))
	// 	for {
	// 		line, error := reader.Read()
	// 		if error == io.EOF {
	// 			break
	// 		} else if error != nil {
	// 			log.Fatal(error)
	// 		}
	// 		// log.Println(line)
	// 		log.Println(chardet.Mostlikein([]byte(strings.Join(line, " ")), []string{"utf8", "gbk", "gb18030"}))
	// 	}
	// }
}
