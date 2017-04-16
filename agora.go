package agora

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

type Config struct {
	Host string
	Port string
	CacheEnabled bool
	CacheHandler map[string]string
}

type App struct {
	Config *Config
	Messenger *Messenger
}

func (a *App) ListenAndServe() error {
	addr := a.Config.Host+":"+a.Config.Port
	return http.ListenAndServe(addr, a.Messenger)
}

func NewApp(configJSON string) (*App, error) {
	conFile, err := os.Open(configJSON)
	if err != nil {
		return nil, err
	}
	defer conFile.Close()
	data, err := ioutil.ReadAll(conFile)
	if err != nil {
		return nil, err
	}

	app := new(App)
	err = json.Unmarshal(data, app.Config)
	if err != nil {
		return nil, err
	}

	if app.Messenger, err = NewMessenger(app.Config); err != nil {
		return nil, err
	}

	return app, nil
}
