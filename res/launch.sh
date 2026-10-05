#!/bin/bash

pkill -f hwt

if [ "$#" -gt 0 ]; then
 /mnt/SDCARD/Apps/Nesroom/nesroom "$@"
else
  progdir=$(dirname "$0")
  cd $progdir
  ./nesroom
fi