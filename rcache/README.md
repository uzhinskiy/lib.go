# rcache

## Внимание!
Что-бы избежать проблем с DATARACE нужно явно вызывать блокировку кэша:

    cache := rcache.New("mycache", 10)
    cache.Lock()
    cache.Add("key_name", []byte("value"))
    cache.Unlock()

