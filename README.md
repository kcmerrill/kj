# King Julien

Nohup. But with a little extra.

[![Build Status](https://travis-ci.org/kcmerrill/kj.svg?branch=master)](https://travis-ci.org/kcmerrill/kj) [![Join the chat at https://gitter.im/kcmerrill/kj](https://badges.gitter.im/kcmerrill/kj.svg)](https://gitter.im/kcmerrill/kj?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

![kj](assets/king-julien.jpg "kj")

## Binaries || Installation

[![MacOSX](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/apple_logo.png "Mac OSX")](http://go-dist.kcmerrill.com/kcmerrill/kj/mac/amd64) [![Linux](https://raw.githubusercontent.com/kcmerrill/go-dist/master/assets/linux_logo.png "Linux")](http://go-dist.kcmerrill.com/kcmerrill/kj/linux/amd64)

via go:

`$ go get -u github.com/kcmerrill/kj`
## About

I wanted something that would run in the background(use the `&`) sign, something that logged my application, something that would keep it alive in case it died for whatever reason, and also something that would keep X number of tasks running. 
## Usage

The most basic usage is simply to run `kj your command here &`. Output will be stored to `kj-1.log`. 

More complex usage is to run kj with params.

`kj --cmd="your command here" --id="name.you.want.to.use" --workers=10 --keep-alive`
