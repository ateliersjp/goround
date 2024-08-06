#!/bin/sh

mkdir -p /var/run/goround

if [ "$1" != 'goround' ]; then
    set -- 'goround' "$@"
fi

exec "$@"
