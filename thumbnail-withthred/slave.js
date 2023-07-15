importScripts('../wasm_exec.js');
const go = new Go();

WebAssembly.instantiateStreaming(
    fetch("./main.wasm"),
    go.importObject
).then((result) => {
    go.run(result.instance);
});


onmessage = function(e) {
    console.log(e.data.name);
    this.self.cropImage(e.data.image);
}

function createImageElement(cropImagebase64,originalImageBase64) {
    postMessage({type: 'createImageElement', from:'slave', cropImagebase64,originalImageBase64});
}