BINARY=sixalert

.DEFAULT_GOAL: $(BINARY)

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) \( -name '*.go' \))

$(BINARY): $(SOURCES)
	go build

.PHONY: clean
clean:
	@rm -f ${BINARY}
