#!/usr/bin/env bash
set -Ceu

dirs=`find *.txt -maxdepth 0 -type f`

go build -o exe
for dir in ${dirs}; do
	echo "${dir}"
>   cat ${dir} | ./exe
	echo -e "\n\n"
done
