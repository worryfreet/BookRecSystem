package utils

import (
	"errors"
	"github.com/wenzhenxi/gorsa"
)

var publicKey = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAwx3OAU//0ATSlJp8mJMO
B/qFXJsc6fvdiUn+JlI/MfDtT0ju83ATeivy5oXvE8Rx9+lq32/tIX1h7mFQAmpX
EZpakhDL9cqgSLAYyJe696ibPtO6Mw46G7CEgMp6nIWTKAI5JcPjX+IuWvCZXjW0
P+iMlmIYDdEZcrtI+D/SuYYvBxFTZj6/0AoH2XeBTx7ZUPcY1Fr/h+RB3uUNDUEq
z7YXpCr/g9Np8KRCssxw4Sffy5AQxpJcEHk2z5SAzegp8q+CnYRfVftHcasTKhxB
gk41NkxlyTjhJZS9v75cNEsP5oILHLosOdxt+MAbJls4beMjyLKTNo5P8UyrqoX4
Lc1ykuAo7wfFbfgnGxbrilAUcN5NzJ92mYQGb6s9BnO9r6lsWvL/Hk+X2+uIPrFn
x0BeQmAaebmE6Y5sGkCklhi4+za73tmVJklNv2+nhnr8OO5D2tqCuqkk2cEGBw9Y
qFItbI5oOJvFfFVtuUKGPvgFnBCeXB6oSS8H19nQTtH7K5fxyXlKt3qFlG0LVvWV
Wpuv5csqKk8IK8xVaX2oFudctMDkiFbVl2fDcMm3PU8KQLANbGWJ2ll+0QYzrPC+
l1Dt19v3JSjNDKS9Wn9VLfv/1tIzizXbkPMESzg+Ta155msTBr23JEahU8cbwrAx
IumoBsCxfDZwkSUJwg0tA+cCAwEAAQ==
-----END PUBLIC KEY-----
`

var privateKey = `-----BEGIN PRIVATE KEY-----
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
RGR96NxQfSIkCHio8tAoT4JiReY2cA==
-----END PRIVATE KEY-----
`

// RsaPriDecode 私钥解密
func RsaPriDecode(str string) (value string) {
	value, _ = gorsa.PriKeyDecrypt(str, privateKey)
	return
}

// RsaPubDecode 公钥解密
func RsaPubDecode(str string) (value string) {
	value, _ = gorsa.PublicDecrypt(str, publicKey)
	return
}

// RsaPriEncode 私钥加密
func RsaPriEncode(str string) (value string) {
	value, _ = gorsa.PriKeyEncrypt(str, privateKey)
	return
}

// RsaPubEncode 公钥加密
func RsaPubEncode(str string) (value string, err error) {
	value, err = gorsa.PublicEncrypt(str, publicKey)
	if err != nil {
		return
	}
	return
}

// ApplyPubEPriD 公钥加密私钥解密
func ApplyPubEPriD() error {
	pubEncrypt, err := gorsa.PublicEncrypt(`hello world`, publicKey)
	if err != nil {
		return err
	}

	priDecrypt, err := gorsa.PriKeyDecrypt(pubEncrypt, privateKey)
	if err != nil {
		return err
	}
	if string(priDecrypt) != `hello world` {
		return errors.New(`解密失败`)
	}
	return nil
}

// ApplyPriEPubD 公钥解密私钥加密
func ApplyPriEPubD() error {
	priEncrypt, err := gorsa.PriKeyEncrypt(`hello world`, privateKey)
	if err != nil {
		return err
	}

	pubDecrypt, err := gorsa.PublicDecrypt(priEncrypt, publicKey)
	if err != nil {
		return err
	}
	if string(pubDecrypt) != `hello world` {
		return errors.New(`解密失败`)
	}
	return nil
}
