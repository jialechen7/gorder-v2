.PHONY: gen
gen: genproto genopenapi

.PHONY: genproto
genproto:
	@./scripts/genproto.sh

.PHONY: genopenapi
genopenapi:
	@./scripts/genopenapi.sh

.PHONY: fmt
fmt:
	goimports -l -w internal/

.PHONY: lint
lint:
	@./scripts/lint.sh

.PHONY: listen
listen:
	stripe listen --forward-to localhost:8282/api/webhook