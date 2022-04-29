package calculator

import "fmt"

type VolumePrinter struct{}

func (vp *VolumePrinter) Print(result VolumeResult) {
	for key, value := range result {
		fmt.Println(key, " : ", value)
	}
}
