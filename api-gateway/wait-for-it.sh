#!/bin/sh
# wait-for-it.sh: wait for a service to be available before executing a command

set -e

host="$1"
shift
cmd="$@"

until nc -z "$host"; do
  >&2 echo "Waiting for $host to be available..."
  sleep 1
done

>&2 echo "$host is available, executing command: $cmd"
exec $cmd
