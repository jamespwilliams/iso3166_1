# iso3166\_1

[![Documentation](https://godoc.org/github.com/jamespwilliams/iso3166_1?status.svg)](https://godoc.org/github.com/jamespwilliams/iso3166_1)

Go library for efficiently accessing the country information specified in
ISO3166-1.

#### Performance

Performance when looking up country information for a given Alpha-2 or Alpha-3
code:

| Module                                                  | Result (Alpha-2)      | Result (Alpha-3) |
|---------------------------------------------------------|-----------------------|------------------|
| iso3166\_1                                              |  19.1 ns/op           |  19.3 ns/op      |
| github.com/launchdarkly/go-country-codes (unmaintained) |  48.9 ns/op           |  48.6 ns/op      |
| github.com/TheBookPeople/iso3166                        |  63.2 ns/op           |  64.8 ns/op      |
| github.com/pariz/gountries                              | 105 ns/op             | 103 ns/op        |
| github.com/biter777/countries                           | 198 ns/op             | 251 ns/op        |

Source:

```
cd _comparison
go test -bench .
```

#### Development

The `generated.go` file is generated by `internal/gen/main.go`. It can be
(re)generated by running

```
go generate .
```

from the root of the repo. Run tests with:

```
go test .
```

#### Contributing

Pull requests are welcome.
