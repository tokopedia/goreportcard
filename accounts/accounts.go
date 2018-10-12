package accounts

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type account struct {
	Username string
	Password string
}

var Account account

func SetAccount() {
	account := account{}

	path := "/etc/goreportcard/cred"
	readFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	creds := strings.Split(string(readFile), ":")
	account.Password = creds[1]
	account.Username = creds[0]
	Account = account
}
