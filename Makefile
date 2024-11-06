generate:
	go generate ./...
	
server:
	go run ./cmds/server/

client:
	go run ./cmds/client/