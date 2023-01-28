package report

import (
	"context"
	"github.com/chaitin/libveinmind/go/plugin/service"
	"golang.org/x/sync/errgroup"
)

const Namespace = "github.com/chaitin/veinmind-tools/veinmind-common/go/service/report"
const BufferSize = 1 << 8

type Service struct {
	EventChannel chan Event
}

type reportClient struct {
	ctx    context.Context
	group  *errgroup.Group
	Report func(Event) error
}

func (s *Service) Report(evt Event) {
	s.EventChannel <- evt
}

func (s *Service) Add(registry *service.Registry) {
	registry.Define(Namespace, struct{}{})
	registry.AddService(Namespace, "report", s.Report)
}

func NewReportService() *Service {
	return &Service{
		EventChannel: make(chan Event, BufferSize),
	}
}
