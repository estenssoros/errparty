package errparty

import (
	"sync"
)

// ErrParty real banger of a party
type ErrParty struct {
	errChan chan error
	wg      sync.WaitGroup
}

// RSVP add a partier to the party
func (p *ErrParty) RSVP(partier func() error) {
	p.wg.Add(1)
	go func(ch chan error, w *sync.WaitGroup) {
		defer w.Done()
		if err := partier(); err != nil {
			p.errChan <- err
		}
	}(p.errChan, &p.wg)
}

// Party throw a party
func (p *ErrParty) Party() error {
	go func() {
		p.wg.Wait()
		close(p.errChan)
	}()
	for err := range p.errChan {
		return err
	}
	return nil
}
