module ricochet/aurora/docker

go 1.20

require (
	github.com/docker/docker v24.0.5+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/google/go-cmp v0.5.9
	ricochet/aurora/types v0.0.0-00010101000000-000000000000
)

require (
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/docker/distribution v2.8.2+incompatible // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	gotest.tools/v3 v3.5.0 // indirect
)

replace ricochet/aurora/types => ../types
