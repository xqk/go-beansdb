# go-beansdb

[![tag](https://img.shields.io/github/tag/xqk/go-beansdb.svg)](https://github.com/xqk/go-beansdb/releases)
![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.17-%23007d9c)
[![License](https://img.shields.io/github/license/xqk/go-beansdb)](./LICENSE)

✨ **`xqk/go-beansdb` 是go版本的beansdb客户端.**

## 💡 用法

```go
import (
    "github.com/xqk/go-beansdb/memcache"
)
```

```go
sch := NewAutoScheduler([]string{"127.0.0.1:7901"}, 16)
client := NewRClient(sch, 1, 1, 1)
item, targets, err := client.Get("foo")
fmt.Println(string(item.Body))
```