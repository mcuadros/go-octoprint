go-octoprint [![Build Status](https://travis-ci.org/mcuadros/go-octoprint.svg?branch=master)](https://travis-ci.org/mcuadros/go-octoprint) [![GoDoc](http://godoc.org/github.com/mcuadros/go-octoprint?status.svg)](http://godoc.org/github.com/mcuadros/go-octoprint)
==============================

Go library for accessing the [OctoPrint](http://octoprint.org/)'s [REST API](http://docs.octoprint.org/en/master/api/index.html).

Installation
------------

The recommended way to install go-octoprint

```
go get github.com/mcuadros/go-octoprint
```

Example
-------

### Retrieving the current connection state:

```go
client, _ := NewClient("<octoprint-url>", "<api-key>")

r := octoprint.ConnectionRequest{}
s, err := r.Do(client)
if err != nil {
  log.Error("error requesting connection state: %s", err)
}

fmt.Printf("Connection State: %q\n", s.Current.State)
```


### Retrieving current temperature for bed and extruders:

```go
r := octoprint.StateRequest{}
s, err := r.Do(c)
if err != nil {
	log.Error("error requesting state: %s", err)
}

fmt.Println("Current Temperatures:")
for tool, state := range s.Temperature.Current {
	fmt.Printf("- %s: %.1f°C / %.1f°C\n", tool, state.Actual, state.Target)
}
```

License
-------

MIT, see [LICENSE](LICENSE)
