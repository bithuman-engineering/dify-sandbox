package main

import "github.com/langgenius/dify-sandbox/internal/core/lib/nodejs"
import "C"

//export DifySeccomp
func DifySeccomp(uid int, gid int) {
	nodejs.InitSeccomp(uid, gid)
}

func main() {}