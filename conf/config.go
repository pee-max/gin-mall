package conf

import (
	"gin_mall/dao"
	"gopkg.in/ini.v1"
	"strings"
)

var (
	AppMode  string
	HttpPort string

	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string

	R2Accesskey string
	R2Secretkey string
	R2Bucket    string
	R2AccountId string
	R2Domain    string

	Host        string
	ProductPath string
	AvatarPath  string
)

func Init() {
	file, err := ini.Load("conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadService(file)
	LoadDB(file)
	LoadOs(file)
	LoadSmtp(file)
	LoadPath(file)

	//8
	pathRead := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True"}, "")
	//2
	pathWrite := strings.Join([]string{DbUser, ":", DbPassword, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True"}, "")
	dao.Database(pathRead, pathWrite)
}

func LoadService(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
}

func LoadDB(file *ini.File) {
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassword = file.Section("mysql").Key("DbPassword").String()
	DbName = file.Section("mysql").Key("DbName").String()
	DB = file.Section("mysql").Key("DB").String()
}

func LoadOs(file *ini.File) {
	R2Accesskey = file.Section("os").Key("R2_AccessKey").String()
	R2Secretkey = file.Section("os").Key("R2_SecretKey").String()
	R2Bucket = file.Section("os").Key("R2_Bucket").String()
	R2AccountId = file.Section("os").Key("R2_AccountId").String()
	R2Domain = file.Section("os").Key("R2_Domain").String()
}

func LoadSmtp(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").String()
	SmtpHost = file.Section("email").Key("SmtpHost").String()
	SmtpEmail = file.Section("email").Key("SmtpEmail").String()
	SmtpPass = file.Section("email").Key("SmtpPass").String()
}

func LoadPath(file *ini.File) {
	AvatarPath = file.Section("path").Key("AvatarPath").String()
	ProductPath = file.Section("path").Key("ProductPass").String()
	Host = file.Section("path").Key("Host").String()
}
