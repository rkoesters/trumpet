#!/bin/sh
set -e

if [ $# -eq 0 ]; then
	output_file="trumpet.conf"
elif [ $# -eq 1 ]; then
	output_file="$1"
else
	echo "Usage: $0 [OUTPUT_FILE]" >&2
	exit 1
fi

printf 'Consumer Key: '
read -r consumer_key

printf 'Consumer Secret: '
read -r consumer_secret

printf 'Access Token: '
read -r access_token

printf 'Access Secret: '
read -r access_secret

{
	echo "consumer-key=$consumer_key"
	echo "consumer-secret=$consumer_secret"
	echo "access-token=$access_token"
	echo "access-secret=$access_secret"
} >"$output_file"

echo "Configuration written to '$output_file'."
