gmail2go
========

Simple gmail multiple accounts CLI mail checker.

Uses atom feed to get the unread emails. The user/passwords are kept
in an encrypted configuration file. Developed and tested on Ubuntu 12.04

Install
-------

go get github.com/rif/gmail2go

OR

[Download the static binary for ubuntu](https://github.com/downloads/rif/gmail2go/gmail2go_ubuntu12.04_amd64_static.tar.xz)

Usage
-----

    rif@grace:~$ gmail2go -help
    Usage of gmail2go:
        -color=false: use terminal output colors
        -config="$HOME/.gmail2gorc": the user account file
        -account="": adds/updates/deletes to user:password to the accounts file (leave password empty to delete)
        -notify=false: send libnotify message

gmail2go -account user:secret - to create an account

gmail2go -account user:changed_secret - to change password to an account

gmail2go -account user: - to delete an account

Examples
--------

- add a google apps account:

    gmail2go -account test@domain.ro:password

- use terminal colors and libnotify messages:

    gmail2go -color -notify

- run it every 10 minutes with sound notification and libnotify message:

    watch -cn 600 "gmail2go -color -notify && ogg123 -q /usr/share/sounds/gnome/default/alerts/drip.ogg"

Continous integration: [![Build Status](https://goci.herokuapp.com/project/image/github.com/rif/gmail2go "Continous integration")](http://goci.me/project/github.com/rif/gmail2go)
