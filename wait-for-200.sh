#!/bin/sh

curl --retry 100 --retry-delay 3 --retry-max-time 30 --retry-connrefused $1 > /dev/null
