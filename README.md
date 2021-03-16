# log

[![Go Reference](https://pkg.go.dev/badge/github.com/tprasadtp/log.svg)](https://pkg.go.dev/github.com/tprasadtp/log)
[![test](https://github.com/tprasadtp/log/workflows/test/badge.svg)](https://github.com/tprasadtp/log/actions?query=workflow%3Atest)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/tprasadtp/log?color=00ADD8&label=go&logo=go)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/tprasadtp/log?label=version&logo=github&sort=semver)](https://github.com/tprasadtp/log/tags)
[![License](https://img.shields.io/badge/license-MIT-orange?labelColor=313131)](https://github.com/tprasadtp/log/blob/master/LICENSE)
![Analytics](https://ga-beacon.prasadt.com/UA-101760811-3/github/log)


## Notes

This repository for log package used by other tools including some internal apps.
This package should be considered alpha.

## Why

Logging libraries are tied too much into the application.
In some cases where data is in hotpath and logging becomes your bottleneck it might
be worth to integrate it with your choice of logging library, however in 98% of other cases
its benificial. This approach has a defined interface that all `adapters` implement. Thus you are free to change logging library as log as your adapters implements the said interface.
