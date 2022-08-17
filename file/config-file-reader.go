package file

import (
	"github.com/spf13/viper"
	"log"
	"path"
	"strings"
)

const (
	ConfigFileName = "redis-config"
)

type ConfigFileReader interface {
	YmlFileRead()
	JsonFileRead()
	PropertiesFileRead()
	XmlFileRead()
	IniFileRead()
}

type configReader struct {
	FileName string
	FileType string
	Path     string
}

func NewConfigReader(filepath string) *configReader {
	filename, fileType, dir := resolveFilename(filepath)
	log.Println(filename, fileType, dir)
	return &configReader{FileName: filename, FileType: fileType, Path: dir}
}

func resolveFilename(filepath string) (filename, fileType, dir string) {
	if strings.Contains(filepath, "\\") {
		index := strings.LastIndex(filepath, "\\")
		filename = filepath[index+1:]
		fileType = filepath[strings.LastIndex(filepath, "."):]
		dir = filepath[:index]
	} else {
		filename = path.Base(filepath)
		fileType = path.Ext(filepath)
		dir = path.Dir(filepath)
	}
	//filename = filename[:len(filename)-len(fileType)]
	fileType = fileType[1:]
	return
}

func (c *configReader) YmlFileRead(filename, path, fileType string) {
	//viper.AddConfigPath(c.Path)
	//log.Println(c.Path, c.FileName, c.FileType)
	viper.AddConfigPath("D:/golang/go_work/go-commons")
	viper.SetConfigFile(filename)
	viper.SetConfigType(fileType)
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("err : %s", err.Error())
		panic("read config error ")
	}
	host := viper.Get("redis-client.host")
	database := viper.Get("redis-client.database")
	pwd := viper.Get("redis-client.pwd")
	log.Println(host, database, pwd)
}

func (c *configReader) JsonFileRead() {
	//TODO implement me
	panic("implement me")
}

func (c *configReader) PropertiesFileRead() {
	//TODO implement me
	panic("implement me")
}

func (c *configReader) XmlFileRead() {
	//TODO implement me
	panic("implement me")
}

func (c *configReader) IniFileRead() {
	//TODO implement me
	panic("implement me")
}
