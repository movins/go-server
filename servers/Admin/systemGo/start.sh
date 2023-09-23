#!/bin/bash
set -ex
make
./systemGo --config=config/config.conf
