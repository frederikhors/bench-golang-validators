# bench-golang-validators

`go test -bench .`

```
goos: windows
goarch: amd64
pkg: benchGolangValidators
BenchmarkAsaskevich-8              94460             12505 ns/op            5439 B/op         72 allocs/op
BenchmarkPlayground-8             704642              1784 ns/op             227 B/op         14 allocs/op
BenchmarkSaddam-8                  97534             12125 ns/op            1858 B/op         68 allocs/op
BenchmarkBuffalo-8                235267              5033 ns/op             756 B/op         15 allocs/op
BenchmarkOzzo-8                    14416             81325 ns/op            5653 B/op        112 allocs/op
PASS
ok      benchGolangValidators   7.300s
```
