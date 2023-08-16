use wasm_bindgen::prelude::*;
use qrcode::QrCode;
use base64;

#[wasm_bindgen]
pub fn generate_qr_code(input: &str) -> String {
    let code = QrCode::new(input).unwrap();
    let image = code.render::<u8>().build();
    base64::encode(&image)
}

