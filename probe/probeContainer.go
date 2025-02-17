package probe

import (
	"encoding/json"
	"health-probe/models"
	"net/http"
)

type ProbeContainer interface {
	GetLocalProbes() []LocalProbe
	GetDependencyProbes() []DependencyProbe
}

func WriteProbes(c ProbeContainer, w http.ResponseWriter) {
	dependencyProbes := c.GetDependencyProbes()
	localProbes := c.GetLocalProbes()

	var dependencyProbesJson []models.DependencyProbe
	var localProbesJson []models.LocalProbe

	for _, probe := range dependencyProbes {
		dependencyProbesJson = append(dependencyProbesJson, models.DependencyProbe{
			ClientName:     probe.GetClientName(),
			ServerName:     probe.GetServerName(),
			TotalFailed:    probe.GetErrorCount(),
			TotalCompleted: probe.GetSuccessCount(),
		})

		probe.Reset()
	}

	for _, probe := range localProbes {
		localProbesJson = append(localProbesJson, models.LocalProbe{
			Name:           probe.GetName(),
			TotalFailed:    probe.GetErrorCount(),
			TotalCompleted: probe.GetSuccessCount(),
		})

		probe.Reset()
	}

	stat := models.Health{
		DependencyProbes: dependencyProbesJson,
		LocalProbes:      localProbesJson,
	}

	err := json.NewEncoder(w).Encode(stat)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
