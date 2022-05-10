rm -rf Config-Vanilla/*
rm -rf Config-CrystalHell/*
rm -rf Config
mkdir Config
go run .
cp -r Config-Shared/* Config
cp -r Config-Vanilla/* Config

