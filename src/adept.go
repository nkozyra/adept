package adept

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Conf      Config
	DB        *sql.DB
	Templates *template.Template
)

type Config struct {
	DB   ConfigDB   `json:"database"`
	Host ConfigHost `json:"host"`
}

type ConfigDB struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ConfigHost struct {
	Port string `json:"port"`
}

func LoadTemplates() {
	Templates = template.Must(template.ParseGlob("templates/*.html"))
}

func LoadConfig(f string) {
	var err error
	c, _ := ioutil.ReadFile(f)

	json.Unmarshal(c, &Conf)

	LoadTemplates()

	DB, err = sql.Open("mysql", getConnectString())
	if err != nil {
		fmt.Println(err)
	}
	Serve()
}
