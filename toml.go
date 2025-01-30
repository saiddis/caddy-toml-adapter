package caddytomladapter

import (
	"bytes"
	"encoding/json"

	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/pelletier/go-toml/v2"
)

func init() {
	caddyconfig.RegisterAdapter("toml", Adapter{})
}

type Adapter struct{}

func (a Adapter) Adapt(body []byte, m map[string]interface{}) ([]byte, []caddyconfig.Warning, error) {
	buf := bytes.NewReader(body)
	if err := toml.NewDecoder(buf).Decode(&m); err != nil {
		return nil, nil, err
	}

	b, err := json.Marshal(m)

	return b, nil, err
}
