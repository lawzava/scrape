[![scrape](https://snapcraft.io/scrape/badge.svg)](https://snapcraft.io/scrape)

# Scrape
CLI utility to scrape emails from websites

### Features

- Asynchronous scraping
- Recursive link follow
- External link follow
- Cloudflare email obfuscation decoding
- Client side rendered pages support through headless `chromium` load awaits
- Simple, grepable output

### Usage
Sample call:

`scrape -w https://lawzava.com` 

Depends on `chromium` or `google-chrome` being available in path if `--js` is used

#### Parameters:
```
          --async             Scrape website pages asynchronously (default true)
      -d, --depth int         Max depth to follow when scraping recursively (default 3)
          --emails            Scrape emails (default true)
          --follow-external   Follow external 3rd party links within website
      -h, --help              help for scrape
          --js                Enables JS execution await
          --debug             Print debug logs
          --recursively       Scrape website recursively (default true)
      -w, --website string    Website to scrape (default "https://lawzava.com")
```
