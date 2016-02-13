BMP CLI [![Build Status](https://travis-ci.org/maximilien/bosh-bmp-cli.svg?branch=master)](https://travis-ci.org/maximilien/bosh-bmp-cli#)
==============================

[![Join the chat at https://gitter.im/maximilien/bosh-bmp-cli](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/maximilien/bosh-bmp-cli?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

This is a simple CLI to help manage (provision, list, and delete) a bare metal server pool that can be used by the bosh-softlayer-cpi when doing deployment that include bare metal servers. This helps speed up these types of deployments.

This tool is licensed under the [Apache 2.0 OSS license](https://github.com/maximilien/i18n4go/blob/master/LICENSE). We'd love to hear from you if you are using, attempting to use, or planning to use this tool. 

Two additional ways, besides Gitter or Slack chat above, to contact us:

1. Feel free to [open an issue (good or bad)](https://github.com/maximilien/bosh-bms-cli/issues) here on Github.
2. Send email to `bosh-bms-cli` at the Gmail domain.

## Getting Started
------------------

### Overview Presentations, Talks, Blogs
----------------------------------------

* TBD

### Getting Latest Executable: bmp
----------------------------------

Assuming you have a valid [Golang 1.4.2](https://golang.org/dl/) or [later](https://golang.org/dl/) installed for your system, you can quickly get the latest `i18n4go` executable by running the following `go` command:

```
$ go get github.com/maximilien/bosh-bmp-cli/bmp
```

This will build and place the `bmp` executable built for your operating system in your `$GOPATH/bin` directory.

### Cloning and Building
------------------------

Clone this repo and build it. Using the following commands on a Linux or Mac OS X system:

```
$ mkdir -p bosh-bms-cli/src/github.com/maximilien
$ export GOPATH=$(pwd)/bosh-bms-cli:$GOPATH
$ cd bosh-bms-cli/src/github.com/maximilien
$ git clone https://github.com/maximilien/bosh-bms-cli.git
$ cd bosh-bms-cli
$ ./bin/build
```

NOTE: if you get any dependency errors, then use `go get path/to/dependency` to get it, e.g., `go get github.com/onsi/ginkgo` and `go get github.com/onsi/gomega`

The executable output should now be located in: `out/bmp`. Place it wherever you want, e.g., `/usr/local/bin` on Linux or Mac OS X.

You can now use the `bmp` executable to issue some commands.

### Running Tests
-----------------

You should run the tests to make sure all is well, do this with: `$ ./bin/test` in your cloned repository.

The output should be similar to:

```
$ bin/test

 Cleaning build artifacts...

 Formatting packages...

...

 Vetting packages for potential issues...

SWEET SUITE SUCCESS
```

### Typical Workflow
--------------------

The recommended workflow is to use the commands (documented below) in the following order. For each command, use the command's help or this README for details and to experiment for your project. So the steps are:

1. TBD

2. ship and profit :)

### Typical Workflow Diagram

![Typical bosh-bms-cli workflow diagram](https://github.com/maximilien/bosh-bms-cli/blob/master/docs/images/typical-workflow.png)

### Help
--------

TBD
## How 

## Commands

### `bmp`

desc 'show bmp bosh sub-commands'
say('bosh bmp sub-commands:')

### `bmp target`

desc 'Set the URL of Bluemix Provision Server'

### `bmp login`

desc 'Login Bluemix Provision Server'

### `bmp status`

desc 'show bmp status'

### `bmp bms`

desc 'List all bare metals'

### `bmp stemcells`

desc 'List all stemcells'

### `bmp update state`

desc 'Update the server state ("bm.state.new", "bm.state.using", "bm.state.loading", "bm.state.failed", "bm.state.deleted")'

### `bmp tasks`

desc 'List tasks'
option "--latest number", "The number of latest tasks, default is 50"

### `bmp task`

desc 'Show the output of the task'
option "--debug", "Get the debug info of the task"

### `bmp create baremetals`

desc 'Create the missed baremetals'
option "--dryrun", "only verify the orders"

### `bmp sl packages`

desc 'List all Softlayer packages'

### `bmp sl package-options`

desc 'List all options of Softlayer package'
