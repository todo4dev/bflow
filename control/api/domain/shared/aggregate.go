// domain/shared/aggregate.go
package shared

type Aggregate[TKind ~string] struct {
	events []Event[TKind]
}

func (a *Aggregate[TKind]) Publish(event Event[TKind]) {
	a.events = append(a.events, event)
}

func (a *Aggregate[TKind]) Pull() []Event[TKind] {
	if len(a.events) == 0 {
		return nil
	}
	events := a.events
	a.events = nil
	return events
}
