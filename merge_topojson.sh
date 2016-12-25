#!/bin/bash

find ./topojson/47都道府県 -type f | while read file
do
    arr=( `echo $file | tr -s '/' ' '`)
    pref=${arr[2]}
    dir="topojson/"$pref

    basename=$(basename $file)
    if [ $basename = '.DS_Store' -o ! ${basename##*.} = 'topojson' ]; then
      continue
    fi

    name1=( `echo $basename | sed -e "s/\.topojson//"`)

    if [ ! ${name1##*.} = 'simple' ]; then
        continue
    fi

    name2=( `echo $name1 | sed -e "s/\.simple//"`)

    echo "merge: "$pref" - "$name2

    # merge
    $(npm bin)/topomerge $name2=$name2 -k d.id -o topojson/$pref/$name2.merge.topojson topojson/$pref/$basename

done