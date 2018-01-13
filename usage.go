package main

func usage() string {

	return `cdash

A cryptocurrency portfolio dashboard that draws market data from the
CoinMarketCap API.

Usage:
  cdash [options]

Options:
  -h --help           Show this help.
  -v --version        Display the version number.
  -c --config=<file>  Path to config file.
  -b --base=<cur>     Base currency code [default: USD].
  -n --name           Display token name in addition to symbol.`
}
