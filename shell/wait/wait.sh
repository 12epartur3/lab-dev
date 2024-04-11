#!/bin/bash

batch_size=20
for ((i=1; i < 100; i++)); do
    if (( $i % $batch_size == 0 )); then
        echo "$i is divisible by 20"
	sleep 3 &
	wait
    elif (( $i % 10 == 0 )); then
        echo "$i is divisible by 10"
	sleep 10 &
    else
        echo "i = $i"
    fi
done
