#!/bin/bash

# constants

AMD64="amd64"
ARM64="arm64"
LINUX="linux"

DEFAULT_OS="linux"
DEFAULT_ARCH="amd64"

NULL = ""

BLUE=34
CYAN=36
GREEN=32
RED=31

# defaults

SRC="."
dst="debian"
arch=""
os="linux"

# helpers

# text, color
color()
{
  echo "\e[$2m$1\e[0m"
}

timestamp()
{
#  t="$date %T"
#  echo $t
#  echo "$(color $(t) $RED)"
  echo "$(date)"
}

error()
{
  echo -e "[$(timestamp)] \e[31m$1\e[0m"
}


log()
{
  echo -e "[$(timestamp)] \e[32m$1\e[0m"
#echo -e "$(timestamp) $(color $1 $CYAN)"
}

# usage
usage()
{
  error "Usage: please set env variables GOOS and GOARCH"
  exit
}

# check

check()
{

  log "Checking env..."

  if [ -z $GOOS ]; then
    log "GOOS env variable not set, using $DEFAULT_OS"
    GOOS=$DEFAULT_OS
  fi

  if [ -z $GOARCH ]; then
    log "GOARCH env variable not set, using $DEFAULT_ARCH"
    GOARCH=$DEFAULT_ARCH
  fi

  if ! [[ $GOOS == $LINUX ]]; then
    log "GOOS not supported: $GOOS, only $LINUX is supported."
  fi

  if ! [[ $GOARCH == $AMD64 || $GOARCH == ARM64 ]]; then
    log "GOARCH not supported: $GOARCH, only $AMD64 and $ARM64 are supported."
  fi

}

# clean

clean()
{
  log "clean"
}

# build:  GOOS, GOARCH

build()
{

  go build
  #migrate -database sqlite3://rpic.db -path db/migrations up
}

# init_database

init_database()
{
  log "Initializing database"
}

# stage

stage()
{
  cp rpic debian/usr/local/rpic/
  cp -R www debian/usr/local/rpic/
  cp rpic.db debian/usr/local/rpic/
}

# package

package()
{
  dpkg-deb --build debian
}

log "Build started..."
check
build
init_database