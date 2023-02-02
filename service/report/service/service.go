package service

import (
	"context"
	"errors"
	"sync"

	"github.com/chaitin/libveinmind/go/plugin/service"

	"github.com/chaitin/veinmind-common-go/service/report/event"
)

const Namespace = "github.com/chaitin/veinmind-tools/veinmind-common/go/service/report"
const BufferSize = 1 << 8

// Service means use a channel to accept event,
// and you can use default listen or custom listen function to deal with those events
// when service channel close, it will check format type and try output a summary report.
type Service struct {
	ctx         context.Context
	cancel      context.CancelFunc
	eventsMutex sync.RWMutex
	closeOnce   sync.Once

	Options   *Options
	Client    *Client
	EventPool *EventPool
}

type Client struct {
	ctx    context.Context
	Report func(event *event.Event) error
}

type EventPool struct {
	Events       []*event.Event
	EventChannel chan *event.Event
}

func NewService(ctx context.Context, opts ...Option) *Service {
	var s *Service

	ctx, cancel := context.WithCancel(ctx)

	s = &Service{
		ctx:    ctx,
		cancel: cancel,
		Options: &Options{
			verbose:      false,
			formatEnable: make(map[Format]struct{}, 0),
		},
		EventPool: &EventPool{
			Events:       make([]*event.Event, 0),
			EventChannel: make(chan *event.Event, BufferSize),
		},
	}

	for _, o := range opts {
		o(s.Options)
	}

	// register client
	hasService := false
	if service.Hosted() {
		ok, err := service.HasNamespace(Namespace)
		if err != nil {
			panic(err)
		}
		hasService = ok
	}
	if hasService {
		var r func(*event.Event) error
		service.GetService(Namespace, "report", &r)
		s.Client = &Client{
			ctx:    ctx,
			Report: r,
		}
	} else {
		s.Client = &Client{
			ctx:    ctx,
			Report: s.AppendEvent,
		}
	}
	return s
}

func (s *Service) AppendEvent(event *event.Event) error {
	s.eventsMutex.Lock()
	s.EventPool.Events = append(s.EventPool.Events, event)
	s.eventsMutex.Unlock()
	return nil
}

func (s *Service) Report(event *event.Event) error {
	// use for hosted
	s.EventPool.EventChannel <- event
	return nil
}

// Listen provide a default function to deal with event in channel
func (s *Service) Listen() {
	for {
		select {
		case evt := <-s.EventPool.EventChannel:
			s.AppendEvent(evt)
		case <-s.ctx.Done():
			return
		}
	}
}

// Close provide a default function to gather event and output
func (s *Service) Close() {
	s.closeOnce.Do(func() {
		s.cancel()
		close(s.EventPool.EventChannel)

		// sync previous sent event
		s.eventsMutex.Lock()
		for e := range s.EventPool.EventChannel {
			s.EventPool.Events = append(s.EventPool.Events, e)
		}
		s.eventsMutex.Unlock()

		// render output
		s.Write()
	})
}

// Add can register report Service
func (s *Service) Add(registry *service.Registry) {
	registry.Define(Namespace, struct{}{})
	registry.AddService(Namespace, "report", s.Report)
}

// Write generate output
func (s *Service) Write() error {
	// check render
	for key, _ := range s.Options.formatEnable {
		switch key {
		case CLI:
			return WriteTable(s)
		case Json:
			return WriteJson(s)
		case Html:
			return WriteHtml(s)
		}
	}
	return errors.New("no format scheme")
}
