#!/bin/sh

find . -regextype egrep -regex '(.*\.go|.*\.tmpl\.html|.*\.all\.json)' | entr -r sh -c "killall kails; go run kails.go"
