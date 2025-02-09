# Price Container
A solution is based on the hash maps, min heap and max heap.

For each origin in MaxHeap we always store the best 10 prices and in MinHeap we store every other price.

Since we are using heap the algorithmic complexity of inserting a new price is `O(log(n))`, where `n = 999` - total number companies. Thus the algorithmic complexity is `O(1)`.

Because on each insertion we recalculate the mean price, processing GET request takes constant time `O(1)`

## run
```
go run main.go
```

## test
```
go test ./...
```