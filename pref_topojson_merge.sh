#!/bin/bash

dir="topojson/47都道府県_単体"
if [ ! -d $dir ]; then
    mkdir -p $dir
fi

find ./topojson/47都道府県 -type f | while read file
do
    basename=$(basename $file)
    prefname=( `echo $basename | sed -e "s/\.topojson//"`)
    echo $prefname
    $(npm bin)/topomerge $prefname=$prefname < ./topojson/47都道府県/$prefname.topojson > $dir/$prefname.json
done