const go = new Go();

let mod, inst;
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
    async result =>{
        mod = result.module;
        inst = result.instance;
        await go.run(inst)
    }
)

function CalculClick(){
    let value = document.getElementById("inp").value
    let result = calculatePolish(value)
    document.getElementById("answerInput").value = result
}