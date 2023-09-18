# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).


## 2.0.0 - 2023-03-10
Initial release.
### change
- adjust dir struct
- maputil/slice:change to  github.com/duke-git/lancet extension
### Add
- test/mock: add condMatcher
- test/assert: support testify  chain call
- idgenerator: support localIdGenerator
### Del
- di: remove

## 1.0.0 - 2023-03-10
Initial release.
### Added
- support generic
- bitmaputil: base on bitmap, support exist and get all
- cputil:  copier extension
- expr: expr, eg: ADD,MIN,MAX,COND
- fileutil: base on afero for file extension
- maputil: support eg:Map
- sliceutil: support generic,eg:Contain,Sum,Map,Reduce
- di: uber/dig extension