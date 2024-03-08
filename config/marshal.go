package config

import (
	"encoding/json"
)

type configMainTemp struct {
	Services      []*ConfigProxyService
	PrivateConfig *Something
}

var (
	_ json.Marshaler   = (*configMain)(nil)
	_ json.Unmarshaler = (*configMain)(nil)
)

func (c *configMain) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		configMainTemp{
			Services:      c.Services,
			PrivateConfig: c.PrivateConfig,
		},
	)
}

func (c *configMain) UnmarshalJSON(data []byte) (err error) {
	configTemp := configMainTemp{}
	err = json.Unmarshal(data, &configTemp)
	if err != nil {
		return err
	}

	c.Services = configTemp.Services
	c.PrivateConfig = configTemp.PrivateConfig
	return nil
}
