BINARY=sixalert
ARM_BINARY=$(BINARY)-arm

.DEFAULT_GOAL: $(BINARY)

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) \( -name '*.go' \))

$(BINARY): $(SOURCES)
	go build

$(ARM_BINARY): $(SOURCES)
	GOOS=linux GOARCH=arm go build -o ${ARM_BINARY}

.PHONY: clean
clean:
	@rm -f ${BINARY}
	@rm -f ${ARM_BINARY}
