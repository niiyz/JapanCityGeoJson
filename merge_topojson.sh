#!/bin/bash

find ./topojson/47都道府県 -type f | while read file
do
    arr=( `echo $file | tr -s '/' ' '`)
    pref=${arr[2]}
    dir="topojson/"$pref

    basename=$(basename $file)
    if [ $basename = '.DS_Store' -o ! ${basename##*.} = '.simple.topojson' ]; then
      continue
    fi

    name=( `echo $basename | sed -e "s/\.topojson//"`)
    echo "merge: "$pref" - "$name

    # merge
    $(npm bin)/topomerge $name=$name -k d.id < topojson/$pref/$basename > topojson/$pref/$name.merge.topojson

done