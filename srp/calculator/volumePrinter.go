package calculator

import "fmt"

type VolumePrinter struct {
}

func (vp *VolumePrinter) PrintVolume(result VolumeResult, geometry string) {
	switch geometry {
	case CONE:
		fmt.Printf("cone volume: %f\n", result.coneVolume)
		break
	case CUBE:
		fmt.Printf("cube volume: %f\n", result.cubeVolume)
		break
	case SPHERE:
		fmt.Printf("sphere volume: %f\n", result.sphereVolume)
		break
	case TUBE:
		fmt.Printf("tube volume: %f\n", result.tubeVolume)
		break
	default:
		fmt.Println("geometries not recognize")
	}
}
