cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./assets/wasm_exec.js
GOOS=js GOARCH=wasm go build -o ./assets/main.wasm ./cmd/wasm/main.go