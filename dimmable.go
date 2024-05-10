package main

import (
	"math"
	"time"
)

type Dimmable struct {
	Device
	Target         float64 `json:"target"`
	Step           float64 `json:"-"`
	Min            int     `json:"-"`
	Max            int     `json:"-"`
	LastSent       int     `json:"-"`
	transition     bool
	TransitionTime int
}

func (d *Dimmable) getMin() int {
	return d.Min
}

func (d *Dimmable) getMax() int {
	return d.Max
}

func (d *Dimmable) getStep() float64 {
	return d.Step
}

func (d *Dimmable) setStep(step float64) {
	d.Step = step
}

func (d *Dimmable) getLastSent() int {
	return d.LastSent
}

func (d *Dimmable) setLastSent(lastSent int) {
	d.LastSent = lastSent
}

func (d *Dimmable) getLastChanged() *time.Time {
	return d.LastChanged
}

func (d *Dimmable) setLastChanged(lastSent *time.Time) {
	d.LastChanged = lastSent
}

func (d *Dimmable) getTarget() float64 {
	return d.Target
}

func (d *Dimmable) setTarget(target float64) {
	d.Target = target
}

func (d *Dimmable) processRequest(request SwitchRequest) {
	request.Value = math.Min(request.Value, float64(d.getMax()))
	request.Value = math.Max(request.Value, float64(d.getMin()))

	d.setTarget(request.Value)

	if d.transition {
		d.TransitionTime = request.Duration
		d.setStep(0)
		return
	}

	diff := int(math.Abs(d.getCurrent() - request.Value))
	var step float64
	cycles := request.Duration * 1000 / cycleLength
	if request.Duration == 0 {
		step = float64(diff)
	} else {
		step = float64(diff) / float64(cycles)
	}
	d.setStep(step)
}

func (d *Dimmable) processRequestChild(request SwitchRequest) {
	d.processRequest(request)
}

func (d *Dimmable) UpdateValue() (float64, bool) {
	current := d.getCurrent()
	if current != d.Target {
		if d.transition {
			d.setCurrent(d.Target)
			return d.Target, true
		}
		if d.Step == 0 {
			d.setStep(100)
		}
		if current > d.Target {
			current -= d.Step
			current = math.Max(current, d.Target)
		} else {
			current += d.Step
			current = math.Min(current, d.Target)
		}
		d.setCurrent(current)
		return current, true
	}
	return 0, false
}

func (d *Dimmable) UpdateValueChild() (float64, bool) {
	return d.UpdateValue()
}
