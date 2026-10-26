// Package behavior owns Eva's quiet, predictable desk routine. It never
// invents a walking, standing or jumping pose, and it never moves windows.
package behavior

import "time"

type Pose int

const (
	Sitting Pose = iota
	Curled
	Horizontal
)

type phase struct {
	pose     Pose
	duration time.Duration
}

var phases = [...]phase{
	{Sitting, 60 * time.Second},
	{Horizontal, 90 * time.Second},
	{Curled, 120 * time.Second},
	{Horizontal, 90 * time.Second},
}

type Controller struct {
	pose       Pose
	autonomous bool
	phase      int
	next       time.Time
}

func New(now time.Time) *Controller {
	return &Controller{pose: Sitting, autonomous: true, next: now.Add(phases[0].duration)}
}

func (c *Controller) Pose() Pose       { return c.pose }
func (c *Controller) Autonomous() bool { return c.autonomous }

// Step advances at most one phase to avoid startling jumps after suspension.
func (c *Controller) Step(now time.Time) bool {
	if !c.autonomous || now.Before(c.next) {
		return false
	}
	c.phase = (c.phase + 1) % len(phases)
	c.next = now.Add(phases[c.phase].duration)
	before := c.pose
	c.pose = phases[c.phase].pose
	return before != c.pose
}

func (c *Controller) Select(p Pose, now time.Time) {
	if p != Sitting && p != Curled && p != Horizontal {
		return
	}
	c.autonomous = false
	c.pose = p
	c.phase = phaseForPose(p)
	c.next = now.Add(phases[c.phase].duration)
}

func (c *Controller) EnableAutonomy(now time.Time, enabled bool) {
	c.autonomous = enabled
	if enabled {
		c.phase = phaseForPose(c.pose)
		c.next = now.Add(phases[c.phase].duration)
	}
}

// Touch wakes Eva without sliding her bitmap or her desktop position.
func (c *Controller) Touch(now time.Time) {
	c.pose = Sitting
	c.phase = 0
	c.next = now.Add(phases[0].duration)
}

func (c *Controller) Breathes() bool {
	return c.pose == Sitting || c.pose == Curled
}

func phaseForPose(p Pose) int {
	switch p {
	case Curled:
		return 2
	case Horizontal:
		return 1
	default:
		return 0
	}
}
