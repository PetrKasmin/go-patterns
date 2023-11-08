package pkg

import "log"

type PublicTransportStrategy struct {
}

func (t *PublicTransportStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 40
	total := endPoint - startPoint
	totalTime := total * avgSpeed

	log.Printf(
		"Tran A:[%d] to B:[%d] AVG speed [%d] Total [%d] Total time [%d] min \n",
		startPoint,
		endPoint,
		avgSpeed,
		total,
		totalTime,
	)
}
