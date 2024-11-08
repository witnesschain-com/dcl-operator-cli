witness: cmd/witness/main.go commands/watchtower_status.go
	go build -o witness -ldflags "-X main.VERSION="$$(git describe --tags --abbrev=0) cmd/witness/main.go 

build: cmd/witnesschain-cli/main.go 
	go build -o witnesschain-cli -ldflags "-X main.VERSION="$$(git describe --tags --abbrev=0) cmd/witnesschain-cli/main.go 
	go build -o witness -ldflags "-X main.VERSION="$$(git describe --tags --abbrev=0) cmd/witness/main.go 

test: build
	./witnesschain-cli registerProver --config-file templates/prover.json
	./witnesschain-cli deRegisterProver --config-file templates/prover.json
	./witnesschain-cli registerChallenger --config-file templates/challenger.json 
	./witnesschain-cli deRegisterChallenger --config-file templates/challenger.json

install: build
	cp witnesschain-cli ${HOME}/.local/bin

watchtowerStatus:
	go run cmd/witness/main.go  watchtowerStatus --watchtower-address 0x5549d62f67c2d8092602b499ac913030ecac50f5
	go run cmd/witness/main.go  watchtowerStatus --watchtower-address 0x18dfdee484d4e1006e011c430755686f60386ab1

