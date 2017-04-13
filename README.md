# King Julien

Nohup. But with a little extra.

[![Build Status](https://travis-ci.org/kcmerrill/kj.svg?branch=master)](https://travis-ci.org/kcmerrill/kj) [![Join the chat at https://gitter.im/kcmerrill/kj](https://badges.gitter.im/kcmerrill/kj.svg)](https://gitter.im/kcmerrill/kj?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

![kj](assets/king-julien.jpg "kj")

## Binaries || Installation

[![MacOSX](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/apple_logo.png "Mac OSX")](http://go-dist.kcmerrill.com/kcmerrill/kj/mac/amd64) [![Linux](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/linux_logo.png "Linux")](http://go-dist.kcmerrill.com/kcmerrill/kj/linux/amd64)

via go:

`$ go get -u github.com/kcmerrill/kj`

## About

`kj` is a simple process manager that keeps a single(or multiple) workers working. The idea being, you can keep a process running while you're away at lunch. If you need something more sophisticated I'd recommend using `supervisord`. `kj` is small and lightweight and should be treated as such.

* keep a process running
* keep X number of processes running
* automatically shove processes to the background
* keep logs of stdin/stdout
* log reaping based on log size(coming soon)

## Demo

[![asciicast](https://asciinema.org/a/113063.png)](https://asciinema.org/a/113063)

## Usage

The most basic usage is simply to run `kj echo hello world` where `echo hello world` is the command you want to run. It will spawn the process off in the background, and make sure it can't be interrupted by the hangup signal.

`kj --cmd="your command here" --id="name.you.want.to.use" --workers=10 --run-once`

The above example will store your logs in a file called `name.you.want.to.use-_worker-id_.log`. By default, `kj` will attempt to keep your process running by restarting it. `--run-once` flag will only let the process run once and then kj will exit out.
