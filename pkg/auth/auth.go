package auth

import (
	"github.com/pelletier/go-toml"
	"io"
	"io/ioutil"
)

type Auth struct {
	Registry string `toml:"registry"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type AuthConfig struct {
	Auths []Auth `toml:"auths"`
}

func ParseAuthConfig(path string) (*AuthConfig, error) {
	authConfig := &AuthConfig{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(data, authConfig)
	if err != nil {
		return nil, err
	}

	return authConfig, nil
}

func ParseAuthConfigFromReader(reader io.Reader) (*AuthConfig, error) {
	authConfig := &AuthConfig{}
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal(data, authConfig)
	if err != nil {
		return nil, err
	}

	return authConfig, nil
}

func checkRegistryAddress(registry string) {

}
