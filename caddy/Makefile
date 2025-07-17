SHELL=/bin/bash

GOVERS=1.22.11
PKG=go$(GOVERS).linux-amd64.tar.gz
XVER=0.3.5
XPKG=xcaddy_$(XVER)_linux_amd64.deb
URL=https://go.dev/dl/$(PKG)
DEST=/usr/local/$(PKG)
BIN=/usr/local/go/bin
SVC=caddy.service
SSVC=/etc/systemd/system/$(SVC)
CADDYUID=901
CADDYGID=901
CADDYPKGS=github.com/caddy-dns/cloudflare
CADDYBIN=/usr/local/bin/caddy

.PHONY: help
help:
	@if [ ! -e $(BIN)/go ]; then echo "go not installed. Run 'make go' and then restart this session"; else \
		echo 'Run "build" or "force-build" to build /usr/local/bin/caddy'; \
	fi

.PHONY: caddy
caddy: $(CADDYBIN) $(SSVC) /usr/local/etc/Caddyfile
xxcaddy: $(CADDYBIN) /usr/local/etc/Caddyfile /usr/local/etc/caddy_template $(SSVC) | xcaddy

.PHONY: check
check: $(CADDYBIN)
	$< validate

.PHONY: fmt
fmt: $(CADDYBIN)
	$< fmt --overwrite

.PHONY: force-build build
build: $(CADDYBIN)


.cloudflare_ips:
	curl -s https://api.cloudflare.com/client/v4/ips | jq -r '.result | .ipv4_cidrs+.ipv6_cidrs | join(" ")' > $@

force-build $(CADDYBIN): /usr/bin/xcaddy
	/usr/bin/xcaddy build $(addprefix --with ,$(CADDYPKGS)) --output $(CADDYBIN)

/var/lib/caddy /var/log/caddy: | users
	@mkdir -p $@
	@chown -R caddy:caddy $@

.PHONY: users
users:
	@grep -q '^caddy:' /etc/group || groupadd -g $(CADDYGID) caddy
	@grep -q '^caddy:' /etc/passwd || useradd --gid $(CADDYGID) --uid $(CADDYUID) --no-create-home --no-user-group --password 'x' --home-dir /var/lib/caddy caddy

build.sh: | xcaddy

/usr/local/etc/Caddyfile: Caddyfile .cloudflare_ips $(wildcard components/*)
	@sed "s!__CLOUDFLARE__!$$(cat .cloudflare_ips)!" < Caddyfile > /tmp/.caddyfile
	@for x in $$(awk '/^#component/ { print $$2}' < /tmp/.caddyfile); do echo Importing $$x; sed -i -e "/^#component $$x/r components/$$x" /tmp/.caddyfile; done
	@$(CADDYBIN) validate -c /tmp/.caddyfile
	@cp /tmp/.caddyfile $@

/usr/local/etc/caddy_template: caddy_template
	cp $< $@

$(SSVC): $(SVC) /usr/local/etc/Caddyfile
	echo 'cp $(SVC) < $@'
#	systemctl enable $(SVC)
#	systemctl restart $(SVC)

.PHONY: xcaddy
xcaddy: /usr/bin/xcaddy

/usr/bin/xcaddy: | $(BIN)/go $(XPKG)
	dpkg -i $(XPKG)

.PHONY: go
go: /var/lib/caddy /var/log/caddy $(BIN)/go
	@echo If this errors, restart your login session to update PATH
	@echo "  The version reported below should be "$(GOVERS)". If it is not,"
	@echo "  run 'make update-go'"
	@go version

clean:
	echo rm -rf $(shell dirname $(BIN))


.PHONY: update-go
update-go $(BIN)/go: $(DEST)
	rm -rf /usr/local/go
	tar -C /usr/local -zxf $(DEST)
	touch -r $(BIN)/go $(DEST)

$(DEST): | bashrc
	wget $(URL) -O $@

$(XPKG):
	wget https://github.com/caddyserver/xcaddy/releases/download/v$(XVER)/xcaddy_$(XVER)_linux_amd64.deb -O $(XPKG)

.PHONY: bashrc
bashrc:
	@grep -q $(BIN) /etc/bash.bashrc || echo 'PATH=$$PATH:$(BIN)' >> /etc/bash.bashrc




