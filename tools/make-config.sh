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
read -r consumerkey

printf 'Consumer Secret: '
read -r consumersecret

printf 'Access Token: '
read -r accesstoken

printf 'Access Secret: '
read -r accesssecret

{
	echo "consumer-key=$consumerkey"
	echo "consumer-secret=$consumersecret"
	echo "access-token=$accesstoken"
	echo "access-secret=$accesssecret"
} >"$output_file"

echo "Configuration written to '$output_file'."
