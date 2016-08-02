godiva [![GoDoc](http://godoc.org/github.com/mappymappy/godiva/lib/godiva?status.svg)](http://godoc.org/github.com/mappymappy/godiva/lib/godiva)
======

A simple DependencyInjectionContainer for Golang.

## Installation:

```
go get github.com/mappymappy/godiva/lib/godiva
```

## Usages:

### register

```
container := CreateGodivaContainer()
container.Register("chocolate",func (c *Container) (interface{}, error) { return Chocolate{})
```

### create

```
chocolate := container.Create("chocolate",false)
girl.eat(chocolate)
```

## License

```
Copyright (c) 2016 marnie_ms4
Released under the MIT license
http://opensource.org/licenses/mit-license.php
```

## About Me
[marnie_ms4](https://github.com/mappymappy?tab=repositories)
