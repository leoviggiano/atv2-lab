#!/bin/bash

fails_dir="./fails"
mv "$fails_dir/repositories.txt" "$fails_dir/repositories-tmp.txt"
repo_list="$fails_dir/repositories-tmp.txt"

if [ ! -f "$repo_list" ]
then
	echo "File $repo_list does not exists"
	exit 1
fi

while IFS= read -r line
do
	rm -rf "$line"
done < "$fails_dir/metrics.txt"

echo "Error metrics folder cleared with success"
rm -rf "$fails_dir/metrics.txt"

./clone.sh $repo_list

rm -rf "$fails_dir/repositories-tmp.txt"
