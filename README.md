# GMO Farming

[![Tested with A20.5 b2](https://img.shields.io/badge/A20.5%20b2-tested-blue.svg)](https://7daystodie.com/) [![Automated Release](https://github.com/jonathan-robertson/gmo-farming/actions/workflows/main.yml/badge.svg)](https://github.com/jonathan-robertson/gmo-farming/actions/workflows/main.yml)

7 Days to Die Modlet: Genetically modify seeds to grow plants with new properties

## Features

TODO

## Development

This modlet is truly massive due to all the plant traits and supported combinations of them.

Becuase of this, we relied on Go to generate [blocks.xml](./Config/blocks.xml), [recipes.xml](./Config/recipes.xml), and [Localization.txt](./Config/Localization.txt).

To make adjustments to generated files:

1. edit the `*.go` files in the [data](./data) and [gen](./gen) packages to suit your preferences.
2. run `go run main.go` to have it dump files generated files into the [Config](./Config) folder.

## Learn Go

If you don't already know Go, you can learn it [over here](https://go.dev/learn/).

I highly recommend the [Tour of Go](https://go.dev/tour/), which can usually get someone up to speed in about 30mins to 1hr.
