build:
	rm -rf messages/*
	go build -o server message.go
	touch messages/placeholder
test:
	rm -rf messages/*
	go test --cover
	touch messages/placeholder
clean:
	rm -rf messages/*
	rm -rf server
	touch messages/placeholder