cdash
=====
A minimalist cryptocurrency portfolio dashboard for the command-line that draws
market data from the [CoinMarketCap][cmc] API.

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

### Specifying the base currency ###
The [base currency][base] may be specified using the `--base` flag. The
following currency codes are supported:

Code  | Currency
------|----------------------
`AUD` | Australia Dollar         
`BRL` | Brazil Real              
`CAD` | Canada Dollar            
`CHF` | Switzerland Franc        
`CLP` | Chile Peso               
`CNY` | China Yuan Renminbi      
`CZK` | Czech Republic Koruna    
`DKK` | Denmark Krone            
`EUR` | Euro Member Countries    
`GBP` | United Kingdom Pound     
`HKD` | Hong Kong Dollar         
`HUF` | Hungary Forint           
`IDR` | Indonesia Rupiah         
`ILS` | Israel Shekel            
`INR` | India Rupee              
`JPY` | Japan Yen                
`KRW` | Korea (South) Won        
`MXN` | Mexico Peso              
`MYR` | Malaysia Ringgit         
`NOK` | Norway Krone             
`NZD` | New Zealand Dollar       
`PHP` | Philippines Peso         
`PKR` | Pakistan Rupee           
`PLN` | Poland Zloty             
`RUB` | Russia Ruble             
`SEK` | Sweden Krona             
`SGD` | Singapore Dollar         
`THB` | Thailand Baht            
`TRY` | Turkey Lira              
`TWD` | Taiwan New Dollar        
`USD` | United States Dollar     
`ZAR` | South Africa Rand        

The `--base` option is case-insensitive, and defaults to `USD`.  


Windows Usage
-------------
`cdash` uses ANSI escape sequences for colorization. Windows users must take
note to use a compatible shell.


[Releases]: https://github.com/chrisallenlane/cdash/releases/
[base]:     https://en.wikipedia.org/wiki/Currency_pair#Base_currency
[cmc]:      https://coinmarketcap.com/
[img]:      ./.github/screen.jpg
