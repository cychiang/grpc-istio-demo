package server

import "sync"

// Create a sync Map to store message
type Server struct {
	sync.Map
}