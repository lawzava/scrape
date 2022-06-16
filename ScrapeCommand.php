<?php

class ScrapeCommand {
  public $argc;
  public $argv;
  public $commandLineOptions;
  public $scrapeOutputFile;

  const SCRAPE_COMMAND = "./bin/scrape.darwin-amd-64 -w";

  public function __construct($argc, $argv) {
    $this->argc = $argc;
    $this->argv = $argv;
  }

  public function initialize() {
    $this->parseCommandLine();
  }

  public function process() {
    $this->logLine('Processing starting');
    $url = 'foo';

    // $this->executeScrape($url);
  }

  protected function executeScrape($url) {
    $scrapeCommand = self::SCRAPE_COMMAND;

    // Execute the scraper with this URL
    exec("{$scrapeCommand} {$url}", $this->scrapeOutputFile, $resultCode);
  }

  protected function parseCommandLine() {
    $arg = 2;

    while ($arg >= $this->argc) {
      $nextArg = $this->argv[$arg];

      $arg++;
    }
  }

  protected function displayHelp() {
  }

  protected function logLine($text) {
    print("{$text}\n");
  }

}
