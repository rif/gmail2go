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
	"path"
	"strings"
)

var (
	config      = flag.String("config", path.Join(os.Getenv("HOME"), ".gmail2gorc"), "the user accounts file")
	set         = flag.String("set", "", "adds/updates/deletes user:password to the accounts file (leave password empty to delete)")
	color       = flag.Bool("color", false, "use terminal output colors")
	accountsMap = make(map[string]string)
)

func main() {
	flag.Parse()
	fin, err := os.Open(*config)
	if err == nil {
		r := bufio.NewReader(fin)
		res, err := passwd.Decrypt(r, make([]byte, 16), make([]byte, 16))
		if err != nil {
			log.Fatal("Could not decrypt accounts file: ", err)
		} else {
			dec := json.NewDecoder(res)
			if err := dec.Decode(&accountsMap); err != nil {
				log.Fatal("Could not decode accounts content to json: ", err)
			}
		}
	}
	if *set != "" {
		fin, err = os.Create(*config)
		if err != nil {
			log.Fatal("Cannot open account file: ", err)
		}
		defer fin.Close()
		w := bufio.NewWriter(fin)
		up := strings.SplitN(*set, ":", 2)
		if len(up) != 2 {
			log.Fatal("Incorrect set string use user:pass")
		}
		if up[1] != "" {
			accountsMap[up[0]] = up[1]
		} else {
			delete(accountsMap, up[0])
		}
		var out bytes.Buffer
		enc := json.NewEncoder(&out)
		err := enc.Encode(accountsMap)
		if err != nil {
			log.Fatal("Could not encode accounts to json", err)
		} else {
			err = passwd.Encrypt(w, &out, make([]byte, 16), make([]byte, 16))
			if err != nil {
				log.Fatal("Could not encrypt accounts json string to file", err)
			}
			w.Flush()
		}
	}

	yellow, green, red := "", "", ""
	if *color {
		yellow, green, red = "\033[1;33m", "\033[0;32m", "\033[0:31m"
	}

	foundAtLeastOne := false
	for user, pass := range accountsMap {
		fmt.Println(yellow+"Account: ", user)
		mails, err := rss.Read("https://mail.google.com/mail/feed/atom", user, pass)
		if err != nil {
			log.Fatal(err)
		}
		for _, m := range mails {
			fmt.Println("\t"+red, m.Title)
			foundAtLeastOne = true
		}
		if len(mails) == 0 {
			fmt.Println("\t" + green + "No unread email.")
		}
	}
	fmt.Println("\033[0m")
	if foundAtLeastOne {
		os.Exit(0)
	}
	os.Exit(1)
}
