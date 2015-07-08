#!/bin/bash

base_package=github.com/mitsuse/matrix-go
base_path=${GOPATH}/src/${base_package}

if [ ! -d ${base_path}/coverprofile ]
then 
    mkdir ${base_path}/coverprofile
else
    rm ${base_path}/coverprofile/*.coverprofile
fi

for package in $(go list ${base_package}/...)
do
    cover_name=$(echo ${package} | sed -e "s/\//__/g").coverprofile
    cover_path=${base_path}/coverprofile/${cover_name}
    go test -covermode=count -coverprofile=${cover_path} ${package}
done

gover ${base_path}/coverprofile coverage.txt
