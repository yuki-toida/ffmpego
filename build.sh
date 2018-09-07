#!/bin/sh

rm -rf dist

go build -o ./dist/ffmpego

mkdir dist/input
mkdir dist/output
cp README.md dist/README.md
