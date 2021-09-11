[![scrape](https://snapcraft.io/scrape/badge.svg)](https://snapcraft.io/scrape)

# Scrape
CLI utility to scrape emails from websites

## Features

- Asynchronous scraping
- Recursive link follow
- External link follow
- Cloudflare email obfuscation decoding
- Client side rendered pages support through headless `chromium` load awaits
- Simple, grepable output

## Install

* MacOS:

    ```bash
    brew tap lawzava/scrape https://github.com/lawzava/scrape
    brew install scrape
    ```

* Linux:

    ```bash
    sudo snap install scrape
    ```


## Usage
Sample call:

`scrape -w https://lawzava.com` 

Depends on `chromium` or `google-chrome` being available in path if `--js` is used

### Parameters:
```
      --async             Scrape website pages asynchronously (default true)
      --debug             Print debug logs
  -d, --depth int         Max depth to follow when scraping recursively (default 3)
      --follow-external   Follow external 3rd party links within website
  -h, --help              help for scrape
      --js                Enables EnableJavascript execution await
      --output string     Output type to use (default 'plain', supported: 'csv', 'json') (default "plain")
      --output-with-url   Adds URL to output with each email
      --recursively       Scrape website recursively (default true)
      --timeout int       If > 0, specify a timeout (seconds) for js execution await
  -w, --website string    Website to scrape (default "https://lawzava.com")
```

## Note about scraper package

For those that are looking for `scraper` package - this repository was intended as a cli-use only thus the scraper package was moved to [lawzava/emailscraper](https://github.com/lawzava/emailscraper).
The `scrape` utility will be maintained as a CLI implementation of `emailscraper` package.
