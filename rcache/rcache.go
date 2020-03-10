// Copyright © 2020 Uzhinskiy Boris <boris.ujinsky@gmail.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rcache

import (
	_ "fmt"
	"sync"
)

const (
	//default capacity
	CAP = 1000
)

type valtype []byte

type Rcache struct {
	//mu sync.Mutex
	sync.Mutex
	Namespace string
	Capacity  int64
	Values    map[string]valtype
	Index     []map[string]string
	keys      []string
	clen      int64
}

func New(namespace string, capacity int64) *Rcache {
	c := new(Rcache)
	c.Namespace = namespace

	// default value to capacity
	if capacity <= 0 {
		capacity = int64(CAP)
	}
	c.Capacity = capacity
	c.Values = make(map[string]valtype)
	c.Index = make([]map[string]string, 0)
	c.keys = make([]string, 0)
	c.clen = 0
	return c
}

func (c *Rcache) Add(key string, value valtype) {
	//c.mu.Lock()
	//defer c.mu.Unlock()
	if c.clen >= c.Capacity {
		c.deloldest()
	}
	// TODO: сделать проверку на дубликаты ключей
	c.keys = append(c.keys, key)
	c.Values[key] = value
	c.clen++
}

func (c *Rcache) Get(key string) (valtype, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.Values[key]
	return val, ok
}

func (c *Rcache) deloldest() {
	var rk string
	rk, c.keys = c.keys[0], c.keys[1:]
	if rk != "" {
		delete(c.Values, rk)
		c.clen--
	}
}

func (c *Rcache) Del(key string) {
	//c.mu.Lock()
	//defer c.mu.Unlock()
	delete(c.Values, key)
	c.keys = remove(c.keys, key)
	c.clen--
}

func (c *Rcache) Purge() {
	c.Capacity = 0
	c.Namespace = ""
	c.Values = make(map[string]valtype)
	c.Index = make([]map[string]string, 0)
	c.keys = make([]string, 0)
	c.clen = 0
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
