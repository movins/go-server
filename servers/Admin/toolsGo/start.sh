#!/bin/bash
set -ex
make
./toolsGo --config=config/config.conf
