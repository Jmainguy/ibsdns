package main

import (
	"fmt"
    "os"
    "flag"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {

    value := flag.String("value", "", "IP to update domain to, example 192.168.1.1")
    flag.Parse()

	c := config()
	recordtype := "A"
	newvalue := *value

    if newvalue == "" {
        fmt.Println("You must provide a value for the IP")
        os.Exit(1)
    }

	tid, status, message := updateDns(c.ApiKey, c.Password, c.Domain, recordtype, newvalue)
	if message != "" {
		fmt.Printf("TransactID: %s, Status: %s, Message: %s\n", tid, status, message)
	} else {
		fmt.Printf("TransactID: %s, Status: %s\n", tid, status)
	}

}
