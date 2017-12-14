cdash ("coin dash")
===================
A minimalist cryptocurrency portfolio dashboard for the command-line that draws
market data from the [Coin Market Cap][cmc] API.

![screenshot][img]


Installing
----------
1. Visit the [Releases][] page
2. Download the appropriate executable
3. Optionally rename the downloaded executable
4. Optionally place the executable on your system `PATH`


Configuring
-----------
Specify your holdings via a `yml` file structured as follows:

```yml
---
- symbol   : BTC
  holdings : 100

- symbol   : LTC
  holdings : 100

- symbol   : ETH
  holdings : 100
```

By default, `cdash` will attempt to open this file at `~/.config/cdash.yml`. If
you choose to store it elsewhere, provide `cdash` the approriate path using the
`--config` option:

```sh
cdash --config=/path/to/cdash.yml
```


[Releases]: https://github.com/chrisallenlane/cdash/releases/
[cmc]:      https://coinmarketcap.com/
[img]:      ./.github/screen.jpg