# sweep-playground

## Compiling Go code to WASM

To compile the Go code to WASM, you need to have Go installed on your machine. You can then use the following command:

```bash
GOOS=js GOARCH=wasm go build -o main.wasm
```

This will generate a WASM file named `main.wasm`. Please ensure that the `main.wasm` file is included in the repository after it is generated.

## Using the WASM file in HTML

To use the WASM file in HTML, you need to load it using the WebAssembly JavaScript API. You can then call the exported Go function and display the returned QR code.

Before using the example, please ensure that the `main.wasm` file exists in the repository. Here is an example:

```html
<script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
    });

    function generateQR() {
        const input = document.getElementById('inputField').value;
        go.exports.GenerateQRCode(input);
    }
</script>