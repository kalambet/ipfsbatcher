module ipfsbatcher

go 1.15

require (
	code.cloudfoundry.org/bytefmt v0.0.0-20200131002437-cf55d5288a48
	github.com/filecoin-project/go-commp-utils v0.0.0-20201119054358-b88f7a96a434
	github.com/filecoin-project/lotus v1.2.2
	github.com/google/uuid v1.1.2
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-ipfs-files v0.0.8
	github.com/ipfs/go-ipfs-http-client v0.1.0
	github.com/ipfs/go-unixfs v0.2.4
	github.com/ipfs/interface-go-ipfs-core v0.4.0
	github.com/ipld/go-car v0.1.1-0.20201119040415-11b6074b6d4d
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/textileio/powergate v1.2.4 // indirect
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

replace github.com/filecoin-project/lotus => ./extern/lotus
