 package asfuncss

// // import (
// // 	"crypto/sha256"
// // 	"encoding/hex"
// // 	"io/ioutil"
// // 	"os"
// // )

// // // Compare the stored hash with the calculated hash
// // func verifyHash(dataFilePath, hashFilePath, password string) (bool, error) {
// // 	// Check if data file is empty
// // 	dataContent, err := ioutil.ReadFile(dataFilePath)
// // 	if err != nil && !os.IsNotExist(err) {
// // 		return false, err
// // 	}

// // 	// Check if hash file is empty
// // 	hashContent, err := ioutil.ReadFile(hashFilePath)
// // 	if err != nil && !os.IsNotExist(err) {
// // 		return false, err
// // 	}

// // 	if len(dataContent) == 0 && len(hashContent) == 0 {
// // 		return true, nil // Both files are empty, return true
// // 	}

// // 	// Get the stored hash from the encrypted file
// // 	storedHash, err := readEncryptedHash(hashFilePath, password)
// // 	if err != nil {
// // 		return false, err
// // 	}

// // 	// Calculate the hash from the data file
// // 	hash := sha256.Sum256(dataContent)
// // 	calculatedHash := hex.EncodeToString(hash[:])

// // 	// Compare the hashes
// // 	return storedHash == calculatedHash, nil
// // }

// import (
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"crypto/sha256"
// 	"encoding/hex"
// 	"fmt"
// 	"os"
// )

// // ... (previous functions remain the same: HashFile, CreateSecuredFile, GeneratePassword)

// // OpenAndVerifySecuredFile opens a secured file, decrypts it, and compares its hash with the provided hash
// func OpenAndVerifySecuredFile(securedFilePath, password string, originalHash string) (bool, error) {
// 	// Read the secured file
// 	ciphertext, err := os.ReadFile(securedFilePath)
// 	if err != nil {
// 		return false, fmt.Errorf("error reading secured file: %v", err)
// 	}

// 	// Create cipher block
// 	block, err := aes.NewCipher([]byte(password))
// 	if err != nil {
// 		return false, fmt.Errorf("error creating cipher: %v", err)
// 	}

// 	// Create GCM
// 	gcm, err := cipher.NewGCM(block)
// 	if err != nil {
// 		return false, fmt.Errorf("error creating GCM: %v", err)
// 	}

// 	// Extract nonce and ciphertext
// 	nonceSize := gcm.NonceSize()
// 	if len(ciphertext) < nonceSize {
// 		return false, fmt.Errorf("ciphertext too short")
// 	}
// 	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

// 	// Decrypt the ciphertext
// 	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
// 	if err != nil {
// 		return false, fmt.Errorf("error decrypting: %v", err)
// 	}

// 	// Hash the decrypted content
// 	hash := sha256.Sum256(plaintext)
// 	decryptedHash := hex.EncodeToString(hash[:])

// 	// Compare hashes
// 	return decryptedHash == originalHash, nil
// }

// func checkvalid() {
// 	inputFile := "input.txt"
// 	securedFile := "secured.bin"

// 	// Hash the input file
// 	hash, err := HashFile(inputFile)
// 	if err != nil {
// 		fmt.Printf("Error hashing file: %v\n", err)
// 		return
// 	}

// 	// Generate a password from the hash
// 	password := GeneratePassword(hash)

// 	// Use the password to create a key for encryption
// 	key := []byte(password)

// 	// Create the secured file
// 	err = CreateSecuredFile(inputFile, securedFile, key)
// 	if err != nil {
// 		fmt.Printf("Error creating secured file: %v\n", err)
// 		return
// 	}

// 	fmt.Printf("Secured file created successfully. Hash: %s\n", hash)
// 	fmt.Printf("Password to access the file: %s\n", password)

// 	// Verify the secured file
// 	match, err := OpenAndVerifySecuredFile(securedFile, password, hash)
// 	if err != nil {
// 		fmt.Printf("Error verifying secured file: %v\n", err)
// 		return
// 	}

// 	if match {
// 		fmt.Println("Verification successful: The secured file's hash matches the original.")
// 	} else {
// 		fmt.Println("Verification failed: The secured file's hash does not match the original.")
// 	}
// }
