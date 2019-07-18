# properties
Command-line utility to work with '\*.properties' file

# Download

Select released version that you are interested in on https://github.com/OpusCapita/properties/releases page and download archive with binary file for your platform.

# Use binary without installation

Run the latest properties binary
```sh
curl -sL https://git.io/oc-properties| bash -s -- get --file ./test.properties --key a.b.c
```

Or, if you want to use an specific version

```sh
curl -sL https://git.io/oc-properties | VERSION=v1.0.0-rc5 bash -s -- get --file ./test.properties --key a.b.c
```

In both cases you should get property value printed into stdout by key 'a.b.c' from file './test.properties'

# Install binary

In fact the same as running script but without passing any arguments to bash

```sh
curl -sL https://git.io/oc-properties | TARGET_DIR=/one/two/tree bash
```

You'll get 'properties' binary file installed into /one/two/tree folder (corresponding message will be printed to stdout)
**P.S.** If you don't define TARGET_DIR then installation will be performed into random temporary folder (which is not what you want)
