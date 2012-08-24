#!/usr/local/bin/perl

use Message;
use ConfigFile;

$obj = new Message(msg => "Hello, World.");
$obj->show();

$obj = new Message or die "Improperly constructed 'Message',";
$obj->show();

warn("foo");

$cfg_file = load ConfigFile(".foobar");

