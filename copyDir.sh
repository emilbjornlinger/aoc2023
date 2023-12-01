#!/bin/bash

cd "days"
for i in {3..25}
do
    cp -r "day2" "./day$i"
    cd "day$i"
    replace="day$i"
    sed -i "s/dayX/$replace/" "puzzle.go"
    cd ..
done
