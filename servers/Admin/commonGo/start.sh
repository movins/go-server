#!/bin/bash
set -ex
make
./commonGo --config=config/config.conf
