module github.com/axelarnetwork/axelar-core

go 1.16

require (
	github.com/armon/go-metrics v0.3.6 // indirect
	github.com/axelarnetwork/tm-events v0.0.0-20210427154304-86290b49ae8f
	// TODO: update to v0.22.0-beta once https://github.com/btcsuite/btcd/issues/1706 is resolved
	github.com/btcsuite/btcd v0.21.0-beta.0.20210506225145-0ec4bdc1b8e1
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/ethereum/go-ethereum v1.10.2
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/matryer/moq v0.2.1
	github.com/miguelmota/go-ethereum-hdwallet v0.0.1
	github.com/rakyll/statik v0.1.7
	github.com/rs/zerolog v1.21.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.9
	github.com/tendermint/tm-db v0.6.4
	google.golang.org/genproto v0.0.0-20210303154014-9728d6b83eeb
	google.golang.org/grpc v1.36.1
	google.golang.org/protobuf v1.26.0
)

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // https://github.com/axelarnetwork/axelar-core/issues/36
