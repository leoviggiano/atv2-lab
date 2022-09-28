#!/bin/bash

[ ! -d "./repos" ] && mkdir repos
[ ! -d "./metrics" ] && mkdir metrics

repo_list=$1
fails_dir="./fails"
echo $repo_list
if [ ! -f "$repo_list" ]
then
	echo "File $repo_list does not exists"
	exit 1
fi

function save_error() {
	repository=$1
	metrics_dir=$2
	[ ! -d "$fails_dir" ] && mkdir "$fails_dir"
	[ ! -f "$fails_dir/metrics.txt" ] && touch "$fails_dir/metrics.txt"
	[ ! -f "$fails_dir/repositories.txt" ] && touch "$fails_dir/repositories.txt"

	echo "$metrics_dir" >> "$fails_dir/metrics.txt"
	echo "$repository" >> "$fails_dir/repositories.txt"
}

while IFS= read -r line
do
	basename=$(basename $line)
	name="${basename%.*}"
	repo_dir="./repos/${name}"
	metrics_dir="./metrics/$name/"
	if [ -d "$metrics_dir" ] 
	then
		echo "$metrics_dir" >> "teste.txt"
		continue
	fi

	[ ! -d "$repo_dir" ] && git clone "$line" "$repo_dir"
	mkdir $metrics_dir
	cloc "$repo_dir" | awk '/SUM/ { printf "%s\n%s\n", $4, $5}' > "$metrics_dir/cloc.txt"
	java -jar ./ck/target/ck-0.7.1-SNAPSHOT-jar-with-dependencies.jar "$repo_dir" true 0 true "$metrics_dir" || save_error $line $metrics_dir
	rm -rf "$repo_dir"
done < "$repo_list"
