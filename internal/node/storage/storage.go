package storage

import "sync"

type Storage struct {
	mutex sync.RWMutex
	data  map[string]string

	storeCh chan storeRequest
	loadCh  chan loadRequest
}

func (storage *Storage) store(req storeRequest) {
	storage.mutex.Lock()
	defer storage.mutex.Unlock()

	storage.data[req.key] = req.value
	req.respCh <- storeResponse{
		success: true,
	}
}

func (storage *Storage) load(req loadRequest) {
	storage.mutex.RLock()
	defer storage.mutex.RUnlock()

	value, success := storage.data[req.key]
	req.respCh <- loadResponse{
		value: value, success: success,
	}
}

func (storage *Storage) Store(key string, value string) bool {
	respCh := make(chan storeResponse)

	storage.storeCh <- storeRequest{
		key: key, value: value, respCh: respCh,
	}

	response := <-respCh
	return response.success
}

func (storage *Storage) Load(key string) (string, bool) {
	respCh := make(chan loadResponse)

	storage.loadCh <- loadRequest{
		key: key, respCh: respCh,
	}

	response := <-respCh
	return response.value, response.success
}

func New() *Storage {
	storage := &Storage{
		data:    make(map[string]string),
		storeCh: make(chan storeRequest),
		loadCh:  make(chan loadRequest),
	}

	go func() {
		for {
			select {
			case req := <-storage.storeCh:
				storage.store(req)
			case req := <-storage.loadCh:
				storage.load(req)
			}
		}
	}()

	return storage
}
