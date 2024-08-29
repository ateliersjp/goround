#!/bin/sh

if [ "$1" != 'goround' ]; then
    set -- 'goround' "$@"
fi

exec "$@"
