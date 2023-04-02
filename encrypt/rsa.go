package encrypt

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "github.com/d3code/pkg/log"
)

func RsaGenerate() (*rsa.PrivateKey, *rsa.PublicKey) {
    privateKey, _ := rsa.GenerateKey(rand.Reader, 4096)
    return privateKey, &privateKey.PublicKey
}

func RsaPrivateToString(privateKey *rsa.PrivateKey) string {
    privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
    encodeToMemory := pem.EncodeToMemory(
        &pem.Block{
            Type:  "RSA PRIVATE KEY",
            Bytes: privateKeyBytes,
        },
    )
    return string(encodeToMemory)
}

func RsaPublicToString(publicKey *rsa.PublicKey) string {
    pkixPublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        log.Log.Error(err)
        return ""
    }

    encodedString := pem.EncodeToMemory(
        &pem.Block{
            Type:  "RSA PUBLIC KEY",
            Bytes: pkixPublicKey,
        },
    )

    return string(encodedString)
}

func RsaPrivateFromString(privateKeyPem string) *rsa.PrivateKey {
    block, decodeError := pem.Decode([]byte(privateKeyPem))
    if block == nil {
        log.Log.Error(decodeError)
        return nil
    }

    privateKey, parseError := x509.ParsePKCS1PrivateKey(block.Bytes)
    if parseError != nil {
        log.Log.Error(parseError)
        return nil
    }

    return privateKey
}

func RsaPublicFromString(publicKey string) *rsa.PublicKey {
    block, decodeError := pem.Decode([]byte(publicKey))
    if block == nil {
        log.Log.Error(decodeError)
        return nil
    }

    pkixPublicKey, parseError := x509.ParsePKIXPublicKey(block.Bytes)
    if parseError != nil {
        log.Log.Error(parseError)
        return nil
    }

    switch pub := pkixPublicKey.(type) {
    case *rsa.PublicKey:
        return pub
    default:
        break
    }

    log.Log.Error("Key type is not RSA")
    return nil
}
