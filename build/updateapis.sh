#!/bin/bash
##############################################################
# Shell script to update the API codes to the latest
# versions, including downloading the API and
# generating the go library.
#
# Requires `curl` in order to download the API discovery doc
# remote repository. There are no arguments to this script,
# so run using:
# 
# updateapi.sh
##############################################################

CURRENT_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
CURL=`which curl`
GO=`which go`

##############################################################

cd "${CURRENT_PATH}/.."

##############################################################
# Sanity checks

if [ ! -d ${CURRENT_PATH} ] ; then
  echo "Not found: ${CURRENT_PATH}" >&2
  exit -1
fi
if [ ! -x ${GO} ] ; then
  echo "go not installed or executable" >&2
  exit -1
fi
if [ ! -x ${CURL} ] ; then
  echo "curl not installed or executable" >&2
  exit -1
fi

##############################################################
# Fetch API's

# Get the API generator from remote source and build it
go get -u google.golang.org/api/google-api-go-generator || exit 1

# Get YouTube Data API
go get -u google.golang.org/api/youtube/v3

# Create the videointelligence folders
install -d googleapi/videointelligence/v1 || exit 1
install -d googleapi/videointelligence/v1beta2 || exit 1

#  Fetch the API documents
curl https://videointelligence.googleapis.com/\$discovery/rest?version=v1 > googleapi/videointelligence/v1/videointelligence.json || exit 1
curl https://videointelligence.googleapis.com/\$discovery/rest?version=v1beta2 > googleapi/videointelligence/v1beta2/videointelligence.json || exit 1

##############################################################
# Generate Go Libraries

API_GEN=`which google-api-go-generator`
if [ ! -x ${API_GEN} ] ; then
  echo "google-api-go-generator not installed or executable" >&2
  exit -1
fi

# Generate the go libraries
google-api-go-generator -cache=false -gendir=googleapi -api_json_file=googleapi/videointelligence/v1/videointelligence.json -api_pkg_base=github.com/djthorpe/EnterpriseTechnologyDay/googleapi
google-api-go-generator -cache=false -gendir=googleapi -api_json_file=googleapi/videointelligence/v1beta2/videointelligence.json -api_pkg_base=github.com/djthorpe/EnterpriseTechnologyDay/googleapi

