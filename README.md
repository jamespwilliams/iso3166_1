# iso3166\_1

[![Documentation](https://godoc.org/github.com/jamespwilliams/iso3166_1?status.svg)](https://godoc.org/github.com/jamespwilliams/iso3166_1)

Go library for efficiently accessing the country information specified in
ISO3166-1.

#### Performance

Performance when looking up country information for a given Alpha-2 or Alpha-3
code:

| Module                                                  | Result (Alpha-2)      | Result (Alpha-3) |
|---------------------------------------------------------|-----------------------|------------------|
| iso3166\_1                                              |  19.7 ns/op           |  25.1 ns/op      |
| github.com/launchdarkly/go-country-codes (unmaintained) |  49.3 ns/op           |  49.8 ns/op      |
| github.com/TheBookPeople/iso3166                        |  64.1 ns/op           |  64.8 ns/op      |
| github.com/biter777/countries                           | 200 ns/op             | 253 ns/op        |
| github.com/pariz/gountries                              | 105 ns/op             | 107 ns/op        |

Source:

```
cd _comparison
go test -bench .
```

#### Contributing

Pull requests are welcome.
