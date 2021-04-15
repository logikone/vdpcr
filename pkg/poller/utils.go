package poller

func isUpOrDown(code int) float64 {
	if code >= 200 && code <= 399 {
		return 0
	} else {
		return 1
	}
}
