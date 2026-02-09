package utils

import "sync"

func NewSafeUnLocker() SafeUnLocker {
	s := safeUnLockAndTryLocker(make(chan struct{}, 1))
	return &s
}

type TryLocker interface {
	TryLock() (l bool)
}

type SafeUnLocker interface {
	sync.Locker
	TryLocker
}

type safeUnLockAndTryLocker chan struct{}

func (u *safeUnLockAndTryLocker) TryLock() (l bool) {
	select {
	case *u <- struct{}{}:
		return true
	default:
		return false
	}
}

func (u *safeUnLockAndTryLocker) Lock() {
	*u <- struct{}{}
}

func (u *safeUnLockAndTryLocker) Unlock() {
	select {
	case <-(*u):
	default:
	}
}
