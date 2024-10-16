#!/bin/bash

tempfile=$(mktemp)

find . -type f | xargs -I {} echo 'shasum -a 256 "'{}'" >> '$tempfile

echo 'cat '$tempfile' | sort > shasums.txt'

echo '# this script is intented to create a shell script that will create a shasum file for all files in the current directory and subdirectories'

echo '# save the output of this script to a file and run the file to create the shasum file'
