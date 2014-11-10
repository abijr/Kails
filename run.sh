#!/bin/sh

cd translations;
./update.sh;
cd -;

(find . -regextype egrep -regex '(.*\.go|.*\.tmpl\.html|.*\.all\.json)' | entr -r sh -c "killall kails; go run kails.go") &

cd webapp;
grunt;
cd -;
