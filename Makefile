build:
	go build -o simpd ./cmd/simpd 
	go build -o simp ./cmd/simp

install:
	go install ./cmd/simpd -o simpd
	go install ./cmd/simp -o simp

run-server: build
	rm -rf /tmp/simpd.sock
	./simpd