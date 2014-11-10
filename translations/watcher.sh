#!/usr/bin/bash

set -e

find . -name "*.interface.json" | entr ./update.sh
