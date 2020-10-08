package lookup

import (
	"github.com/pkg/errors"
	"sigs.k8s.io/yaml"
)

type LoopkupData struct {
	Info map[string]string
}

//todo lets find a better way to lookup this metadata rather than using a csv in the codebase?
// NewLookupData will return struct containing maps for looking up static health information and suggested fix links
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
