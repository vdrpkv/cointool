# cointool
Simple command line tool for working with cryptocurrency coins.

It uses `coinmarketcap` service to retrieve coin information under the hood.

If not specified, `cointool` send requests to `sandbox` environment of `coinmarketcap`.

Examples below features `sandbox` environment of `coinmarketcap`. **That's why example values doesn't look real**.

Set `--api-key` and `--api-prefix` flags to `pro` environment in order to get real values.

See [Flags](https://github.com/vdrpkv/cointool/new/main?readme=1#flags) section below for more information.


# Usecases

## Convert coins
```
cointool convert AMOUNT FROM TO
```

### Example

Convert `10000` `USD` to `BTC`
```
cointool convert 10000 USD BTC
38227.273628733994
```


## Get exchange rate
```
cointool rate FROM TO
```

### Example

Get exchange rate from `USD` to `BTC`
```
cointool rate USD BTC
0.7368213384961149
```

### Note
`cointool rate` command expects cryptocurrency symbol as first argument and any currency symbol as the second one.

`cointool rate` on on `pro` environment responds with error on requests where first symbol is fiat currency:

```
cointool rate --api-prefix=pro --api-key=*** USD BTC
ERROR: get exchange rate: exchange rate not found
```

## Check is currency fiat
```
cointool fiat SYMBOL
```

### Example
```
cointool fiat USD
true
```
```
cointool fiat BTC
false
```

# Flags

## CoinMarketCap API Key
Set `coinmarketcap` API key. Sandbox key is shown below and used by default.
```
--api-key=b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c
```

## CoinMarketCap API Prefix
Set `coinmarketcap` API prefix. `pro` prefix is shown below and requires valid `coinmarketcap` API key to access.
```
--api-key=pro
```

## Timeout
Set maximum time to wait for command to perform. The default timeout is `7` seconds.
```
--timeout=7s
```