package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	help = `
This is a simple app, to check if mysql is available and contain table
You can pass all needed information to app by using ENV. 
Below you will find ENV and their default value 

ENV:
	APP_CHECK_TRIES="10"
	MYSQL_PORT="3306"
	MYSQL_HOST="127.0.0.1"
	MYSQL_USER="user"
	MYSQL_PASSWD="passwd"
	MYSQL_DATABASE="appdb"
	MYSQL_CHECK_TABLE="application"

Flags:
	-h	print help information

`
)

type vars struct {
	App_chek_tries int
	Mysql_port     string
	Mysql_host     string
	Mysql_user     string
	Mysql_passwd   string
	Mysql_database string
}

func setdefaultvalue(env string, defvalue string) string {
	if value, ok := os.LookupEnv(env); ok {
		env = value
	} else {
		env = defvalue
	}
	return env
}

func newvars() *vars {

	app_check_tries, err := strconv.Atoi(setdefaultvalue("APP_CHECK_TRIES", "10"))
	if err != nil {
		fmt.Println("APP_CHECK_TRIES env variable must be a int")
		os.Exit(1)
	}
	mysql_port := setdefaultvalue("MYSQL_PORT", "3306")
	mysql_host := setdefaultvalue("MYSQL_HOST", "127.0.0.1")
	mysql_user := setdefaultvalue("MYSQL_USER", "user")
	mysql_passwd := setdefaultvalue("MYSQL_PASSWD", "passwd")
	mysql_database := setdefaultvalue("MYSQL_DATABASE", "appdb")

	return &vars{
		App_chek_tries: app_check_tries,
		Mysql_port:     mysql_port,
		Mysql_host:     mysql_host,
		Mysql_user:     mysql_user,
		Mysql_passwd:   mysql_passwd,
		Mysql_database: mysql_database,
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "%s", help)
	os.Exit(0)
}

func check(v *vars) {
	c := 0
	fmt.Println("try to connect to mysql")

	creds := v.Mysql_user + ":" + v.Mysql_passwd + "@tcp(" + v.Mysql_host + ":" + v.Mysql_port + ")/" + v.Mysql_database

	db, err := sql.Open("mysql", creds)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		var check string

		for {
			err := db.QueryRow("SELECT 1").Scan(&check)

			if err != nil {

				fmt.Printf("try %d, sleep 5s\n", c)
				c += 1

				if c == v.App_chek_tries {
					fmt.Println("mysql database not available or database not correct")
					os.Exit(1)
				}

				time.Sleep(5 * time.Second)

			} else {
				db.Close()
				fmt.Println("mysql is available")
				wg.Done()
				break
			}
		}
	}
}

func main() {

	flag.Usage = usage
	flag.Parse()

	v := newvars()

	wg.Add(1)
	go check(v)
	wg.Wait()

	fmt.Printf("mysql is ready to connect\n")

	os.Exit(0)
}
