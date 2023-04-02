package encrypt

import (
    "strings"
    "testing"
)

func Test_GenerateRsaKeyPair(t *testing.T) {
    t.Run("Generate", func(t *testing.T) {
        privateKey, publicKey := RsaGenerate()
        if &privateKey.PublicKey != publicKey {
            t.Error("Public keys are not the same")
        }
    })
}

func Test_ExportRsaPrivateKeyAsPemStr(t *testing.T) {
    privateKey, _ := RsaGenerate()
    privateKeyAsPemStr := RsaPrivateToString(privateKey)

    t.Run("Header", func(t *testing.T) {
        if !strings.HasPrefix(privateKeyAsPemStr, "-----BEGIN RSA PRIVATE KEY-----") {
            t.Error("Pem block does not start with correct string")
        }
    })

    t.Run("Footer", func(t *testing.T) {
        if !strings.HasSuffix(privateKeyAsPemStr, "-----END RSA PRIVATE KEY-----\n") {
            t.Error("Pem block does not end with correct string")
        }
    })
}
