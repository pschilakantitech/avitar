package timer

import (
	"bytes"
	"encoding/gob"
	"time"
)

// Timer holds the processing time details, the main data item that will travel through the app.
type Timer struct {
	ObservedAt,
	EmittedAt,
	SubmittedAt,
	StartedAt,
	CompletedAt time.Time
}

// NewTimer returns a newly initialized Timer.
func NewTimer() Timer {
	return Timer{SubmittedAt: time.Now()}
}

// NewTimerFromBytes deserializes a Timer from Gob.
func NewTimerFromBytes(data []byte) (Timer, error) {
	buf := bytes.NewReader(data)
	dec := gob.NewDecoder(buf)
	timer := Timer{}
	err := dec.Decode(&timer)
	return timer, err
}

// ToBytes serializes a Timer to Gob.
func (t Timer) ToBytes() ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(t)
	return buf.Bytes(), err
}

// WaitTime returns the time the job waited in the queue before being picked up.
func (t Timer) WaitTime() time.Duration {
	return t.StartedAt.Sub(t.SubmittedAt)
}

// WaitExt returns the duration the event waited externally (the difference between when
// it was available to when it was emitted).
func (t Timer) WaitExt() time.Duration {
	return t.EmittedAt.Sub(t.ObservedAt)
}

// WaitInt returns the duration the event waited internally (since it was emitted until
// it was started).
func (t Timer) WaitInt() time.Duration {
	return t.StartedAt.Sub(t.EmittedAt)
}

// RunTime returns the job duration.
func (t Timer) RunTime() time.Duration {
	return t.CompletedAt.Sub(t.StartedAt)
}
