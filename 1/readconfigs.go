package mailfetcher

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

//TagClassInfo 存储配置信息
type TagClassInfo struct {
	className    string
	homeworkPath string
	mailserver   string
	mailUser     string
	mailPassword string
	prefixFlag   string
	stuLists     []string
	DateStart    time.Time
	DateEnd      time.Time
	VIOLATELIST  []string
}

//ClassName 班级名称
func (clsInfo *TagClassInfo) ClassName() string {
	return clsInfo.className
}

// readConfig Read config options form txt file
// readConfig从txt文件读取配置选项
func readConfig(txtPath string, classConfigs *[]TagClassInfo) {
	var currentClassInfo TagClassInfo

	//fmt.Println("read configs from: ", txtPath)

	file, err := os.Open(txtPath) //打开文件
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	buf := bufio.NewReader(file) //新建一个缓冲区，把内容先放在缓冲区

	//读取配置信息
LabelReadOptions:
	for {
		line, err := buf.ReadString('\n') //以'\n' 为一行，读取文件内容。
		line = strings.TrimSpace(line)
		//会返回一个string类型的slice，并将最前面和最后面的ASCII定义的空格去掉，中间的空格不会去掉，如果遇到了\0等其他字符会认为是非空格。

		if line == "" {
			break LabelReadOptions
		}

		//解析配置key和value
		keyvalue := strings.TrimSpace(strings.SplitN(line, "#", 2)[0]) //以"#"分割字符串，[0]带代表前半段 [1]后半段
		key := strings.SplitN(keyvalue, "=", 2)[0]
		value := strings.SplitN(keyvalue, "=", 2)[1]

		switch key {
		case "homework_path":
			currentClassInfo.homeworkPath = value
		case "mailserver":
			currentClassInfo.mailserver = value
		case "mail_user":
			currentClassInfo.mailUser = value
		case "mail_passwd":
			currentClassInfo.mailPassword = value
		case "prefix_flag":
			currentClassInfo.prefixFlag = value
			currentClassInfo.className = value
		default:
			fmt.Println("Unknown: ", key, value)
		}

		if err != nil {
			log.Fatal(err)
		}
	}

LabelStudents:
	//读取学员名单信息
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "" {
			break LabelStudents
		}

		//初始化学员列表和违纪列表
		currentClassInfo.stuLists = append(currentClassInfo.stuLists, line)
		currentClassInfo.VIOLATELIST = append(currentClassInfo.VIOLATELIST, line)

		if err != nil {
			if err == io.EOF {
				break LabelStudents
			}
			log.Fatal(err)
		}
	}

	*classConfigs = append(*classConfigs, currentClassInfo)
	//fmt.Println(txtPath, "done.")
}

//ReadConfigDir Read configs from config txt files
func ReadConfigDir(configpath string) []TagClassInfo {
	var configTxtFiles []string
	var classConfigs []TagClassInfo

	if configpath == "" {
		configpath = "./"
	}

	files, err := ioutil.ReadDir(configpath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "NR") && strings.HasSuffix(file.Name(), "txt") {
			configTxtFiles = append(configTxtFiles, file.Name())
		}
	}

	for _, item := range configTxtFiles {
		readConfig(path.Join(configpath, item), &classConfigs)
		//Join函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加斜杠。结果是经过简化的，所有的空字符串元素会被忽略。
	}

	return classConfigs
}
