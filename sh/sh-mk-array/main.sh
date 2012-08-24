#!/bin/sh



for ii in 0 1 2 3 
do
  arr="$arr $ii"
done


for jj in $arr
do
  echo $jj
done