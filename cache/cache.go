package cache

import (
	"errors"
	"time"
)

type Interface interface {
	Configure(map[string]string) error
	GetItem(string) (*Item, error)
	GetItems([]string) ([]*Item, error)
	HasItem(string) (bool, error)
	Clear() error
	DeleteItem(string) error
	DeleteItems([]string) error
	Save(*Item) error
	SaveDeferred(*Item) error
	Commit() error
}

type Provider struct{}

func (h *Provider) Configure(config map[string]string) error {
	return errors.New("Not implemented Configure() method")
}

func (h *Provider) GetItem(key string) (*Item, error) {
	return nil, errors.New("Not implemented GetItem() method")
}

func (h *Provider) GetItems(keys []string) ([]*Item, error) {
	return nil, errors.New("Not implemented GetItems() method")
}

func (h *Provider) HasItem(key string) (bool, error) {
	return false, errors.New("Not implemented HasItem() method")
}

func (h *Provider) Clear() (error) {
	return errors.New("Not implemented Clear() method")
}

func (h *Provider) DeleteItem(key string) (error) {
	return errors.New("Not implemented DeleteItem() method")
}

func (h *Provider) DeleteItems(keys []string) (error) {
	return errors.New("Not implemented DeleteItems() method")
}

func (h *Provider) Save(item *Item) (error) {
	return errors.New("Not implemented Save() method")
}

func (h *Provider) SaveDeferred(item *Item) (error) {
	return errors.New("Not implemented SaveDeferred() method")
}

func (h *Provider) Commit() (error) {
	return errors.New("Not implemented Commit() method")
}

var providers map[string]Interface

type Item struct {
	key string
	data interface{}
	hit bool
	expiration int64
}

func (i *Item) Key() string {
	return i.key
}

func (i *Item) Data() interface{} {
	return i.data
}

func (i *Item) Hit() bool {
	return i.hit
}

func (i *Item) SetKey(key string) *Item {
	i.key = key
	return i
}

func (i *Item) SetData(data interface{}) *Item {
	i.data = data
	return i
}

func (i *Item) ExpiresAt(datetime time.Time) *Item {
	i.expiration = datetime.Unix()
	return i
}

func (i *Item) ExpiresAfter(interval time.Duration) *Item {
	i.expiration = time.Now().Add(interval).Unix()
	return i
}


func New(config map[string]string) (h Interface, err error) {
	var driver string
	var ok bool

	if driver, ok = config["driver"]; !ok {
		return nil, errors.New("Undefined cacheHandler:driver config param")
	}
	if h, ok = providers[driver]; !ok {
		return nil, errors.New("Can't find driver "+driver)
	}
	if err := h.Configure(config); err != nil {
		return nil, errors.New("Can't initialize provider")
	}

	return h, nil
}

func init() {
	providers = make(map[string]Interface, 0)
}

func Register(key string, provider Interface) error {
	providers[key] = provider
	return nil
}