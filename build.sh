#!/bin/bash
APP_NAME=`pwd`
PROJECT_PATH=$(cd `dirname $0`; pwd)
PROJECT_NAME="${PROJECT_PATH##*/}"
echo "go build -o "${PROJECT_NAME}
go build -o ${PROJECT_NAME} ${APP_NAME}
if [ -f ${PROJECT_NAME} ]; then
  echo "mv "${PROJECT_NAME}" to "${APP_NAME}/bin
  mv ${PROJECT_NAME} ${APP_NAME}/bin;
fi



