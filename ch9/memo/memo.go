package memo

import "fmt"

// Func is the type of the function to memoize.
type Func func(key string, done <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	response chan<- result   // the client wants a single result
	done     <-chan struct{} // closed when res is done
}

type Memo struct{ requests, cancels chan request }

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request), cancels: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	req := request{key, response, done}
	memo.requests <- req
	fmt.Println("get: waiting for response")
	res := <-response
	fmt.Println("get: checking if cancelled")
	select {
	case <-done:
		memo.cancels <- req
	default:
		// Do nothing.
	}
	fmt.Println("get: return")
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	// for req := range memo.requests {
	//     e := cache[req.key]
	//     if e == nil {
	//         // This is the first request for this key.
	//         e = &entry{ready: make(chan struct{})}
	//         cache[req.key] = e
	//         go e.call(f, req.key) // call f(key)
	//     }
	//     go e.deliver(req.response)
	// }

Loop:
	for {
	Cancel:
		// Process all cancellations before requests.
		// After Get has returned a cancellation for some key, any subsequent
		// requests for that key should return the result of a new call to
		// Func. If select is allowed to choose randomly between processing
		// requests and cancellations it can't be predicted whether a request
		// will be cancelled by a previous cancellation or not without looking
		// at the cancels queue, which client obviously can't do.
		for {
			select {
			case req := <-memo.cancels:
				fmt.Println("server: deleting cancelled entry (early)")
				delete(cache, req.key)
			default:
				break Cancel
			}
		}
		select {
		case req := <-memo.requests:
			fmt.Println("server: deleting cancelled entry")
			e := cache[req.key]
			if e == nil {
				// This is the first request for this key.
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done) // call f(key)
			}
			go e.deliver(req.response)
		case req := <-memo.cancels:
			delete(cache, req.key)
			continue Loop

		}
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key, done)
	fmt.Println("call: returned from f")
	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Wait for the ready condition.
	<-e.ready
	// Send the result to the client.
	response <- e.res
}
