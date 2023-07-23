set GOOS=js
set GOARCh=wasm
go build -o D:\webasm\assets\main.wasm
copy C:\Program Files\Go\misc\wasm\wasm_exec.js D:\webasm\assets\wasm_exec.js