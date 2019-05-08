```
$ go run main.go
2019-05-08 13:17:51 +0900 JST
0001-01-01 00:00:00 +0000 UTC
2019-05-08 13:17:51 +0900 JST
```
```
$ go run -tags prod main.go
# command-line-arguments
./main.go:12:2: undefined: mytime.FixTime
./main.go:14:2: undefined: mytime.Reset
```
