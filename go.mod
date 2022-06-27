module github.com/chaitin/veinmind-common-go

go 1.16

require (
	github.com/chaitin/libveinmind v1.1.1
	github.com/containerd/containerd v1.6.6
	github.com/docker/cli v20.10.17+incompatible
	github.com/docker/distribution v2.8.1+incompatible
	github.com/docker/docker v20.10.17+incompatible
	github.com/fvbommel/sortorder v1.0.2 // indirect
	github.com/google/go-containerregistry v0.10.0
	github.com/moby/sys/mount v0.3.3 // indirect
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6 // indirect
	github.com/pelletier/go-toml v1.9.4
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/theupdateframework/notary v0.7.0 // indirect
	golang.org/x/sync v0.0.0-20220513210516-0976fa681c29
	google.golang.org/grpc/naming v0.0.0-00010101000000-000000000000 // indirect
)

replace google.golang.org/grpc/naming => github.com/xiegeo/grpc-naming v1.29.1-alpha
