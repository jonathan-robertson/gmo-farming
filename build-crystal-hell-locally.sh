#!/bin/sh
rm -rf Config
mkdir Config
go run .
cp -r Config-Shared/* Config
cp -r Config-CrystalHell/* Config

