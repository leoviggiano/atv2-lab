#!/bin/bash
repo_list=$1
echo $repo_list
if [ ! -f "$repo_list" ]
then
	echo "File $repo_list does not exists"
	exit 1
fi

while IFS= read -r line
do
	basename=$(basename $line)
	name="${basename%.*}"
	repo_dir="./repos/${name}"
	metrics_dir="./metrics/$name/"

	[ ! -d "$repo_dir" ] && git clone "$line" "$repo_dir" --depth 1
    cloc "$repo_dir" | awk '/SUM/{ printf "%s\n%s\n", $4, $5}' > "$metrics_dir/cloc.txt"
	rm -rf "$repo_dir"
done < "$repo_list"
