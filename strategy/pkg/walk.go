package pkg

import "log"

type WalkStrategy struct {
}

func (w *WalkStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 4
	total := endPoint - startPoint
	totalTime := total * 60

	log.Printf(
		"Walk A:[%d] to B:[%d] AVG speed [%d] Total [%d] Total time [%d] min \n",
		startPoint,
		endPoint,
		avgSpeed,
		total,
		totalTime,
	)
}
