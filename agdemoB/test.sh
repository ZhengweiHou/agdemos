#!/bin/bash
fname="$1"

for i in {1..10}
do 
    if [ ${i} -lt 5 ]
    then
        echo ${i} check job status: running
    fi

    if [ ${i} -gt 5 ]
    then
        echo ${i} check job status: completed
        break
    fi
done

# 若fname以xxx开头，则报错
# if [ "$fname" = "xxx"* ]
if [[ $fname == *xxx* ]]
then
    echo "fname must not start with xxx"
    exit 1
fi