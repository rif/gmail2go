gmail2go
========

Simple gmail multiple accounts cli mail checker.

Uses atom feed to get the unread emails. The user/passwords are kept in an encrypted configuration file.

Install
-------

go get github.com/rif/gmail2go

Usage
-----

rif@grace:~$ gmail2go -help
Usage of gmail2go:
  -config="$HOME/.gmail2gorc": the user account file
  -set="": adds/updates/deletes to user:password to the accounts file (leave password empty to delete)

gmail2go -set user:secret - to create an account

gmail2go -set user:changed_secret - to change password to an account

gmail2go -set user: - to delete an account

Example: run it every 5 minutes with sound notification

watch -n 600 "gmail2go && play -q /usr/share/sounds/gnome/default/alerts/drip.ogg"

Continous integration: [![Build Status](https://goci.herokuapp.com/project/image/github.com/rif/gmail2go "Continous integration")](http://goci.me/project/github.com/rif/gmail2go)
