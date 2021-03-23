#!/usr/bin/env php
<?php
//$options = getopt("f:hp:");
/**
 * File loops through atheleteTrax sites and parses for a url and email
 */

$url = 'https://app.mysportsort.com/main/view/index?an=';

for($i=1;$i<1000;$i++){
  //$i=3;
  $email_match =  $url_match = '';
  $site_url = $url."$i";

  if(!$html = file_get_contents($site_url)){
    continue;
  }

  $url_regex = '/Website:.*(https?:\/\/\w*\.\w*\.?(\w*)?)/';
  if(preg_match($url_regex,$html,$url_matches)){
    $url_match = $url_matches[1];
  }

  $email_regex = '/Email:.*>([a-zA-Z0-9]*@\w*\.\w*)/';
  if(preg_match($email_regex,$html,$email_matches)){
    $email_match = $email_matches[1];
  }


  //print_r($url_matches);
  echo  $site_url."\t".$url_match."\t".$email_match."\n";

  //sleep(.5);

 //exit();

}