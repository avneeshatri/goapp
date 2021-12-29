package main

import (
	"fmt"

	"github.com/avneeshatri/goapp/annpurna/util"
)

func main() {
	//	testPKCS8Signature()
	testPKCS8Crypto()
}

func testPKCS8Signature() {
	util := util.CryptoUtil{}
	//	secretBase64 := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCVvlOdNaNd+cXX/T7YJoyd5ahBNgyp/vkkj1TJRaCYbEKwTAlyqFAFfuOCRBpUDlQv2kmxp5pA7ykjTnnl/kE016euUTT2oOKh3sVu5cbAWNKf7N1wsG+SLgOMM+1zZMbrM67H4CvdkPkwEgVtq0QwFSHrzt8ypRnUjjvrqRctUWV91sxtNzxOCstHugjmGy8B75VaQdzfdS9kHwYalzCLYzeQy/vZp1kk0pGdE8p7NHUkVlC0u6+RsVxBwmTDhC2f2aL2I4LFoY8/8dlABzPULV+M2RX+giKbgXuSVUpQkdUHzCQCVeddBRnbGfUcm78MZcL/Jvr0KsBmoqlQpQvjAgMBAAECggEAAa3rxCSY23H0rRJzBPuRsiekNm8Dv1hTtPxCiCtWvL7iFJYVuQp4M0eUB5756J1shkjdcLeffEDyZin4JgMu+ge9IrBYdrehMw2Li03aZ8fXPOHsLdpctff8JDNk41lq/bJLsIQsGcUJdmeNVgiIsD1cbweX6zp8yqx5sY/o6aIujQSFykLpYOU6lqz+VNYYhW1TMUlZfEDwtUw0C5xFaoa7/9F4320BpncQzXk3a6lzU3qJ4TTGukkAEKVbOc+ehhcJHOXbGdCbFBgked9eQp2IPaGr9w7bUyj7nWTn66iCeby49ZMhT4kwvljvo1xJDR2WfjuJ71Km/OSpnIw2EQKBgQD3aRQAhGy4HWnvaEargzl9mWFF/dUfzA9n8QrrTatlOK/ctWmCYijZQJy6m5SKO8b81LWypOHd4IP4uPLk1+r7nhwv2eqS4uMxFPvB4Ko5mG42ktaf6Bn1Qpxp0UPNYq/K1VF0bkJHdJhBbioZ2YeRSLeOMB97e99mus6t2YugGwKBgQCa8TV0NpDS+QmHRPOJVM0E0XKUHH9HI/e60+/DGAshxZZDVBSQJErsvrFkk22GTFwPJjoXtRCGcCPc3aYYw/I+TG33fLZTRAV1QpGVGnREbUcMAr9hSJ/FjFX/KbmbYhCqVDH/PEYq/uPdz54C3WbuO/gumhOhTZjprVfP2sZP2QKBgQCnZYtJBlBhClXAzDb0mObvGjX8m0/2IaVS2H5g0r5i0EX8+SmYND4bMGJayCVqmiN6sZj2dCT2cJAaW1jMiWLOOB4uKi1SkPzTGrV0akQCkCYR+qwIqYvWq+kUl6hecKUsgcSJGIjzu+nAa39E2i+llyPKqrjCXf+7jC2G5yhXfQKBgH7C0lmmfqurxQje92Omdij39gaInfSkz3sRAzoYhJr3nOyucDhRCN9RT64aE2PpOvPF8YcNwlxKiHyc9dOTE6+2Nc1hQYUzxEdf090pjG+i+5ou8UX09YnYO08LYudlvn6rbVAEjt+EMwZ4Yiyz2A8WDwCJW3//DRg4kfviX5ZZAoGBALzUaxTaWeJWUF98o3VuVD3aGYot5KqESclt1ESAuGqb0dx4m75JxgVYVAnXQ9beJ6WzLBGvadca6Xhv9Un8oqcbb3zKe5o85+tjVli0EXcdm8x5VvJYKKJJbJ0u5CHysDzrCKjBymQ8HIOeAyUiKRRS4PHgfqwZMCObv7w7YPCN"
	signatureBase64 := "ZQ8PKbrVl46Rxv6RRykS5dFIH8u7hGFT1NnGMNaDgZjBWv9+5SVZQfc+ycuAnQwXZ/BQC9gJ49GV09JJ2gOxAzSWzvqZtkMn1AxX7CFtBDsRdyB3pV+fzwVXeUOtf1t/TzME4bUKdd3Wite0QG2pSibGSX9mK2PlBUThqAdZSaTlzNHC6IkqV/nHyoqPx5iSO8etRGih/Asf/qM/W/5hLftwNOaOQ1YAVF0/+1HeRcART0ekgF1apfOM0fcYOMEQ7gyO776n19fxa/GDWVIuQucHT6CO9F7Cc1YgTT7VViYYdoVCnY1JMnv9HKSB1YdSVc3FL/L5hsmhciEuuXQrFA=="
	publicKeyBase64 := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlb5TnTWjXfnF1/0+2CaMneWoQTYMqf75JI9UyUWgmGxCsEwJcqhQBX7jgkQaVA5UL9pJsaeaQO8pI0555f5BNNenrlE09qDiod7FbuXGwFjSn+zdcLBvki4DjDPtc2TG6zOux+Ar3ZD5MBIFbatEMBUh687fMqUZ1I4766kXLVFlfdbMbTc8TgrLR7oI5hsvAe+VWkHc33UvZB8GGpcwi2M3kMv72adZJNKRnRPKezR1JFZQtLuvkbFcQcJkw4Qtn9mi9iOCxaGPP/HZQAcz1C1fjNkV/oIim4F7klVKUJHVB8wkAlXnXQUZ2xn1HJu/DGXC/yb69CrAZqKpUKUL4wIDAQAB"
	msg := "e1d19b0adbf634346a4956624950c0d1d04142b524805dca7a4ee3d1c138253a"
	status := util.VerifyWithPublicKeyPKCS8(util.DecodeBase64(signatureBase64), msg, util.DecodeBase64(publicKeyBase64))
	fmt.Println("Verification status:", status)
}

func testPKCS8Crypto() {
	util := util.CryptoUtil{}
	msg := "hi there"
	data := []byte(msg)
	encodedString := util.EncodeBase64(data)
	fmt.Println("Encoded String:", encodedString)
	fmt.Println(string(util.DecodeBase64(encodedString)))

	privKeyBytes, pubKeyBytes := util.GeneratePKCS8PrivatePublicKeyPair()
	privateKeyBase64Str := util.EncodeBase64(privKeyBytes)
	publicKeyBase64Str := util.EncodeBase64(pubKeyBytes)

	signature := util.SignWihPKCS8PrivateKey(msg, util.DecodeBase64(privateKeyBase64Str))
	signatureBase64Str := util.EncodeBase64(signature)

	status := util.VerifyWithPublicKeyPKCS8(util.DecodeBase64(signatureBase64Str), msg, util.DecodeBase64(publicKeyBase64Str))

	fmt.Println(status)
}

func testCrypto() {
	util := util.CryptoUtil{}
	msg := "hi there"
	data := []byte(msg)
	encodedString := util.EncodeBase64(data)
	fmt.Println("Encoded String:", encodedString)
	fmt.Println(string(util.DecodeBase64(encodedString)))

	privKeyBytes, pubKeyBytes := util.GeneratePrivatePublicKeyPair()
	privateKeyBase64Str := util.EncodeBase64(privKeyBytes)
	publicKeyBase64Str := util.EncodeBase64(pubKeyBytes)

	signature := util.SignWihPrivateKey(msg, util.DecodeBase64(privateKeyBase64Str))
	signatureBase64Str := util.EncodeBase64(signature)

	status := util.VerifyWithPublicKey(util.DecodeBase64(signatureBase64Str), msg, util.DecodeBase64(publicKeyBase64Str))

	fmt.Println(status)
}
