# bench-golang-validators

`go test -bench .`

```
goos: windows
goarch: amd64
pkg: benchGolangValidators
BenchmarkAsaskevich-8   	   82749	     13119 ns/op	    5452 B/op	      72 allocs/op
BenchmarkPlayground-8   	  705242	      1808 ns/op	     228 B/op	      14 allocs/op
BenchmarkSaddam-8       	   93002	     13140 ns/op	    1854 B/op	      68 allocs/op
BenchmarkBuffalo-8      	  203504	      5042 ns/op	     760 B/op	      15 allocs/op
BenchmarkOzzo-8         	  176402	      6674 ns/op	    5093 B/op	      99 allocs/op
PASS
ok      benchGolangValidators   7.300s
```
