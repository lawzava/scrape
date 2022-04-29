#!/usr/bin/env php
<?php
// Don't forget to make this file executable or the above won't work

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
