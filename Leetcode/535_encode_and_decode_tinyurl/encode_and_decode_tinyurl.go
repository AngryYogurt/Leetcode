package main

import (
	"crypto/md5"
	"hash"
)

func main() {

}
/*
 * 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗：3 MB, 在所有 Go 提交中击败了100.00%的用户
 */
type Codec struct {
	m map[string]string
	h hash.Hash
}

func Constructor() Codec {
	return Codec{
		m: make(map[string]string),
		h: md5.New(),
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	st := this.h.Sum([]byte(longUrl))
	_, ok := this.m[string(st)]
	for ok {
		st = this.h.Sum(st)
	}
	this.m[string(st)] = longUrl
	return string(st)
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.m[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
