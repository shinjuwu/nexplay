package captcha

import (
	"errors"
	"sync"
	"time"

	"github.com/mojocn/base64Captcha"
)

var (
	ErrCaptchaExpired  = errors.New("verification code has expired")
	ErrCaptchaNotExist = errors.New("verification code does not exist")
)

// expValue stores timestamp and id of captchas. It is used in the list inside
// memoryStore for indexing generated captchas by timestamp to enable garbage
// collection of expired captchas.
type idByTimeValue struct {
	timestamp time.Time
	id        string
	value     string
}

// memoryStore is an internal store for captcha ids and their values.
type memoryStore struct {
	sync.RWMutex
	digitsById map[string]idByTimeValue
	// Number of items stored since last collection.
	numStored int
	// Number of saved items that triggers collection.
	collectNum int
	// Expiration time of captchas.
	expiration time.Duration
}

// NewMemoryStore returns a new standard memory store for captchas with the
// given collection threshold and expiration time (duration). The returned
// store must be registered with SetCustomStore to replace the default one.
func NewMemoryStore(collectNum int, expiration time.Duration) base64Captcha.Store {
	s := new(memoryStore)
	s.digitsById = make(map[string]idByTimeValue)
	s.collectNum = collectNum
	s.expiration = expiration
	return s
}

func (s *memoryStore) Set(id string, value string) error {
	s.Lock()
	s.digitsById[id] = idByTimeValue{time.Now(), id, value}
	s.numStored++
	s.Unlock()
	if s.numStored > s.collectNum {
		go s.collect()
	}
	return nil
}

func (s *memoryStore) Verify(id, answer string, clear bool) bool {
	v := s.Get(id, clear)
	return v == answer
}

func (s *memoryStore) Get(id string, clear bool) (value string) {
	specifyTime := time.Now()
	if !clear {
		// When we don't need to clear captcha, acquire read lock.
		s.RLock()
		defer s.RUnlock()
	} else {
		s.Lock()
		defer s.Unlock()
	}
	tmp, ok := s.digitsById[id]
	if !ok {
		return
	}
	if clear {
		delete(s.digitsById, id)
	}
	if tmp.timestamp.Add(s.expiration).Before(specifyTime) {
		value = ErrCaptchaExpired.Error()
		return
	}
	value = tmp.value
	return
}

func (s *memoryStore) collect() {
	specifyTime := time.Now()
	s.Lock()
	defer s.Unlock()
	for id, v := range s.digitsById {
		if v.timestamp.Add(s.expiration).Before(specifyTime) {
			delete(s.digitsById, id)
		}
	}
}
