.SILENT:

proto-gen:
	mkdir -p pkg
	for file in $$(git ls-files '*.proto'); do \
		# mkdir -p $$(dirname $$file); \
		protoc \
			--go_out=pkg/ \
			--go_opt=paths=source_relative \
			--go-grpc_out=pkg \
			--go-grpc_opt=paths=source_relative \
			$$file; \
	done