package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"src/github.com/ajph/nbclassifier-go"
		"github.com/huichen/sego"
	"net"
	"syscall"
	"flag"
)

type Config struct {
	SockFile string `json:"SOCK_FILE"`
	WorkSpace string `json:"WORKSPACE"`
	Classes []string `json:"CLASSES"`
	DictionaryFile string `json:"DICTIONARYFILE"`
	SocketBufferSize int `json:"SOCKET_BUFFER_SIZE"`
}

type JsonStruct struct {}

func NewJsonParser() *JsonStruct {
	return &JsonStruct{}
}

// 将配置文件加载到config中，需要传一个Config结构体
func (jt *JsonStruct) LoadConfig(filename string, model interface{})  {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = json.Unmarshal(data, model)
	if err != nil {
		log.Fatal(err)
		return
	}
}

// 以换行符读取文件到字符串数组中
func ReadLines(filename string) []string {
    file, err := os.Open(filename)
    if err != nil {
    	log.Fatal(err)
	}
    defer file.Close()
    reader := bufio.NewReader(file)
    lines := []string{}
    for {
    	line, err := reader.ReadString('\n')
    	if err != nil || err == io.EOF {
    		if line == "" {
    			break
			}
		}
    	line = strings.Trim(line, "\n")
    	lines = append(lines, line)
	}
    return lines

}

func TrainModel(config *Config) *nbclassifier.Model{
	model := nbclassifier.New()
	workspace := config.WorkSpace
	if strings.HasSuffix(workspace, "/") == false {
		workspace = workspace + "/"
	}
	// 根据配置生成对应的分类
	classes := config.Classes
	for _, class := range classes {
		filepath := workspace + class + ".txt"
		fmt.Println("filepath:" + filepath)
		words := ReadLines(filepath)
		fmt.Println(words)
		model.NewClass(class)
		model.Learn(class, words...)
	}
	return model
}

func LoadDictionary(dictpath string) sego.Segmenter {
	var segmenter sego.Segmenter
	segmenter.LoadDictionary(dictpath)
	return segmenter
}

func GetClassifyResult(model *nbclassifier.Model, segmenter sego.Segmenter, content string) string {
	// 分词
	segments := segmenter.Segment([]byte(content))
	words := sego.SegmentsToSlice(segments, false)
	fmt.Println("分词结果：", words)
	cls, unsure,_ := model.Classify(words...)
	fmt.Println(content + "> 检测到分类为：" + cls.Id, " unsure:", unsure)

	result := ""
	if unsure == false {
		result = cls.Id
	}
	return result
}

func RunAsSock(config *Config) {
	socket, _ := net.Listen("unix", config.SockFile)
	defer syscall.Unlink(config.SockFile)

	segmenter := LoadDictionary(config.DictionaryFile)
	model := TrainModel(config)

	fmt.Println("spamcheck run as `sock`, the sock file exists at: " + config.SockFile)

	// handle request
	for {
		client, _ := socket.Accept()
		buf := make([]byte, config.SocketBufferSize)
		datalength, _ := client.Read(buf)
		data := buf[:datalength]

		class := GetClassifyResult(model, segmenter, string(data))
		response := []byte("")
		if len(class) > 0 {
			response = []byte(class)
		}
		_, _ = client.Write(response)
	}
}

func RunAsRPC(config *Config) {
	fmt.Println("run spamcheck in rpc way...")
}

func main() {

	configpath := flag.String("c", "", "配置文件绝对路径")
	runway := flag.String("a", "", "以`Unix Domain Socket`或者`RPC`方式运行服务")
	flag.Parse()

	// 检测命令行参数
	if *configpath == "" || strings.HasSuffix(*configpath, ".json") == false {
		fmt.Println("config.json 不能为空~")
		os.Exit(-1)
	}
	if *runway == "" {
		fmt.Println("runway 不能为空")
		os.Exit(-1)
	}

	// 根据配置文件构造配置对象
	JsonParser := NewJsonParser()
	config := &Config{}
	JsonParser.LoadConfig(*configpath, config)
	// choose the way to run this service
	choose := strings.ToLower(*runway)

	if choose == "sock" {
		RunAsSock(config)
	}else if choose == "rpc" {
		RunAsRPC(config)
	}else{
		fmt.Println("暂不支持的服务运行方式：" + choose)
	}



}