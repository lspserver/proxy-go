#!/bin/bash

tag="$1"

docker build -f Dockerfile -t ghcr.io/lspserver/proxy:"$tag" .
