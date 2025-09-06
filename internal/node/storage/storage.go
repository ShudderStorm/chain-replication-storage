package storage

import "sync"

type Storage struct {
	mutex   sync.RWMutex
	data    map[string]string
	writeCh <-chan WriteRequest
	readCh  <-chan ReadRequest
}

type Option func(*Storage)

func WithSize(size int) Option {
	return func(storage *Storage) {
		storage.data = make(map[string]string, size)
	}
}

func New(options ...Option) *Storage {
	storage := &Storage{
		data: make(map[string]string),
	}

	for _, option := range options {
		option(storage)
	}

	return storage
}

func (storage *Storage) ListenOnWrite(stream WriteRequestStream) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	storage.writeCh = stream()
}

func (storage *Storage) ListenOnRead(stream ReadRequestStream) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	storage.readCh = stream()
}

func (storage *Storage) write(key, value string) bool {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()
	storage.data[key] = value
	return true
}

func (storage *Storage) read(key string) (string, bool) {
	storage.mutex.RLock()
	defer storage.mutex.RUnlock()
	value, ok := storage.data[key]
	return value, ok
}

func (storage *Storage) Run() {
	for {
		select {
		case writeRequest := <-storage.writeCh:
			success := storage.write(writeRequest.key, writeRequest.value)
			writeRequest.callback <- WriteResponse{success}
		case readRequest := <-storage.readCh:
			value, success := storage.read(readRequest.key)
			readRequest.callback <- ReadResponse{success, value}
		}
	}
}
