module github.com/jamespwilliams/iso3166_1_comparison

go 1.15

replace github.com/jamespwilliams/iso3166_1 => ../

require (
	github.com/TheBookPeople/iso3166 v0.0.0-20160919111731-b671ca53a641
	github.com/biter777/countries v1.3.4
	github.com/jamespwilliams/iso3166_1 v0.0.0-00010101000000-000000000000
	github.com/launchdarkly/go-country-codes v0.0.0-20191008001159-776cf5214f39
	github.com/pariz/gountries v0.0.0-20200430155801-1c6a393df9c7
	github.com/tchap/go-patricia v2.3.0+incompatible // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
