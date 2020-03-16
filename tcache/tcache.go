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

package tcache

import (
	"sync"
)

const (
	//default capacity
	CAP = 1000
)

//type TCAche - holds data in separated Tables
type TCache struct {
	sync.RWMutex
	Tables   map[string]*Table
	Capacity uint64
}

// type Table - describes single table with data
type Table struct {
	Values  map[string][]byte
	Index   []map[string]string
	keys    []string
	counter uint64
	ccap    uint64
}

func tnew(capacity uint64) *Table {
	return &Table{Values: map[string][]byte{}, Index: []map[string]string{}, keys: []string{}, counter: 0, ccap: capacity}
}

func (t *Table) add(key string, value []byte) {
	if t.counter >= t.ccap {
		t.deloldest()
	}
	_, present := t.Values[key]
	if !present {
		t.keys = append(t.keys, key)
	}
	t.Values[key] = value
	t.counter += 1
}

func (t *Table) get(key string) ([]byte, bool) {
	val, ok := t.Values[key]
	return val, ok
}

func (t *Table) deloldest() {
	var rk string
	ckeys := t.keys
	rk, ckeys = ckeys[0], ckeys[1:]
	if rk != "" {
		delete(t.Values, rk)
		t.counter -= 1
	}
}

func (t *Table) del(key string) {
	delete(t.Values, key)
	t.keys = remove(t.keys, key)
	t.counter--
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// func New(capacity ...uint64) *TCache - initialize Cache with capacity
// or with default capacity=1000
func New(capacity ...uint64) *TCache {
	cc := uint64(CAP)
	t := new(TCache)
	if len(capacity) > 0 {
		cc = capacity[0]
	}
	t.Capacity = cc
	t.Tables = map[string]*Table{}
	return t
}

func (c *TCache) Add(table string, key string, value []byte) {
	c.RLock()
	t, t_exist := c.Tables[table]
	c.RUnlock()
	if !t_exist {
		t = tnew(c.Capacity)
	}
	t.add(key, value)
	c.Lock()
	c.Tables[table] = t
	c.Unlock()
}

func (c *TCache) Get(table string, key string) ([]byte, bool) {
	c.RLock()
	t, t_exist := c.Tables[table]
	c.RUnlock()
	if t_exist {
		val, okv := t.get(key)
		return val, okv
	} else {
		return nil, false
	}
}

func (c *TCache) Del(table string, key string) {
	c.Lock()
	t, t_exist := c.Tables[table]
	if t_exist {
		t.del(key)
		c.Tables[table] = t
	}
	c.Unlock()
}

func (c *TCache) GetTables() []string {
	c.RLock()
	var t []string
	for k := range c.Tables {
		t = append(t, k)
	}
	c.RUnlock()
	return t
}

func (c *TCache) List(table string) (*Table, bool) {
	//c.mx.RLock()
	//defer c.mx.RUnlock()
	val, ok := c.Tables[table]

	return val, ok
}

func (c *TCache) PurgeTable(table string) {
	c.Lock()
	delete(c.Tables, table)
	c.Unlock()
}
