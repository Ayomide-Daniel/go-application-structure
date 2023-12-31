#!/bin/sh

set -e
COMMAND=$@

echo "Waiting for database to start..."
maxTries=10

while [ "$maxTries" -gt 0 ] && ! mysql -h "$MYSQL_HOST" -P "$MYSQL_PORT" -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" "$MYSQL_DB" -e 'SHOW TABLES'; do
  maxTries=$((maxTries-1))
  sleep 3
done

echo 
if [ "$maxTries" -le 0 ]; then
  echo >&2 "error: unable to contact database after 30 seconds."
  exit 1
fi

exec $COMMAND