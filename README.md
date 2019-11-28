[![scrape](https://snapcraft.io/scrape/badge.svg)](https://snapcraft.io/scrape)

# Scrape
CLI utility to scrape emails from websites

### Usage
Sample call:

`scrape -w https://lawzava.com` 

Depends on `chromium` or `google-chrome` being available in path if `--js-wait` is used

#### Parameters:
```
          --async             Scrape website pages asynchronously (default true)
      -d, --depth int         Max depth to follow when scraping recursively (default 3)
          --emails            Scrape emails (default true)
          --follow-external   Follow external 3rd party links within website
      -h, --help              help for scrape
          --js-wait           Should wait for JS to execute
          --logs              Print debug logs
          --recursively       Scrape website recursively (default true)
      -w, --website string    Website to scrape (default "https://lawzava.com")
```
