.SILENT:

proto-gen:
	echo "##############################################"
	echo "# MAKE SURE THAT YOU STAGE ALL PROTO CHANGES #"
	echo "##############################################"
	rm -r pkg
	mkdir pkg
	for file in $$(git ls-files '*.proto'); do \
		protoc \
			--go_out=pkg/ \
			--go_opt=paths=source_relative \
			--go-grpc_out=pkg \
			--go-grpc_opt=paths=source_relative \
			$$file; \
	done
