package conf

import (
	"encoding/json"
	"io/ioutil"
	"github.com/blind-platform/pkg/util"
)

type AwsDB struct {
	User		string `json:"User"`
	Password	string `json:"Password"`
	Endpoint	string `json:"Endpoint"`
	DBName		string `json:"DBName"`
	Region		string `json:"Region"`
}

type RSA struct {
	PathPrivate		string `json:"PathPrivate"`
	PathPublic		string `json:"PathPublic"`
}

type Conf struct {
	AwsDB	AwsDB	`json:"AwsDB"`
	RSA		RSA		`json:"RSA"`
}

func Load() (Conf) {
	var conf Conf

	data, err := ioutil.ReadFile("cfg/conf.json")
	util.Fatal(err)
	err = json.Unmarshal(data, &conf)
	util.Fatal(err)
	return conf
}