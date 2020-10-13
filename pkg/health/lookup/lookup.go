package lookup

import (
	"github.com/pkg/errors"
	"sigs.k8s.io/yaml"
)

// LoopkupData contains bind data from yaml
type LoopkupData struct {
	Info map[string]string
}

// NewLookupData will return struct containing maps for looking up static health information and suggested fix links
// todo lets find a better way to lookup this metadata rather than using a csv in the codebase?
func NewLookupData() (LoopkupData, error) {

	data := LoopkupData{}
	data.Info = make(map[string]string)

	byteArray, err := Asset("pkg/health/lookup/static_data/info.yaml")
	if err != nil {
		return LoopkupData{}, errors.Wrapf(err, "failed to find health status information data")
	}
	err = yaml.Unmarshal(byteArray, &data.Info)
	if err != nil {
		return LoopkupData{}, errors.Wrapf(err, "failed to unmarshal health status information data")
	}

	return data, nil
}
