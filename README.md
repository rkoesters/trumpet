Trumpet 🎺
=======

A bot that generates tweets based on the accounts it follows.

[![GoDoc](https://godoc.org/github.com/rkoesters/trumpet?status.svg)](https://godoc.org/github.com/rkoesters/trumpet)
[![Build Status](https://travis-ci.org/rkoesters/trumpet.svg?branch=master)](https://travis-ci.org/rkoesters/trumpet)
[![Go Report Card](https://goreportcard.com/badge/github.com/rkoesters/trumpet)](https://goreportcard.com/report/github.com/rkoesters/trumpet)

Building
--------

First, you will need to make sure you have the dependencies:

	$ make deps

Next, you are able to build the program:

	$ make

Installing
----------

To install to `$GOPATH/bin`:

	$ make install

To install to a custom location, for example `/usr/local/bin`:

	$ sudo make install prefix=/usr/local

Configuration
-------------

Trumpet needs to be configured to connect to Twitter. You will need a
consumer key, a consumer secret, an access token, and an access secret.
See
<https://developer.twitter.com/en/docs/basics/authentication/guides/access-tokens.html>
for more information.

Once you have the required information, you can use the
`trumpet-mkconfig` tool to create the configuration file.

	$ trumpet-mkconfig

For more information on how to use the tool, run:

	$ trumpet-mkconfig -help

Running
-------

Once you have created the configuration file, you are ready to run the
bot:

	$ trumpet

For more information on the program's usage, run:

	$ trumpet -help

License
-------

See [LICENSE](LICENSE).
