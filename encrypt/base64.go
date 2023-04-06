package encrypt

import "encoding/base64"

func Base64Encode(message []byte) []byte {
    bytes := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
    base64.StdEncoding.Encode(bytes, message)
    return bytes
}

func Base64Decode(message []byte) ([]byte, error) {
    bytes := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
    length, err := base64.StdEncoding.Decode(bytes, message)
    if err != nil {
        return bytes, err
    }
    return bytes[:length], nil
}
