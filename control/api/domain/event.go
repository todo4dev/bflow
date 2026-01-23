// domain/event.go
package domain

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/leandroluk/gox/util"
)

type Event struct {
	Kind           string    `json:"kind"`
	IdempotencyKey string    `json:"idempotency_key"`
	OccurredAt     time.Time `json:"occurred_at"`
	Payload        any       `json:"payload"`
}

func NewEvent(kind string, payload any, optionalIdempotencyKey ...string) Event {
	event := Event{
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

type EventMapper struct {
	decoders map[string]DecoderFunc
}

func NewEventMapper() *EventMapper {
	return &EventMapper{decoders: map[string]DecoderFunc{}}
}

func (m *EventMapper) Decoder(kind string, decoder DecoderFunc) *EventMapper {
	m.decoders[kind] = decoder
	return m
}

func (m *EventMapper) Decode(e Event) (any, error) {
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
