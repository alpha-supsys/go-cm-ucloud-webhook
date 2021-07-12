#!/bin/sh

version=$1

docker build -t ${hub}/alpha_supsys/go-cm-ucloud-webhook:${version} .