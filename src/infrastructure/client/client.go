package client

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Oleg-Skalozub/testtask/src/domain/entity"
	"github.com/Oleg-Skalozub/testtask/src/infrastructure/config"
	"github.com/Oleg-Skalozub/testtask/src/infrastructure/errscan"
)

//
var Client ClientInterface

// ClientInterface ...
type ClientInterface interface {
	Get(path string, day, month int) (entity.Contain, error)
}

type client struct {
	config config.Configuration
}

// NewClient ...
func NewClient() ClientInterface {
	return client{
		config: config.Config,
	}
}

// Get ...
func (c client) Get(path string, day, month int) (entity.Contain, error) {

	url := path + "/" + strconv.Itoa(day) + "/" + strconv.Itoa(month)
	resp, err := http.Get(url)
	if err != nil {
		return entity.Contain{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entity.Contain{}, errscan.WrongStatusCodeError
	}

	res := entity.Contain{}
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return entity.Contain{}, nil
	}

	return res, nil
}
