package main

import (
	"fmt"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
	bindgen "github.com/second-state/wasmedge-bindgen/host/go"
)

func main() {
	// Expected Args[1]: wasm or wasm-so file (rust_bindgen_funcs_lib.wasm))

	//wasmFile := os.Args[1:]
	wasmFile := []string{"../rust/target/wasm32-wasi/release/hello_lib.wasm"}

	// Set not to print debug info
	wasmedge.SetLogErrorLevel()

	// Create configure
	var conf = wasmedge.NewConfigure(wasmedge.WASI)

	// Create VM with configure
	var vm = wasmedge.NewVMWithConfig(conf)

	// Init WASI
	var wasi = vm.GetImportModule(wasmedge.WASI)
	wasi.InitWasi(
		wasmFile,        // The args
		os.Environ(),    // The envs
		[]string{".:."}, // The mapping preopens
	)

	vm.LoadWasmFile(wasmFile[0])
	vm.Validate()

	// Instantiate the bindgen and vm
	bg := bindgen.New(vm)
	bg.Instantiate()

	// Run bindgen functions
	res, _, err := bg.Execute("say", "called from Go")

	if err == nil {
		fmt.Println("called WASM -- say:", string(res[0].([]uint8)))
	} else {
		fmt.Println("called WASM -- say FAILED")
	}
}
