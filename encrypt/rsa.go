package encrypt

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "fmt"
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
        fmt.Println(err)
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
        fmt.Println(decodeError)
        return nil
    }

    privateKey, parseError := x509.ParsePKCS1PrivateKey(block.Bytes)
    if parseError != nil {
        fmt.Println(parseError)
        return nil
    }

    return privateKey
}

func RsaPublicFromString(publicKey string) *rsa.PublicKey {
    block, decodeError := pem.Decode([]byte(publicKey))
    if block == nil {
        fmt.Println(decodeError)
        return nil
    }

    pkixPublicKey, parseError := x509.ParsePKIXPublicKey(block.Bytes)
    if parseError != nil {
        fmt.Println(parseError)
        return nil
    }

    switch pub := pkixPublicKey.(type) {
    case *rsa.PublicKey:
        return pub
    default:
        break
    }

    fmt.Println("Key type is not RSA")
    return nil
}
