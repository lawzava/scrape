#!/usr/bin/env php

<?php
  // TODO: Make this file executable or the above won't work

  include "ScrapeCommand.php";

  $command = new ScrapeCommand($argc, $argv);
  $command->initialize();
  $command->process();
