package health

import (
	"time"
)

func (h *healthChecker) Report() (Report, error) {
	t := time.Now().UTC()

	l, ld, lerr := collect(h.liveness)
	r, rd, rerr := collect(h.readiness)

	rpt := Report{
		ReportAsOf: &t,
		Duration:   ld + rd,
		Interval:   0,
		Liveness:   l,
		Readiness:  r,
	}

	if lerr != nil {
		return rpt, lerr
	}

	return rpt, rerr
}

func (h *healthChecker) Liveness() (results []Result, err error) {
	results, _, err = collect(h.liveness)
	return
}

func (h *healthChecker) Readiness() (results []Result, err error) {
	results, _, err = collect(h.liveness, h.readiness)
	return
}