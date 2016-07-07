BINARY=sixalert
WEBDISTDIR=web/dist
JSFILES=$(shell find web/js -name '*.js')
CSSFILES=$(shell find web/css -name '*.*css')
GLYPHFILES=$(shell find web/dist -name 'glyphicons-*.*')
BUNDLEFILE=bundle.js
BINDATAFILE=server/bindata.go
BINDATAFLAGS=-pkg server -ignore .gitignore -o ${BINDATAFILE} ${WEBDISTDIR}

.DEFAULT_GOAL: $(BINARY)

SOURCEDIR=.
SOURCES := $(shell find $(SOURCEDIR) \( -name '*.go' -and -not -name 'bindata.go' \))

$(BINARY): $(SOURCES) $(SOURCEDIR)/server/bindata.go
	go build

$(WEBDISTDIR):
	mkdir -pv ${WEBDISTDIR}

$(WEBDISTDIR)/index.html: web/index.html
	cp web/index.html ${WEBDISTDIR}/index.html

$(WEBDISTDIR)/$(BUNDLEFILE): $(JSFILES) $(CSSFILES) $(GLYPHFILES)
	cd web && npm run build

$(SOURCEDIR)/$(BINDATAFILE): $(WEBDISTDIR)/index.html $(WEBDISTDIR)/$(BUNDLEFILE)
	go-bindata ${BINDATAFLAGS}

.PHONY: clean
clean:
	@rm -f ${SOURCEDIR}/server/bindata.go
	@rm -f ${BINARY}
	@rm -f ${WEBDIST}/*.html ${WEBDIST}/glyphicons-*.* ${WEBDIST}/*.js
