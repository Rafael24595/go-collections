package collection

type DictionaryLimited[T comparable, K any] struct {
	DictionarySync[T, K]
	size     int
	timeline []T
}

func DictionaryLimitedFromMap[T comparable, K any](size int, items map[T]K) *DictionaryLimited[T, K] {
	instance := &DictionaryLimited[T, K]{
		size: size,
		timeline: make([]T, 0),
	}

	instance.items = make(map[T]K)

	count := 0
	for k, v := range items {
		if count == size || len(instance.timeline) > size {
			break
		}
		instance.items[k] = v
		instance.timeline = append(instance.timeline, k)
		count++
	}

	return instance
}

func DictionaryLimitedEmpty[T comparable, K any](size int) *DictionaryLimited[T, K] {
	return DictionaryLimitedFromMap(size, make(map[T]K))
}

func DictionaryLimitedFromVector[T comparable, K any](size int, vector Vector[K], mapper func(K) T) *DictionaryLimited[T, K] {
	return DictionaryLimitedFromList(size, vector.items, mapper)
}

func DictionaryLimitedFromList[T comparable, K any](size int, vector []K, mapper func(K) T) *DictionaryLimited[T, K] {
	mapp := DictionaryLimitedEmpty[T, K](size)
	count := 0
	for _, v := range vector {
		if count == size ||  len(mapp.timeline) > size {
			break
		}
		key := mapper(v)
		mapp.Put(key, v)
		mapp.timeline = append(mapp.timeline, key)
		count++
	}
	return mapp
}

func (c *DictionaryLimited[T, K]) Put(key T, item K) (*K, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	old, exists := c.items[key]
	c.items[key] = item
	c.timeline = append(c.timeline, key)

	if len(c.timeline) > 0 && len(c.timeline) > c.size {
		first := c.timeline[0]
		delete(c.items, first)
		c.timeline = c.timeline[1:len(c.timeline)]
	}

	return &old, exists
}

func (c *DictionaryLimited[T, K]) PutIfAbsent(key T, item K) (*K, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	old, exists :=  c.items[key]
	if exists {
		return &old, exists	
	}

	c.items[key] = item
	c.timeline = append(c.timeline, key)

	if len(c.timeline) > 0 && len(c.timeline) > c.size {
		first := c.timeline[0]
		delete(c.items, first)
		c.timeline = c.timeline[1:len(c.timeline)]
	}

	return &old, exists
}

func (c *DictionaryLimited[T, K]) PutAll(items map[T]K) *DictionaryLimited[T, K] {
	c.mu.Lock()
	defer c.mu.Unlock()

	count := 0
	for key := range items {
		if count == c.size {
			break
		}
		c.items[key] = items[key]
		c.timeline = append(c.timeline, key)
		count++
	}

	for range count {
		if len(c.timeline) > 0 && len(c.timeline) > c.size {
			first := c.timeline[0]
			delete(c.items, first)
			c.timeline = c.timeline[1:len(c.timeline)]
		}
	}

	return c
}
