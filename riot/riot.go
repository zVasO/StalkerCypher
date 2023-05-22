package riot

import (
	"StalkerCypher/config"
	"fmt"
	"github.com/Kyagara/equinox"
	"sync"
)

var (
	riotClient *equinox.Equinox
	once       sync.Once
)

// GetRiotClient singleton for riot api
func GetRiotClient() (*equinox.Equinox, error) {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	once.Do(func() {
		riotClient, err = equinox.NewClient(config.RiotToken)
		if err != nil {
			fmt.Println(err.Error())
		}
	})
	return riotClient, err
}
