package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/rif/gmail2go/passwd"
	"github.com/rif/gmail2go/rss"
	"log"
	"os"
	"strings"
)

var (
	accountFile = flag.String("account-file", "test.gmail2go", "the user account filee to be used")
	set         = flag.String("set", "", "adds/updates/deletes au user:password to the accounts file")
	accountsMap = make(map[string]string)
)

func main() {
	flag.Parse()
	fin, err := os.Open(*accountFile)
	if err != nil {
		if *set != "" {
			fin, err = os.Create(*accountFile)
			defer fin.Close()
			w := bufio.NewWriter(fin)
			up := strings.SplitN(*set, ":", 2)
			if len(up) != 2 {
				log.Fatal("Incorrect set string use user:pass")
			}
			accountsMap[up[0]] = up[1]
			var out bytes.Buffer
			enc := json.NewEncoder(&out)
			err := enc.Encode(accountsMap)
			if err != nil {
				log.Fatal("Could not encode accounts to json", err)
			}
			err = passwd.Encrypt(w, &out, make([]byte, 16), make([]byte, 16))
			if err != nil {
				log.Fatal("Could not encrypt accounts json string to file", err)
			}
		}
		if err != nil {
			log.Fatal("Cannot open account file: ", err)
		}
	} else {

		r := bufio.NewReader(fin)
		res, err := passwd.Decrypt(r, make([]byte, 16), make([]byte, 16))
		if err != nil {
			log.Fatal("Could not decrypt accounts file: ", err)
		}
		dec := json.NewDecoder(res)
		if err := dec.Decode(&accountsMap); err != nil {
			log.Fatal("Could not decode accounts content to json: ", err)
		}
	}

	if *set != "" {
		//
	}

	for user, pass := range accountsMap {
		fmt.Println("Account: ", user)
		mails, err := rss.Read("https://mail.google.com/mail/feed/atom", user, pass)
		if err != nil {
			log.Fatal(err)
		}
		for _, m := range mails {
			fmt.Println("\t", m.Title)
		}
	}
}
