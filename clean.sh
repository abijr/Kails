#!/bin/bash

arangosh --javascript.execute tools/db_init.js;
touch kails.go;
