package file

import (
	"log"
	"testing"
)

func TestNewConfigReader(t *testing.T) {
	config := NewConfigReader("D:/golang/go_work/go-commons/redis-config.yaml")
	//log.Println(len("D:\\config.ini"))
	log.Println(config.FileName, config.FileType, config.Path)

	//YmlFileRead(config.FileName, config.Path, config.FileType)
}
