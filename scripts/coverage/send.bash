#!/bin/bash

service_name=wercker
profile_path=coverprofile/gover.coverprofile

goveralls -service=${service_name} -repotoken=$COVERALLS_TOKEN -coverprofile=${profile_path}
