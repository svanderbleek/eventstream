package eventstream

import "time"

// Store
type Store interface {
	Increment()
}

// Query
type Query interface {
	NumLastSecond() int
	NumLastMinute() int
	NumLastHour() int
}

//Clock
type Clock interface {
	Now() time.Time
}

var store Store
var query Query
var clock Clock

func Increment() {
	store.Increment()
}

func NumLastSecond() int {
	return query.NumLastSecond()
}

func NumLastMinute() int {
	return query.NumLastMinute()
}

func NumLastHour() int {
	return query.NumLastHour()
}

func SetClock(c Clock) {
	clock = c
}

type StandardClock struct{}

func (c StandardClock) Now() time.Time {
	return time.Now()
}

func SetStore(s Store) {
	store = s
}

func SetQuery(q Query) {
	query = q
}

type NaiveEventStore struct {
	events []time.Time
}

func (b NaiveEventStore) Increment() {
	b.events = append(b.events, clock.Now())
}

func (b NaiveEventStore) NumLastSecond() int {
	return 0
}

func (b NaiveEventStore) NumLastMinute() int {
	return 0
}

func (b NaiveEventStore) NumLastHour() int {
	return 0
}

func SetDefaults() {
	var naive = NaiveEventStore{}
	SetStore(naive)
	SetQuery(naive)
	SetClock(StandardClock{})
}
