package controler

import (
	"encoding/json"
	mod "health-probe/models"
	"log"
	"net/http"
)

type Controler struct {
	services map[string]string
}

const resource = "/health"

func NewControler(services map[string]string) *Controler {
	return &Controler{
		services: services,
	}
}

func (c *Controler) GetStatus() mod.OverallHealth {
	stats := mod.OverallHealth{}
	stats.HealthStats = make(map[string]mod.Health)

	for name, url := range c.services {
		resp, err := http.Get(url + resource)

		if err != nil {
			log.Printf("Error when getting status for %s: %s\n", name, err)
			break
		}

		defer resp.Body.Close()

		var health mod.Health

		if err := json.NewDecoder(resp.Body).Decode(&health); err != nil {
			log.Printf("Error when decoding status for %s: %s\n", name, err)
			break
		}

		stats.HealthStats[name] = health
	}

	return stats
}
