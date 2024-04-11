#!/bin/bash

batch_size=10
for ((i=1; i < 100; i++)); do
	second=$(expr 10 + $i)
	echo "i = ${i} start sleep $second s" >> log
	sleep $second 
done | parallel -j $batch_size
