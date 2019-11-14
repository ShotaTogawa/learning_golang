# Remember to BET
 - Benchmark
 - Example
 - Test
 
 ```cassandraql
BecnhmarkCat(b *testnig.B)
ExampleCat()
TestCat(t *testing.T)
```

# Commands
```cassandraql
go test // test実行
go test -bench . // 実行速度を図る
go test -cover // testのカバー率をだす
go test -coverprofile c.out 
go tool cover -html=c.out 


godoc -http:=8080 
```