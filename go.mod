module github.com/photon-storage/bor

// Note: Change the go image version in Dockerfile if you change this.
go 1.26.5

require (
	github.com/Microsoft/go-winio v0.6.2
	github.com/consensys/gnark-crypto v0.19.2
	github.com/crate-crypto/go-eth-kzg v1.4.0
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
	github.com/deckarep/golang-set/v2 v2.6.0
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0
	github.com/ethereum/c-kzg-4844/v2 v2.1.5
	github.com/ethereum/go-verkle v0.2.2
	github.com/google/gofuzz v1.2.0
	github.com/gorilla/websocket v1.5.3
	github.com/holiman/uint256 v1.3.2
	github.com/influxdata/influxdb-client-go/v2 v2.13.0
	github.com/influxdata/influxdb1-client v0.0.0-20220302092344-a9ab5670611c
	github.com/jedisct1/go-minisign v0.0.0-20230811132847-661be99b8267
	github.com/kylelemons/godebug v1.1.0
	github.com/shirou/gopsutil v3.21.11+incompatible
	github.com/stretchr/testify v1.11.1
	golang.org/x/crypto v0.53.0
	golang.org/x/sys v0.46.0
)

require (
	github.com/ProjectZKM/Ziren/crates/go-runtime/zkvm_runtime v0.0.0-20260104020744-7268a54d0358
	github.com/bits-and-blooms/bitset v1.24.4 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20210922203350-b1ad95c89adf // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/tklauser/go-sysconf v0.3.16 // indirect
	github.com/tklauser/numcpus v0.11.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	gotest.tools v2.2.0+incompatible
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/crate-crypto/go-ipa v0.0.0-20240724233137-53bbb0ceb27a // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/oapi-codegen/runtime v1.1.0 // indirect
	github.com/supranational/blst v0.3.16 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	golang.org/x/sync v0.21.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	cosmossdk.io/client/v2 => github.com/0xPolygon/cosmos-sdk/client/v2 v2.0.0-beta.6 // Same as heimdall-v2.
	github.com/Masterminds/goutils => github.com/Masterminds/goutils v1.1.1
	github.com/cometbft/cometbft => github.com/0xPolygon/cometbft v0.3.8-polygon
	github.com/cometbft/cometbft-db => github.com/0xPolygon/cometbft-db v0.14.1-polygon
	github.com/cosmos/cosmos-sdk => github.com/0xPolygon/cosmos-sdk v0.2.13-polygon
	github.com/ethereum/go-ethereum => github.com/0xPolygon/bor v1.14.14-0.20260219070410-6b0405c0a5ca
	github.com/ethereum/go-ethereum/common/math => github.com/0xPolygon/bor/common/math v1.5.5
	go.mongodb.org/mongo-driver => go.mongodb.org/mongo-driver v1.14.0
)
