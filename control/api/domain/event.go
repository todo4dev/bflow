// domain/event.go
package domain

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/leandroluk/gox/util"
)

type Event[TKind ~string] struct {
	Kind           TKind     `json:"kind"`
	IdempotencyKey string    `json:"idempotency_key"`
	OccurredAt     time.Time `json:"occurred_at"`
	Payload        any       `json:"payload"`
}

func NewEvent[TKind ~string](kind TKind, payload any, optionalIdempotencyKey ...string) Event[TKind] {
	event := Event[TKind]{
		Kind:       kind,
		OccurredAt: time.Now(),
		Payload:    payload,
	}
	if len(optionalIdempotencyKey) > 0 {
		event.IdempotencyKey = optionalIdempotencyKey[0]
	} else {
		event.IdempotencyKey = util.Must(uuid.NewV7()).String()
	}
	return event
}

type DecoderFunc func(payload any) (any, error)

type EventMapper[TKind ~string] struct {
	decoders map[TKind]DecoderFunc
}

func NewEventMapper[TKind ~string]() *EventMapper[TKind] {
	return &EventMapper[TKind]{decoders: map[TKind]DecoderFunc{}}
}

func (m *EventMapper[TKind]) Decoder(kind TKind, decoder DecoderFunc) *EventMapper[TKind] {
	m.decoders[kind] = decoder
	return m
}

func (m *EventMapper[TKind]) Decode(e Event[TKind]) (any, error) {
	decoder, ok := m.decoders[e.Kind]
	if !ok {
		return nil, errors.New("unmapped event kind")
	}
	return decoder(e.Payload)
}

func JSONDecoder[TPayload any]() DecoderFunc {
	return func(payload any) (any, error) {
		if typed, ok := payload.(TPayload); ok {
			return typed, nil
		}
		if typedPtr, ok := payload.(*TPayload); ok && typedPtr != nil {
			return *typedPtr, nil
		}

		b, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		var out TPayload
		if err := json.Unmarshal(b, &out); err != nil {
			return nil, err
		}

		return out, nil
	}
}

