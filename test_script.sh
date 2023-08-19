#!/bin/bash -e
curl --header "Content-Type: application/json" \
	--request POST \
	--data '{"shizbot":"indeed"}' \
	http://localhost:8080/
