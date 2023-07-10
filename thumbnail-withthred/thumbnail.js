importScripts('../wasm_exec.js');
const go = new Go();
WebAssembly.instantiateStreaming(
    fetch("./main.wasm"),
    go.importObject
).then((result) => {
    go.run(result.instance);
});


onmessage = function(e) {
    this.self.cropImage(e.data);
}

function createImageElement(cropImagebase64,originalImageBase64) {
    postMessage({type: 'createImageElement', cropImagebase64,originalImageBase64});
}


