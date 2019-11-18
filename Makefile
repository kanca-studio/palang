all: compile start

DIRS         = $(shell cd cmd && ls -d */ | grep -v "_output")
ODIR        := bin

export VAR_SERVICES       ?= $(DIRS:/=)

$(ODIR):
	@mkdir -p $(ODIR)

compile: $(ODIR)
	@$(foreach svc, $(VAR_SERVICES), \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ODIR)/$(svc) cmd/$(svc)/main.go;)

test:
	go test -v -coverprofile=cover.out ./...

start:
	export $$(cat .env | grep -v ^\# | xargs) && \
	./bin/api