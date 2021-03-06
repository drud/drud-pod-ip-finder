#make this shiz

all: push

TAG = latest
PREFIX = drud/drud-pod-ip-finder

main: main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o main ./main.go

container: main
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

clean:
	rm -f main
