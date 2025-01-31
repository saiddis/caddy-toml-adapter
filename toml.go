package caddytomladapter

import (
	"encoding/json"
	"log"

	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/pelletier/go-toml/v2"
)

func init() {
	caddyconfig.RegisterAdapter("toml", Adapter{})
}

type Adapter struct{}

func (a Adapter) Adapt(body []byte, m map[string]interface{}) ([]byte, []caddyconfig.Warning, error) {
	if err := toml.Unmarshal(body, &m); err != nil {
		return nil, nil, err
	}

	delete(m, "filename")
	b, err := json.Marshal(m)
	log.Printf("config: %s", b)

	return b, nil, err
}
