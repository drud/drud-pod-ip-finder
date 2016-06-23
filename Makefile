#make this shiz

all: push

# 0.0 shouldn't clobber any released builds
TAG = 1.0
PREFIX = drud/drud-pod-ip-finder

main: main.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o main ./main.go

container: main
	docker build -t $(PREFIX):$(TAG) .

push: container
	docker push $(PREFIX):$(TAG)

clean:
	rm -f main
