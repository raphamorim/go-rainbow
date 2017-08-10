#!/bin/sh
for i in $(seq 0 255)
do
    printf '\033[1;38;5;%dmcolor %d\n' "$i" "$i"
done
