package behavior

import (
	"testing"
	"time"
)

func TestAutonomousRoutineRespectsApprovedPoses(t *testing.T) {
	start := time.Date(2026, 10, 25, 16, 0, 0, 0, time.Local)
	c := New(start)
	if !c.Autonomous() || c.Pose() != Sitting {
		t.Fatalf("autonomous launch must start sitting, got %v, on=%v", c.Pose(), c.Autonomous())
	}
	if c.Step(start.Add(54 * time.Second)) {
		t.Fatal("changed pose before the scheduled rest")
	}
	if !c.Step(start.Add(61*time.Second)) || c.Pose() != Horizontal {
		t.Fatalf("expected a horizontal rest, got %v", c.Pose())
	}
	if !c.Step(start.Add(152*time.Second)) || c.Pose() != Curled {
		t.Fatalf("expected a curled rest, got %v", c.Pose())
	}
	if !c.Step(start.Add(273*time.Second)) || c.Pose() != Horizontal {
		t.Fatalf("expected return to horizontal pose, got %v", c.Pose())
	}
	if !c.Step(start.Add(365*time.Second)) || c.Pose() != Sitting {
		t.Fatalf("expected wakeful sitting pose, got %v", c.Pose())
	}
}

func TestManualPoseSuspendsAutonomyUntilEnabled(t *testing.T) {
	start := time.Now()
	c := New(start)
	c.Select(Curled, start)
	if c.Autonomous() {
		t.Fatal("manually selected rest must not be overridden")
	}
	if c.Step(start.Add(24*time.Hour)) || c.Pose() != Curled {
		t.Fatal("manual pose was changed by automatic timer")
	}
	c.EnableAutonomy(start.Add(24*time.Hour), true)
	if !c.Autonomous() || c.Pose() != Curled {
		t.Fatal("resuming autonomy must not immediately replace manually selected pose")
	}
}

func TestPettingWakesEvaWithoutMovingHerAutomatically(t *testing.T) {
	start := time.Now()
	c := New(start)
	c.Step(start.Add(61 * time.Second))
	c.Step(start.Add(152 * time.Second))
	if c.Pose() != Curled {
		t.Fatal("test setup expected resting pose")
	}
	c.Touch(start.Add(160 * time.Second))
	if c.Pose() != Sitting {
		t.Fatal("petting Eva should wake her gently into the approved sitting pose")
	}
	if c.Step(start.Add(180 * time.Second)) {
		t.Fatal("should keep responding to person briefly without switching poses")
	}
}

func TestBreathingNeverDistortsHorizontalPose(t *testing.T) {
	start := time.Now()
	c := New(start)
	if !c.Breathes() {
		t.Fatal("sitting Eva must have an unobtrusive idle breath")
	}
	c.Select(Horizontal, start)
	if c.Breathes() {
		t.Fatal("approved horizontal pose must remain completely still until tail is separately validated")
	}
}
