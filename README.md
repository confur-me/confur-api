confur-api
============
[![Buid status][travis-badge]][travis-url]

[![Build Status](https://travis-ci.org/confur-me/confur-api.svg?branch=master)](https://travis-ci.org/confur-me/confur-api)

### Quickstart

```
cd $GOPATH/src
git clone https://github.com/confur-me/confur-api.git
cd confur-api
make
make install
confur -c=/path/to/config/file.yml db:migrate
confur -c=/path/to/config/file.yml start
```

### Development

```
cd $GOPATH/src
git clone https://github.com/confur-me/confur-api.git
cd confur-api
go run confur.go db:migrate
go run confur.go start
```

[travis-badge]: https://travis-ci.org/confur-me/confur-api.svg?branch=master
[travis-url]: https://travis-ci.org/confur-me/confur-api
