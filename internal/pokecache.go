package pokecache

type Cache struct {
	data map[string][]byte
}

func NewCache() Cache {
	return Cache{
		data: make(map[string][]byte),
	}
}

