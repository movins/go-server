#!/bin/bash
set -ex
make
./monitorGo --config=config/config.conf
