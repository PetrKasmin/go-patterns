package pkg

import "log"

type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 30
	trafficJam := 2
	total := endPoint - startPoint
	totalTime := total * avgSpeed * trafficJam

	log.Printf(
		"Road A:[%d] to B:[%d] AVG speed [%d] Traffic [%d] Total [%d] Total time [%d] min \n",
		startPoint,
		endPoint,
		avgSpeed,
		trafficJam,
		total,
		totalTime,
	)
}
