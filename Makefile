generate: generate_go generate_rust

generate_go:
	cd go && go generate && go build

generate_rust:
	cd rust && cargo build

clean:
	cd rust && cargo clean
	-rm rust/src/lib.rs
	touch rust/src/lib.rs

