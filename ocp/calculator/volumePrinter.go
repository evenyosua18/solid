package calculator

import (
	"encoding/json"
	"fmt"
)

type VolumePrinter struct{}

func (vp *VolumePrinter) JSON(result VolumeResult) []byte {
	volumeJson, err := json.Marshal(result)

	if err != nil {
		panic(err)
	}

	return volumeJson
}

func (vp *VolumePrinter) PrintVolume(result VolumeResult) {
	for key, value := range result {
		fmt.Printf("%s: %f\n", key, value)
	}
}
