package service

type Options struct {
	verbose      bool
	output       string
	formatEnable map[Format]struct{}
}

// Option provide some custom options with report
type Option func(o *Options)

// WithTableRender output table to cli console
func WithTableRender() Option {
	return func(o *Options) {
		o.formatEnable[CLI] = struct{}{}
	}
}

// WithJsonRender output origin event to json file
func WithJsonRender() Option {
	return func(o *Options) {
		o.formatEnable[Json] = struct{}{}
	}
}

// WithHtmlRender output table to html file
func WithHtmlRender() Option {
	return func(o *Options) {
		o.formatEnable[Html] = struct{}{}
	}
}

// WithOutputDir change output file dirs
func WithOutputDir(path string) Option {
	return func(o *Options) {
		o.output = path
	}
}

// WithVerbose means report will output info events
// like: image details, container details, cluster details
// also, application details
func WithVerbose() Option {
	return func(o *Options) {
		o.verbose = true
	}
}
