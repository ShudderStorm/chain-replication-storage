package storage

type storeRequest struct {
	key, value string
	respCh     chan storeResponse
}

type loadRequest struct {
	key    string
	respCh chan loadResponse
}
