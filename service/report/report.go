package report

import (
	"context"
	"strings"

	"github.com/chaitin/libveinmind/go/cmd"
	libService "github.com/chaitin/libveinmind/go/plugin/service"

	"github.com/chaitin/veinmind-common-go/service/report/service"
)

type Service = service.Service

func NewService(ctx context.Context, opts ...service.Option) *service.Service {
	return service.NewService(ctx, opts...)
}

// MapReportCmd for plugins to simple report
func MapReportCmd(c *cmd.Command, s *service.Service, opts ...service.Option) *cmd.Command {
	c.Flags().StringP("format", "f", "cli", "cli/json/html, support multiple with `,` split")
	c.Flags().BoolP("verbose", "v", false, "show detail info at cli")
	// only standalone mode needs PostRun to gather output
	if !libService.Hosted() {
		c.PostRun = func(c *cmd.Command, args []string) {
			s.Write()
		}
	}
	// all mode need init a report service
	c.PreRun = func(c *cmd.Command, args []string) {
		*s = *service.NewService(c.Context())
		// verbose
		v, err := c.Flags().GetBool("verbose")
		if err == nil && v {
			opts = append(opts, service.WithVerbose())
		}
		// format
		format, err := c.Flags().GetString("format")
		if err != nil {
			format = "cli"
		}

		formatList := strings.Split(format, ",")
		for _, f := range formatList {
			switch f {
			case "cli":
				opts = append(opts, service.WithTableRender())
			case "json":
				opts = append(opts, service.WithJsonRender())
			case "html":
				opts = append(opts, service.WithHtmlRender())
			}
		}
		// user options
		for _, o := range opts {
			o(s.Options)
		}
	}
	return c
}
