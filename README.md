# forecast
A Forecast ([forecastapp.com](https://www.forecastapp.com)) API Client For Go.

## Installation
Add the dependency to your project:

```bash
$ dep ensure github.com/joefitzgerald/forecast
```

# Usage
First, construct an API:

```golang
c := &forecast.Config{
  AccountID: "000000",
  Scheme:    "https",
  Host:      "api.forecastapp.com",
  Username:  "jsmith@example.com",
  Password:  "password-here",
}
api, err := forecast.New(c)
```

Then, make use of the API. Consult [godoc](http://godoc.org/github.com/joefitzgerald/forecast) for detailed API documentation.

## License

[Apache 2.0](https://github.com/joefitzgerald/forecast/blob/master/LICENSE)
