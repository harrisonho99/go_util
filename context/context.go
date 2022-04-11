/**
 *
 * A simple package reimplement stadard context library
 *
 */

package context

import (
	"reflect"
	"sync"
	"time"
)

type noValue struct{}
type any interface{}

// context type
type Context interface {
	Done() <-chan noValue                    //close chan when context is canceled or timeout
	Err() error                              //error indicate why it cancel
	Deadline() (deadline time.Time, ok bool) //deadline is time context should cancel, if deadline is set (ok = true)
	Value(key any) any                       //context can carry some value,get value from context
}

// error type
type canceledError struct{}

func (c canceledError) Error() string {
	return "Context canceled"
}

type deadlineExceededError struct{}

func (d deadlineExceededError) Error() string {
	return "Context deadline exceeded"
}

// implement default context
type emptyContext struct{}

func (e *emptyContext) Done() <-chan noValue {
	return nil
}
func (e *emptyContext) Err() error {
	return nil
}
func (e *emptyContext) Deadline() (deadline time.Time, ok bool) {
	return
}
func (e *emptyContext) Value(key any) any {
	return nil
}

//Return empty context
//should be wrapped by other context
func Background() Context {
	return new(emptyContext)
}

//Function invoke to cancel context
type CancelFunc func()

// ** Cancel context
type CancelContext struct {
	Context
	done chan noValue
	err  error
}

func (c *CancelContext) Done() <-chan noValue {
	return c.done
}

func (c *CancelContext) Err() error {
	return c.err
}

//Timeout context wrapper
func WithCancel(parrent Context) (Context, CancelFunc) {
	ctx := &CancelContext{parrent, make(chan noValue), nil}
	var cancel CancelFunc = func() {
		close(ctx.done)
		ctx.err = canceledError{}
	}
	return ctx, cancel
}

// ** Timeout context
type TimeoutContext struct {
	Context
	done     chan noValue
	isDone   bool
	err      error
	deadline time.Time
}

func (c *TimeoutContext) Done() <-chan noValue {
	return c.done
}

func (c *TimeoutContext) Err() error {
	return c.err
}
func (c *TimeoutContext) Deadline() (deadline time.Time, ok bool) {
	return c.deadline, true
}

//Cancel context wrapper
func WithTimeout(parrent Context, duration time.Duration) (Context, CancelFunc) {
	ctx := &TimeoutContext{parrent, make(chan noValue), false, nil, time.Now().Add(duration)}

	//Alarm
	go func() {
		select {
		case <-time.After(duration):
			if !ctx.isDone {
				var mu sync.Mutex
				mu.Lock()
				ctx.isDone = true
				ctx.err = deadlineExceededError{}
				mu.Unlock()
				close(ctx.done)
			}
			return
		case <-ctx.Done():
			return
		}
	}()

	var cancelFunc CancelFunc = func() {
		var mu sync.Mutex
		mu.Lock()
		defer mu.Unlock()
		if !ctx.isDone {
			ctx.err = canceledError{}
			close(ctx.done)
		}
	}

	return ctx, cancelFunc
}

// ** Value context
// we dont use map here because map is unsafe for concurency
type ValueContext struct {
	Context
	key, value any
}

func (v ValueContext) Value(key any) any {
	var mu sync.RWMutex
	mu.RLock()
	defer mu.RUnlock()

	return find(v, key)
}

//recursive find
func find(v Context, want any) any {
	if v, ok := v.(ValueContext); ok {
		if want == v.key {
			return v.value
		}
		return find(v.Context, want)
	} else {
		return nil
	}
}

//carry Key-Value pair
//key should be comparable
func WithValue(ctx Context, key any, value any) Context {
	key_v := reflect.ValueOf(key)
	if !key_v.Type().Comparable() {
		panic("context value : key is not comparable")
	}
	if key == nil {
		panic("key is nil")
	}

	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	return ValueContext{ctx, key, value}
}
