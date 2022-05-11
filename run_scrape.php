#!/usr/bin/env php

<?php
  // TODO: Make this file executable or the above won't work

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

  $command = "./bin/scrape.darwin-amd-64 -w";

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
      exec("{$command} {$url}", $output, $result);

      // No problem! print the output
      foreach($output as $email) {
        print("{$url}\t$email\n");
      }
    }

  }

?>
