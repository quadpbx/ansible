SHELL=/bin/bash

GOVERS=1.22.11
PKG=go$(GOVERS).linux-amd64.tar.gz
URL=https://go.dev/dl/$(PKG)
DEST=/usr/local/$(PKG)
GOBIN=/usr/local/go/bin

.PHONY: gocheck
gocheck:
	@if [ ! -e $(GOBIN)/go ]; then echo "go not installed. Run 'make go' and then restart this session"; exit 99; fi

.PHONY: go
go: $(GOBIN)/go
	@echo If this errors, restart your login session to update PATH
	@echo "  The version reported below should be "$(GOVERS)". If it is not,"
	@echo "  run 'make update-go'"
	@go version

.PHONY: update-go
update-go $(GOBIN)/go: $(DEST) | bashrc
	rm -rf /usr/local/go
	tar -C /usr/local -zxf $(DEST)
	touch -r $(BIN)/go $(DEST)

$(DEST):
	wget $(URL) -O $@

.PHONY: bashrc
bashrc:
	@grep -q $(GOBIN) /etc/bash.bashrc || echo 'PATH=$$PATH:$(GOBIN)' >> /etc/bash.bashrc




