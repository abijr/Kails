#!/usr/bin/bash

set -e

./update.sh

find . -name "*.interface.json" | entr ./update.sh
