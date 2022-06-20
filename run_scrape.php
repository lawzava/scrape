#!/usr/bin/env php

<?php
  // TODO: Make this file executable or the above won't work

/**
 * INSTRUCTIONS:
 *
 * --urls <default to run_scrape_urls.csv>
 * --emails <default to run_scrape_emails.csv
 * --output <default to run_scrape_output.csv
 * --url-filters <default to run_scrape_url_filters.csv
 * --email-domain-filters <default to run_scrape_email_domain_filters.csv
 * --url-column <defaults to column 1>
 *
 * --help
 * --verbose
 * --scrape-only
 * --email-filter-only
 * --no-url-filter
 * --no-email-filter
 * --no-dup-filter
 *
 *
 * Given a CSV file of urls extract each url and run the scrape command to get emails
 * Output the emails to a file
 *
 * Optionally apply a list of url domains to NOT run the scrape
 * Optionally apply list of email domains to filter out of output csv
 * Optionally choose which column to grab the urls in the csv file
 *
 *
 *
 */

  // INSTRUCTIONS:
  // --in_urls


  //
  // TODO:
  //   1) Add command line args
  //     1. help
  //     b. CSV URL file name as input
  //     c. CSV mail file name for output (if no default)
  //     d. column to use in input csv
  //     e. create default filter file for domains to skip scrape
  //     f. create default filter file for email domains to not put in output file
  //     g. additional or replacement scrape filter file to use (different from default)
  //     h. additional or replacement email domain filter file to use (different from default)
  //   2) Add optional post processing step to filter domains from input file
  //   3) Add post processing step to de-dup email address by full domain

  // Flow
  // Read command line
  // If not scraping, skip to
  // Open scrape URL filter file
  // Loop over scrape URLs, skipping ones filtered
  // Open output file and stuff output data into file
  // Quit if just scraping
  //
  // If processing from command line
  // Open email domain filter file and build filter array
  // Loop over emails and filter out domains

  include "ScrapeCommand.php";

  $command = new ScrapeCommand($argc, $argv);
  $command->initialize();
  $command->process();

  die;



  $scrapeCommand = "./bin/scrape.darwin-amd-64 -w";


  // Validate the input
  //   Usage is "run_scrape <filename>"
  //   Therefore there should be one and only 2 args
  if ($argc !== 2) {
    print("Invalid command line. Use \"run_scrape.php <filename>\"\n");
    die;
  }

  // Get the input file and validate it can be opened
  $filename = $argv[1];
  $file = fopen($filename, "r");

  if (!$file) {
    print("Could not open file {$filename}\n");
    die;
  }

  // Print out that the run is starting
  $start = date('n/d/y h:i:s');

  // Loop over each line of the input file
  while(!feof($file)) {
    $output = [];

    $line = fgets($file);
    $columns = explode("\t", $line);

    // Start Processing
    $url = $columns[0];

    // Validate the URL
    $isValid = filter_var($url, FILTER_VALIDATE_URL);
    if (!$isValid) {
      print("{$url}\tINVALID URL\n");
    } else {
      // Execute the scraper with this URL
      exec("{$scrapeCommand} {$url}", $output, $result);

      // No problem! print the output
      foreach($output as $email) {
        print("{$url}\t$email\n");
      }
    }

  }

  function parseCommandLine() {

  }

  function showHelp() {

  }

?>
