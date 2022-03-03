# gokv
go版本 创建kv

### 使用

> 列
> ps: 可以专门建一个 rediskvs.go 用来存放项目所有用到的键名称， 只需要配置 init 代码即可
```go
package main

import "github.com/putyy/gokv"

var cache = make(map[string]gokv.RedisIns)

func init() {
	cache["user"] = createKv(map[string]string{"tail": "user"})
	cache["register"] = struct{ gokv.RedisKv }{gokv.RedisKv{gokv.Kv{Tail: "u_register"}}}
}

func BuildKv(key string) gokv.RedisIns {
	return cache[key].(gokv.RedisIns)
}

func createKv(r map[string]string) gokv.RedisIns {
	kv := struct{ gokv.RedisKv }{gokv.RedisKv{gokv.Kv{}}}
	for i, v := range r {
		switch i {
		case "project":
			fallthrough
		case "pro":
			kv.RedisKv.Kv.Project = v
		case "tail":
			fallthrough
		case "t":
			kv.RedisKv.Kv.Tail = v
		case "Prefix":
			fallthrough
		case "pre":
			kv.RedisKv.Kv.Prefix = v
		case "separator":
			fallthrough
		case "sep":
			kv.RedisKv.Kv.Separator = v
		}
	}
	return kv
}

func main() {
	fmt.Println(tests.BuildKv("user").GetListKey("1", "2"))
	fmt.Println(tests.BuildKv("register").GetListKey("1", "2"))
}
```