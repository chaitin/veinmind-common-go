module github.com/chaitin/veinmind-common-go

go 1.16

require (
	github.com/chaitin/libveinmind v1.5.2
	github.com/containerd/containerd v1.6.6
	github.com/docker/cli v20.10.20+incompatible
	github.com/docker/distribution v2.8.1+incompatible
	github.com/docker/docker v20.10.20+incompatible
	github.com/fvbommel/sortorder v1.0.2 // indirect
	github.com/google/go-containerregistry v0.12.1
	github.com/moby/sys/mount v0.3.3 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/pelletier/go-toml v1.9.4
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/theupdateframework/notary v0.7.0 // indirect
	golang.org/x/sync v0.1.0
)

replace google.golang.org/grpc/naming => github.com/xiegeo/grpc-naming v1.29.1-alpha
