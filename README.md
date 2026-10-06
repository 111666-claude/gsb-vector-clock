# vector-clock

向量时钟的合并、因果判定与消息去重。只用标准库。

```
go test ./...
go vet ./...
go run ./cmd/vclock --sample=merge
go run ./cmd/vclock --sample=concurrent
go run ./cmd/vclock --sample=dup
go run ./cmd/vclock --sample=work
```

## 口径

- **合并**：收到消息时逐位取两边较大的值，然后本地位自增。
- **因果**：a 严格早于 b 当且仅当每一位都不大于、且至少有一位严格小于。
- **相等**：两边每一位都相同，既不是早于也不是晚于。
- **并发**：互不满足"早于"的两个时钟。
- **幂等**：同一个消息号只处理一次，重复收到不改时钟。

## 不变量

- 合并不会让任何一位变小。
- 并发与相等都不会被判成"早于"。
- `scanned` 不随消息数平方级放大：两千条消息的 `scanned` 不超过 9000。

## 输出契约

每个场景打印一行 JSON；`--sample=work` 只打印 `{"ids": N, "unique": N, "scanned": N}`。
