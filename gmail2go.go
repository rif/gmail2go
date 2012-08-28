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
	"os/exec"
	"path"
	"strings"
)

var (
	config      = flag.String("config", path.Join(os.Getenv("HOME"), ".gmail2gorc"), "the user accounts file")
	account     = flag.String("account", "", "adds/updates/deletes user:password to the accounts file (leave password empty to delete)")
	color       = flag.Bool("color", false, "use terminal output colors")
	notify      = flag.Bool("notify", false, "send libnotify message")
	accountsMap = make(map[string]string)
)

func main() {
	flag.Parse()
	fin, err := os.Open(*config)
	if err == nil {
		// if the config file was opened decrypt it and unmarshal accounts map
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
	// if account parameter has a value operate the changes on the map and write the encrypted json 
	if *account != "" {
		fin, err = os.Create(*config)
		if err != nil {
			log.Fatal("Cannot open account file: ", err)
		}
		defer fin.Close()
		w := bufio.NewWriter(fin)
		up := strings.SplitN(*account, ":", 2)
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

	// prepare the colors for display
	yellow, green, red, reset := "", "", "", ""
	if *color {
		yellow, green, red, reset = "\033[1;33m", "\033[0;32m", "\033[0:31m", "\033[0m"
	}

	emailCount := 0
	// iterate over accounts
	for user, pass := range accountsMap {
		fmt.Println(yellow+"Account: ", user)
		mails, err := rss.Read("https://mail.google.com/mail/feed/atom", user, pass)
		if err != nil {
			fmt.Println(reset)
			log.Print("Error: ", err)
			continue
		}
		// iterate over mails
		for _, m := range mails {
			fmt.Println("\t"+red, m.Title)
			emailCount++
		}
		if len(mails) == 0 {
			fmt.Println("\t" + green + "No unread email.")
		}
	}
	fmt.Println(reset)
	// show the notification
	if emailCount > 0 {
		if *notify {
			message := "You have 1 unread mail!"
			if emailCount > 1 {
				message = fmt.Sprintf("You have %v unread mails!", emailCount)
			}
			cmd := exec.Command("/usr/bin/notify-send",
				"-i",
				"/usr/share/notify-osd/icons/gnome/scalable/status/notification-message-email.svg",
				"gmail2go",
				message)
			cmd.Run()
		}
		os.Exit(0)
	}
	os.Exit(1)
}
