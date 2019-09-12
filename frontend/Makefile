DISTDIR ?= dist

ALL_TARGETS := \
	$(DISTDIR)/app.js \
	$(DISTDIR)/app.css \
	assets_vfsdata.go

all: $(ALL_TARGETS)

clean:
	-rm -f assets_vfsdata.go

du:
	@du -h $(ALL_TARGETS)

assets_vfsdata.go:
	cd .. && vfsgendev  -source=\"github.com/tonyjia87/editor/frontend\".Assets && mv $@ frontend/

.PHONY: all clean du assets_vfsdata.go

