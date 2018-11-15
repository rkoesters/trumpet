Trumpet
=======

A bot that generates tweets based on the accounts it follows.

[![Build Status](https://travis-ci.org/rkoesters/trumpet.svg?branch=master)](https://travis-ci.org/rkoesters/trumpet)
[![Go Report Card](https://goreportcard.com/badge/github.com/rkoesters/trumpet)](https://goreportcard.com/report/github.com/rkoesters/trumpet)

Building
--------

	$ make deps
	$ make

Configuration
-------------

Trumpet needs to be configured to connect to Twitter. You will need a
consumer key, a consumer secret, an access token, and an access secret.
See
<https://developer.twitter.com/en/docs/basics/authentication/guides/access-tokens.html>
for more information.

	$ make config
