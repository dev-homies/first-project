#!/bin/sh

if [ "${1}" == "" ]; then
  docker compose logs --follow --timestamps
else
  docker compose logs "${1}" --follow --timestamps
fi
