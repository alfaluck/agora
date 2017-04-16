package agora

import (
	"net/http"
	"github.com/alfaluck/agora/cache"
)

type Messenger struct {
	cache cache.Interface
}

func (m *Messenger) ServeHTTP(w http.ResponseWriter, r *http.Request)  {

}

func (m *Messenger) Cache() cache.Interface {
	return m.cache
}

func NewMessenger(config *Config) (m *Messenger, err error) {
	if config.CacheEnabled {
		if m.cache, err = cache.New(config.CacheHandler); err != nil {
			return nil, err
		}
	}

	return m, nil
}
