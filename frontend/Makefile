RESOURCESDIR ?= resources
BUILD ?= pro

ALL_TARGETS := \
	$(RESOURCESDIR)/js/app.js \
	$(RESOURCESDIR)/sass/app.scss \
	resources \
	assets_vfsdata.go

all: $(ALL_TARGETS)

clean:
	-rm -f assets_vfsdata.go

du:
	@du -h $(ALL_TARGETS)

resources:
	@npm run prod

assets_vfsdata.go:
	cd .. && vfsgendev  -source=\"github.com/tonyjia87/editor/frontend\".Assets && mv $@ frontend/

.PHONY: all clean du resources assets_vfsdata.go

