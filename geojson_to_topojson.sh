#!/bin/bash

find ./geojson -type f | while read file
do
    arr=( `echo $file | tr -s '/' ' '`)
    pref_code=${arr[2]}
    # ディレクトリ
    dir="topojson/"$pref_code
    if [ ! -d $dir ]; then
      mkdir -p $dir
    fi
    basename=$(basename $file)
    name=( `echo $basename | sed -e "s/\.json/.topojson/"`)
    echo $name
    npx geo2topo -o $dir/$name -q 1e4 geojson/$pref_code/$basename
done