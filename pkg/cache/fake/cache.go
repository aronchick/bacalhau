package fake

// MockCache is a simple map-based cache useable
// in tests without using a real cache.
type FakeCache[T any] struct {
	inner map[string]FakeCacheItem[T]
}

type FakeCacheItem[T any] struct {
	Value  T
	Cost   uint64
	Expiry int64
}

func NewFakeCache[T any]() FakeCache[T] {
	return FakeCache[T]{
		inner: make(map[string]FakeCacheItem[T]),
	}
}

func (m FakeCache[T]) Get(key string) (T, bool) {
	v, ok := m.inner[key]
	if ok {
		return v.Value, ok
	}
	return *new(T), ok
}

func (m FakeCache[T]) Set(key string, value T, cost uint64, expiresInSeconds int64) error {
	m.inner[key] = FakeCacheItem[T]{
		Value:  value,
		Cost:   cost,
		Expiry: expiresInSeconds,
	}
	return nil
}

func (m FakeCache[T]) Delete(key string) {
	delete(m.inner, key)
}

func (m FakeCache[T]) Close() {}
