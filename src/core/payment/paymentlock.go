package payment

import (
	"sync"
	"time"
)

// Need to lock when deleting or adding

type MapLock struct {
	lock  sync.RWMutex
	locks map[string]*PaymentLock
}

func (l *MapLock) Set(key string, pl *PaymentLock) {
	l.lock.Lock()
	defer l.lock.Unlock()

	pl.LastUpdated = time.Now().UTC()

	l.locks[key] = pl
}

func (l *MapLock) Get(key string) (*PaymentLock, bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	pl, ok := l.locks[key]
	if !ok {
		pl = &PaymentLock{}
	}

	pl.LastAccessed = time.Now().UTC()
	return pl, ok
}

func (l *MapLock) UnlockPayment(username string, pl *PaymentLock) {
	pl.Unlock()
	l.Set(username, pl)
}

type PaymentLock struct {
	sync.RWMutex
	LastAccessed time.Time
	LastUpdated  time.Time
}
