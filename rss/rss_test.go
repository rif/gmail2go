package rss

import (
	"testing"
	"time"
)

var (
	text = []byte(`
<?xml version="1.0" encoding="UTF-8"?>
<feed version="0.3" xmlns="http://purl.org/atom/ns#">
<title>Gmail - Inbox for zzz@gmail.com</title>
<tagline>New messages in your Gmail Inbox</tagline>
<fullcount>29</fullcount>
<link rel="alternate" href="http://mail.google.com/mail" type="text/html" />
<modified>2012-08-26T09:51:02Z</modified>
<entry>
<title>Confirming Your Registration (&gt;CBGQ077859&lt;)</title>
<summary>Your registration information has been received and saved. Thank you for registering with Cabelas.com</summary>
<link rel="alternate" href="http://mail.google.com/mail?account_id=ustest@gmail.com&amp;message_id=138e13f5b7a4bad0&amp;view=conv&amp;extsrc=atom" type="text/html" />
<modified>2012-08-01T08:13:42Z</modified>
<issued>2012-08-01T08:13:42Z</issued>
<id>tag:gmail.google.com,2004:1409085679482485456</id>
<author>
<name>cabelasnews</name>
<email>cabelasnews@cabelas.com</email>
</author>
</entry>
<entry>
<title>Enjoy 25% off new Save the Dates, Baby Announcements, and Invitations</title>
<summary>View email online Snapfish by HP Shop | Photos | Special Offers Celebrate every new beginning | 25%</summary>
<link rel="alternate" href="http://mail.google.com/mail?account_id=ustest@gmail.com&amp;message_id=138d82412dfd293b&amp;view=conv&amp;extsrc=atom" type="text/html" />
<modified>2012-07-30T13:47:19Z</modified>
<issued>2012-07-30T13:47:19Z</issued>
<id>tag:gmail.google.com,2004:1408925474892884283</id>
<author>
<name>Snapfish</name>
<email>snapfish@emails.snapfish.com</email>
</author>
</entry>
</feed>
`)
)

func TestParse(t *testing.T) {
	entries, err := unmarshal(text)
	if err != nil || len(entries) != 2 {
		t.Error("Error parsing feed: ", entries, err)
	}
}

func TestModifiedTime(t *testing.T) {
	entries, _ := unmarshal(text)
	mt, err := entries[0].ModifiedTime()
	if err != nil || mt != time.Date(2012, 8, 1, 8, 13, 42, 0, time.UTC) {
		t.Error("Error parsing date: ", mt, err)
	}
}

func TestAuthor(t *testing.T) {
	entries, _ := unmarshal(text)
	if entries[1].Author.Name != "Snapfish" || entries[1].Author.Email != "snapfish@emails.snapfish.com" {
		t.Error("Error parsing author: ", entries[1].Author)
	}
}
