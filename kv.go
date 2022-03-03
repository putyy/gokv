package gokv

import (
	"strings"
)

type Kv struct {
	// 项目标识
	Project string
	// 功能描述(尾部)
	Tail string
	// 头标识(前缀)
	Prefix string
	// 分隔符
	Separator string
}

func (b *Kv) CompileKey(arr []string, prefix string) string {
	if b.Separator == "" {
		b.Separator = ":"
	}
	k := prefix
	if b.Project != "" {
		k = k + b.Separator + b.Project
	}

	if b.Tail != "" {
		k = k + b.Separator + b.Tail
	}

	if b.Prefix != "" {
		k = b.Prefix + b.Separator + k
	}

	if 0 == len(arr) {
		return k
	}

	return k + b.Separator + strings.Join(arr, b.Separator)
}
