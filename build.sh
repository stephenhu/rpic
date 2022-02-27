#!/bin/bash

# constants

AMD64="amd64"
ARM="arm"
ARM64="arm64"
LINUX="linux"

# colors

BLUE=34
CYAN=36
GREEN=32
RED=31

# defaults

DEFAULT_OS="linux"
DEFAULT_ARCH="amd64"

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
  exit
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

  if ! [[ $GOARCH == $AMD64 || $GOARCH == ARM || $GOARCH == ARM64 ]]; then
    log "GOARCH not supported: $GOARCH, only $AMD64, $ARM, and $ARM64 are supported."
  fi

  go version > /dev/null 2>&1

  if [[ $? -ne 0 ]]; then
    error "golang has not been installed, please see https://golang.org."
  fi

  migrate -version > /dev/null 2>&1

  if [[ $? -ne 0 ]]; then
    error "golang-migrate/migrate has not been installed, please see https://github.com/golang-migrate/migrate"
  fi

  if ! [ -d debian/usr/local/rpic ]; then
    mkdir -p debian/usr/local/rpic
  fi

}

# clean

clean()
{
  log "Cleaning environment"

  if [[ -f "rpic" ]]; then
    log "Removing rpic binary"
    rm rpic
  fi

  if [[ -f "rpic.db" ]]; then
    log "Removing rpic.db"
    rm rpic.db
  fi

  if [[ -f "debian.deb" ]]; then
    log "Removing debian.deb packages"
    rm debian.deb
  fi

}

# build:  GOOS, GOARCH

build()
{
  go build
}

# init_database

init_database()
{
  log "Initializing database"
  migrate -database sqlite3://rpic.db -path db/migrations up
}

# stage

stage()
{
  
  log "Copying rpic binary to debian for staging"
  cp rpic debian/usr/local/rpic/

  log "Copying rpic web templates to debian for staging"
  cp -R www debian/usr/local/rpic/

  log "Copying rpic database to debian for staging"
  cp rpic.db debian/usr/local/rpic/

}

# package

package()
{
  log "Building debian package"
  dpkg-deb --build debian

  if [[ -f debian.deb ]]; then

    log "Renaming debian package..."

    mv debian.deb "rpic-0.1-$GOOS-$GOARCH.deb"

  fi

}

log "Build started..."
clean
check
build
init_database
stage
package
