package eventstream

import "testing"
import "time"

var frozenMoment = time.Now()

type FrozenClock struct{}

func (f FrozenClock) Now() time.Time {
	return frozenMoment
}

func TestInsert(t *testing.T) {
	SetDefaults()
	SetClock(FrozenClock{})

	increments := 10
	for i := 0; i < increments; i++ {
		Increment()
	}

	scount := NumLastSecond()
	if increments != scount {
		t.Errorf("Expected %d, actual %d", increments, scount)
	}

	mcount := NumLastMinute()
	if increments != mcount {
		t.Errorf("Expected %d, actual %d", increments, mcount)
	}

	hcount := NumLastHour()
	if increments != hcount {
		t.Errorf("Expected %d, actual %d", increments, hcount)
	}
}
