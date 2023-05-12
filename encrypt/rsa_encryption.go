package encrypt

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha512"
    "github.com/d3code/pkg/zlog"
    "os"
)

func RsaEncrypt(toEncrypt string, privateKeyPath string) *string {
    if m, err := os.Stat(privateKeyPath); err != nil || m.IsDir() {

        zlog.Log.Error("Could not find private key to encrypt")
        return nil

    } else {

        privateKey, _ := os.ReadFile(privateKeyPath)
        pem := RsaPrivateFromString(string(privateKey))
        if pem == nil {
            return nil
        }
        encrypted := EncryptWithPublicKey(toEncrypt, &pem.PublicKey)
        return &encrypted
    }
}

func RsaDecrypt(toDecrypt string, privateKeyPath string) *string {
    if m, err := os.Stat(privateKeyPath); err != nil || m.IsDir() {

        zlog.Log.Error("Could not find private key to decrypt")
        return nil

    } else {

        privateKey, _ := os.ReadFile(privateKeyPath)
        pem := RsaPrivateFromString(string(privateKey))
        if pem == nil {
            return nil
        }
        dec := DecryptWithPrivateKey(toDecrypt, pem)
        return &dec
    }
}

// EncryptWithPublicKey encrypts data with public key
func EncryptWithPublicKey(message string, publicKey *rsa.PublicKey) string {
    hash := sha512.New()
    ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, []byte(message), nil)
    if err != nil {
        zlog.Log.Error(err)
    }
    base64Bytes := Base64Encode(ciphertext)
    return string(base64Bytes)
}

// DecryptWithPrivateKey decrypts data with private key
func DecryptWithPrivateKey(ciphertext string, privateKey *rsa.PrivateKey) string {
    hash := sha512.New()
    base64Decode, _ := Base64Decode([]byte(ciphertext))
    plaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, base64Decode, nil)
    if err != nil {
        zlog.Log.Error(err)
    }
    return string(plaintext)
}
