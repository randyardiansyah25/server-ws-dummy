package request

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var Pool = make(map[string](chan string))

//=======================================================================================================

var Table sourceRegistry

type sourceItem struct {
	Item      *gin.Context
	CreatedAt int64
}

func NewTable(ttl int) (table *sourceRegistry) {
	table = &sourceRegistry{source: make(map[string]*sourceItem, 0)}
	go func() {
		for now := range time.Tick(time.Second) {
			table.obj.Lock()
			for key, value := range table.source {
				if now.Unix()-value.CreatedAt > int64(ttl) {
					delete(table.source, key)
				}
			}
			table.obj.Unlock()
		}
	}()
	return
}

type sourceRegistry struct {
	source map[string]*sourceItem
	obj    sync.Mutex
}

func (s *sourceRegistry) Len() int {
	return len(s.source)
}

func (s *sourceRegistry) Add(key string, ctx *gin.Context) {
	s.obj.Lock()
	defer func() {
		s.obj.Unlock()
	}()

	_, ok := s.source[key]
	if !ok {
		newItem := sourceItem{
			Item:      ctx,
			CreatedAt: time.Now().Unix(),
		}
		s.source[key] = &newItem
	}
}

func (s *sourceRegistry) Get(key string) (ctx *gin.Context) {
	s.obj.Lock()
	defer func() {
		s.obj.Unlock()
	}()

	if source, ok := s.source[key]; ok {
		ctx = source.Item
	}

	return
}
