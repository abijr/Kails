#!/usr/bin/bash

set -e

goi18n -outdir all/ {english,spanish}/*.json


(
    cd all;
    echo -e "Generating translation files...";
    mv en-US.all.json en-us.all.json
    mv es-MX.all.json es-mx.all.json
    echo -e "Done";
)
