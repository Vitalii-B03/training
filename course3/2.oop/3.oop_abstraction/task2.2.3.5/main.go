package main

import (
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"time"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}
type HashMap struct {
	size     int
	table    []*Backet
	funcHash func(string) uint64
}
type Backet struct {
	key   string
	value interface{}
}

func NewHashMap(size int, options ...func(table *HashMap)) *HashMap {
	m := &HashMap{
		size:  size,
		table: make([]*Backet, size),
	}
	for _, option := range options {
		option(m)
	}
	return m
}
func WithHashCRC64() func(*HashMap) {
	return func(table *HashMap) {
		table.funcHash = func(key string) uint64 {
			return crc64.Checksum([]byte(key), crc64.MakeTable(crc64.ISO))
		}
	}
}
func WithHashCRC32() func(*HashMap) {
	return func(table *HashMap) {
		table.funcHash = func(key string) uint64 {
			return uint64(crc32.ChecksumIEEE([]byte(key)))
		}
	}
}
func WithHashCRC16() func(*HashMap) {
	return func(table *HashMap) {
		table.funcHash = func(key string) uint64 {
			return uint64(crc16([]byte(key)))
		}
	}
}
func crc16(data []byte) uint16 {
	crc := uint16(0xFFFF)
	for _, b := range data {
		crc ^= uint16(b) << 8
		for i := 0; i < 8; i++ {
			if crc&0x8000 != 0 {
				crc = crc<<1 ^ 0x1021
			} else {
				crc <<= 1
			}
		}
	}
	return crc
}
func WithHashCRC8() func(*HashMap) {
	return func(table *HashMap) {
		table.funcHash = func(key string) uint64 {
			return uint64(crc8([]byte(key)))
		}
	}
}
func crc8(data []byte) uint8 {
	crc := uint8(0)
	for _, b := range data {
		crc ^= b
		for i := 0; i < 8; i++ {
			if crc&0x80 != 0 {
				crc = crc<<1 ^ 0x07

			} else {
				crc <<= 1
			}
		}
	}
	return crc
}
func (m *HashMap) Set(key string, value interface{}) {
	hash := m.funcHash(key)
	index := hash % uint64(m.size)
	m.table[index] = &Backet{key, value}
}
func (m *HashMap) Get(key string) (interface{}, bool) {
	hash := m.funcHash(key)
	index := hash % uint64(m.size)
	if m.table[index] != nil && m.table[index].key == key {
		return m.table[index].value, true
	}
	return nil, false
}
func MeassureTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}
func main() {
	m := NewHashMap(16, WithHashCRC64())
	since := MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC32())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC16())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)

	m = NewHashMap(16, WithHashCRC8())
	since = MeassureTime(func() {
		m.Set("key", "value")

		if value, ok := m.Get("key"); ok {
			fmt.Println(value)
		}
	})
	fmt.Println(since)
}
