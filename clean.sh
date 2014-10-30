#!/bin/bash

arangosh --javascript.execute tools/db_init.js;
touch kails.go;

cd webapp;
rm -rf dist;
grunt build;
