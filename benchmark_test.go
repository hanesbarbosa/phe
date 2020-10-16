package phe

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"testing"

	paillier "github.com/roasbeef/go-go-gadget-paillier"
)

var lambda = [10]int64{64, 128, 256, 512, 1024, 2048, 4096, 8192, 16384, 32768}
var message1 = big.NewInt(100)
var message2 = big.NewInt(200)

func BenchmarkPHEKeyGeneration(b *testing.B) {
	var bits int
	b.StopTimer()
	for index := 0; index < len(lambda); index++ {
		bits = int(lambda[index]) / 8
		b.ResetTimer()
		b.StartTimer()
		b.Run(fmt.Sprintf("%s %d-BitPrivateKey", "With", bits), func(b *testing.B) {
			for round := 0; round < b.N; round++ {
				GenerateKeys(lambda[index])
			}
		})
		b.StopTimer()
	}
}

func BenchmarkPaillierKeyGeneration(b *testing.B) {
	var bits int
	b.StopTimer()
	for index := 0; index < len(lambda); index++ {
		bits = int(lambda[index]) / 8
		b.ResetTimer()
		b.StartTimer()
		b.Run(fmt.Sprintf("%s %d-BitPrivateKey", "With", bits), func(b *testing.B) {
			for round := 0; round < b.N; round++ {
				paillier.GenerateKey(rand.Reader, bits)
			}
		})
		b.StopTimer()
	}
}

func BenchmarkPHEEncryption(b *testing.B) {
	b.StopTimer()
	var secretKey *SecretKey
	var publicKey *PublicKey
	var bits = 0

	for index := 0; index < len(lambda); index++ {
		secretKey, publicKey = GenerateKeys(lambda[index])
		bits = int(lambda[index]) / 8
		b.ResetTimer()
		b.StartTimer()
		b.Run(fmt.Sprintf("%s %d-BitKeys", "EncryptionWith", bits), func(b *testing.B) {
			for round := 0; round < b.N; round++ {
				Encrypt(secretKey, publicKey, message1)
			}
		})
		b.StopTimer()
	}
}

func BenchmarkPaillierEncryption(b *testing.B) {
	b.StopTimer()
	var privKey *paillier.PrivateKey
	var bits = 0

	for index := 0; index < len(lambda); index++ {
		bits = int(lambda[index]) / 8
		privKey, _ = paillier.GenerateKey(rand.Reader, bits)
		b.ResetTimer()
		b.StartTimer()
		b.Run(fmt.Sprintf("%s %d-BitPrivateKey", "EncryptionWith", bits), func(b *testing.B) {
			for round := 0; round < b.N; round++ {
				paillier.Encrypt(&privKey.PublicKey, message1.Bytes())
			}
		})
		b.StopTimer()
	}
}

func BenchmarkPHEAddition(b *testing.B) {
	b.StopTimer()
	var secretKey *SecretKey
	var publicKey *PublicKey
	var bits = 0
	var cyphertexts1, cyphertexts2 [10]*Multivector
	var index = 0

	fmt.Print("Generating keys and cyphertexts...")
	for index < len(lambda) {
		secretKey, publicKey = GenerateKeys(lambda[index])
		cyphertexts1[index] = Encrypt(secretKey, publicKey, message1)
		cyphertexts2[index] = Encrypt(secretKey, publicKey, message2)
		index++
	}
	fmt.Println("done.")

	index = 0

	for index < len(lambda) {
		bits = int(lambda[index]) / 8
		b.ResetTimer()
		b.StartTimer()
		b.Run(fmt.Sprintf("%s %d-BitKeys", "AdditionWith", bits), func(b *testing.B) {
			for round := 0; round < b.N; round++ {
				Addition(publicKey, cyphertexts1[index], cyphertexts2[index])
			}
		})
		b.StopTimer()
		index++
	}
}

func BenchmarkPaillierAddition(b *testing.B) {
	b.StopTimer()
	var privKey *paillier.PrivateKey
	var bits = 0
	var cyphertexts1, cyphertexts2 [10][]byte
	var index = 0

	fmt.Print("Generating keys and cyphertexts...")
	for index < len(lambda) {
		bits = int(lambda[index]) / 8
		privKey, _ = paillier.GenerateKey(rand.Reader, bits)
		cyphertexts1[index], _ = paillier.Encrypt(&privKey.PublicKey, message1.Bytes())
		cyphertexts2[index], _ = paillier.Encrypt(&privKey.PublicKey, message2.Bytes())
		index++
	}
	fmt.Println("done.")

	index = 0

	for index < len(lambda) {
		bits = int(lambda[index]) / 8
		b.ResetTimer()
		b.StartTimer()
		b.Run(fmt.Sprintf("%s %d-BitKeys", "AdditionWith", bits), func(b *testing.B) {
			for round := 0; round < b.N; round++ {
				paillier.AddCipher(&privKey.PublicKey, cyphertexts1[index], cyphertexts2[index])
			}
		})
		b.StopTimer()
		index++
	}
}

func BenchmarkPHEDecryption(b *testing.B) {
	b.StopTimer()
	var secretKey *SecretKey
	var publicKey *PublicKey
	var bits = big.NewInt(0)
	var cyphertexts [10]*Multivector
	var index = 0

	fmt.Print("Generating keys and cyphertexts...")
	for index < len(lambda) {
		secretKey, publicKey = GenerateKeys(lambda[index])
		cyphertexts[index] = Encrypt(secretKey, publicKey, message1)
		index++
	}
	fmt.Println("done.")

	index = 0

	for index < len(lambda) {
		bits = bits.Div(big.NewInt(lambda[index]), big.NewInt(8))
		b.ResetTimer()
		b.StartTimer()
		b.Run(fmt.Sprintf("%s %d-BitKeys", "DecryptionWith", bits), func(b *testing.B) {
			for round := 0; round < b.N; round++ {
				Decrypt(secretKey, publicKey, cyphertexts[index])
			}
		})
		b.StopTimer()
		index++
	}
}

func BenchmarkPaillierDecryption(b *testing.B) {
	b.StopTimer()
	var privKey *paillier.PrivateKey
	var bits = 0
	var cyphertexts [10][]byte
	var index = 0

	fmt.Print("Generating keys and cyphertexts...")
	for index < len(lambda) {
		bits = int(lambda[index]) / 8
		privKey, _ = paillier.GenerateKey(rand.Reader, bits)
		cyphertexts[index], _ = paillier.Encrypt(&privKey.PublicKey, message1.Bytes())
		index++
	}
	fmt.Println("done.")

	index = 0

	for index < len(lambda) {
		bits = int(lambda[index]) / 8
		b.ResetTimer()
		b.StartTimer()
		b.Run(fmt.Sprintf("%s %d-BitKeys", "DecryptionWith", bits), func(b *testing.B) {
			for round := 0; round < b.N; round++ {
				paillier.Decrypt(privKey, cyphertexts[index])
			}
		})
		b.StopTimer()
		index++
	}
}
