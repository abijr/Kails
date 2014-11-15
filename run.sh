#!/bin/sh

set -e

(
    cd translations;
    ./watcher.sh;
) &

(
    find . -regextype egrep -regex '(.*\.go|.*\.tmpl\.html|.*\.all\.json)' |
    entr -r sh -c "echo 'kails: killing process...'; killall kails; go run kails.go"
) &

(
    cd webapp;
    grunt;
)
