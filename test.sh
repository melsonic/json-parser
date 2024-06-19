#!/bin/bash

for file in tests/*
do 
  if [ -d $file ]
  then 
    for rfile in $file/*
    do
      echo "$rfile"
      ./jp $rfile
      echo ""
    done
  else
    echo "$file"
    ./jp $file
    echo ""
  fi
done

