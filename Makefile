build:
	rm -rf messages/*
	go build -o server message.go
test:
	rm -rf messages/*
	go test --cover
clean:
	rm -rf messages/*