.PHONY: models generate

# ==============================================================================
# Swagger Models
models:
	$(call print-target)
	find ./internal/api/ -type f -not -name '*_test.go' -delete
	swagger generate model -f docs/shop.yml -m internal/api

generate: models