# Learning day Febr 2023

# Web Assembly

## Links

Some links for getting familiar with WASM / WASI etc.

- https://www.redhat.com/en/blog/red-hat-and-webassembly
- https://opensource.com/article/22/10/wasm-containers
- https://wasmedge.org/book/en/index.html
- https://webassembly.org/getting-started/developers-guide/
- https://developer.mozilla.org/en-US/docs/WebAssembly
- https://github.com/appcypher/awesome-wasm-langs#awesome-webassembly-languages-

- https://github.com/redhat-et/wasm-demo-app
- https://www.youtube.com/watch?v=3fudsMOkRCM

## Experiment

- Create a Rust library
- Compile to WASM
- Call from Go

based on https://wasmedge.org/book/en/sdk/go/function.html

### Rust

some commands running along the way... not sure if complete / all are needed

```
# Setup env
$ rustup target add wasm32-was
$ cargo install wasm-pack

# Setup project
$ cargo add wasmedge_bindgen
$ cargo add wasmedge_bindgen_macro

# Build
$ cargo build --target wasm32-wasi --release
```

### Go

```
$ go run main.go
called WASM -- say: hello from Rust: called from Go
```
