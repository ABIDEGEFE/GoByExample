package hashing

import (
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
	"unsafe"
)

func GoHashing(){
	fmt.Println("Go hashing is running.")
}

func HashingCaller(){
	s := "Hash mee"
	digest := sha256.Sum256([]byte(s)) // returns 32 byte array
	fmt.Printf("Digest value %x\n", digest) // prints out hexadecimal value of the array.
	fmt.Println("Size of:", unsafe.Sizeof(digest)) // returns the size of the array in memory which is 32 byte

	// Generating 512 bits digest value
	s2 := "I ma 512 bits"
	digest512 := sha512.Sum512([]byte(s2)) // returns 64 bytes array
	fmt.Printf("Digest value of 512 %x\n", digest512) // Prints out hexadecimal value of the array
	fmt.Println("Size of 512:", unsafe.Sizeof(digest512)) // returns the size of the array which is 64 bytes

	// sha512 also supports using sha384
	s3 := "I am 384 bits"
	digest384 := sha512.Sum384([]byte(s3)) // returns 384 bit of array
	fmt.Printf("Digest value of 384 %x\n", digest384) // prints out hexadecimal value of the array
	fmt.Println("Size of 384: ", unsafe.Sizeof(digest384)) // returns the size of the array which is 384 bits or 48 bytes

	fileInstaller()
}
// ========================================SCENARIO (SIMULATION): CHECK WHETHER THE FILE IS ORIGINAL PUBLISHED BY THE OWNER OR NOT WHILE INSTALLING IT FROM INTERNET=================
func publisherCheckSum()[32]byte{
	originalFile := "original_file.exe"
	digest := sha256.Sum256([]byte(originalFile))
	return digest // this digest value is considered as the original deterministic value of the file which is published by the owner for users
}

func fileInstaller(){
	originalDigest := publisherCheckSum()
	downloadedFile := "original_filee.exe"

	downloadedDigest := sha256.Sum256([]byte(downloadedFile))
	// when the user install the file they can take the original pubished checksum and compare with their hashed value
	if originalDigest == downloadedDigest{
		fmt.Println("You got original uncompromised file")
	}else{
		fmt.Println("The file is not the one you expected. Insecure!")
	}
}