MAKEFLAGS += --no-builtin-rules
.SUFFIXES:

BINARY=sixalert
ARM_BINARY=$(BINARY)-arm

.DEFAULT: $(BINARY)

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) -type f \( -path vendor \) -prune -o -name '*.go')

$(BINARY): $(SOURCES)
	@go build

$(ARM_BINARY): $(SOURCES)
	GOOS=linux GOARCH=arm @go build -o ${ARM_BINARY}

.PHONY: clean
clean:
	@rm -f ${BINARY}
	@rm -f ${ARM_BINARY}
	@go clean -i

