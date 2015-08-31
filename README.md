# godepex
Exclude packages from [godep](https://github.com/tools/godep "godep github page")

## Install

    $ go get github.com/mmcloughlin/godepex

## Usage

In your go project, after you've run `godep save` you can remove dependencies matching a prefix with 

    $ godepex example.com/pkg

This will both

* Filter out the entries from your `Godeps/Godeps.json`
* Delete corresponding directories from `Godeps/_workspace/src`
