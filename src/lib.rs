use wasm_bindgen::prelude::*;
use qrcode::QrCode;
use base64;

#[wasm_bindgen]
pub fn generate_qr_code(input: &str) -> String {
    let code = match QrCode::new(input) {
        Ok(code) => code,
        Err(_) => return "Error: Invalid input for QR code".to_string(),
    };
    let image = code.render::<u8>().build();
    base64::encode(&image)
}

