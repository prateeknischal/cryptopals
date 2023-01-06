package set1

func RepeatingKeyXor(plaintext, key string) (string, error) {
	plaintextBytes := []byte(plaintext)
	keyBytes := []byte(key)

	blocksize := len(keyBytes)

	ciphertextBytes := make([]byte, len(plaintextBytes))

	for i := 0; i < len(plaintextBytes); i++ {
		ciphertextBytes[i] = plaintextBytes[i] ^ keyBytes[i%blocksize]
	}

	return BytesToHex(ciphertextBytes[:]), nil
}
