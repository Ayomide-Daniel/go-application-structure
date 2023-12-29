#!/bin/sh

set -e
COMMAND=$@

echo "Waiting for database to start..."
maxTries=10

while [ "$maxTries" -gt 0 ] && ! mysql -h "$MYSQl_HOST" -P "$MYSQl_PORT" -u "$MYSQl_USER" -p "$MYSQl_PASSWORD" -e "show databases;" > /dev/null 2>&1; do
  maxTries=$((maxTries-1))
  sleep 3
done

echo 
if [ "$maxTries" -le 0 ]; then
  echo >&2 "error: unable to contact database after 30 seconds."
  exit 1
fi

exec $COMMAND