[![scrape](https://snapcraft.io/scrape/badge.svg)](https://snapcraft.io/scrape)

# Scrape
CLI utility to scrape emails from websites

### Usage
Sample call:

`scrape -w https://lawzava.com` 

#### Parameters:
          --async             Scrape website pages asynchronously (default true)
      -d, --depth int         Max depth to follow when scraping recursively (default 3)
          --emails            Scrape emails (default true)
          --follow-external   Follow external 3rd party links within website
      -h, --help              help for scrape
          --logs              Print debug logs
          --recursively       Scrape website recursively (default true)
      -w, --website string    Website to scrape (default "https://lawzava.com")
```
