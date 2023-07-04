package controllers

import (
	"crypto/rand"
	"log"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	random "math/rand"
)
var IV = []byte("1234567812345678")
const deckSize = 52

// func (caller *CallData)test() {
// 	// Generate private keys for four players
// 	playerKeys,err := caller.GeneratePlayerKeys(4)
// 	if err != nil {
// 		fmt.Println("Error generating keys:", err)
// 		return
// 	}

// 	fmt.Println("Generated AES keys:")
// 	for i, key := range playerKeys {
// 		fmt.Printf("Key %d: %s\n", i+1, key)
// 	}
// 	// Create and shuffle a deck of 52 cards
// 	deck := caller.CreateDeck()
// 	// fmt.Println("deck:",deck)
// 	caller.ShuffleDeck(deck)
// 	// deck:=[]string{"1"}
// 	// Encrypt the deck using the private keys of each player
// 	encryptedDeck := caller.EncryptDeck(deck, playerKeys)

// 	// Print the encrypted deck
// 	fmt.Println("Encrypted Deck:",encryptedDeck)
// 	decryptedDeck := caller.DecryptDeck(encryptedDeck,playerKeys)
// 	fmt.Println("Decrypted Deck:",decryptedDeck)

// }

// func (caller *CallData)GeneratePlayerKeys(numPlayers int) ([]string, error) {
func (caller *CallData)GeneratePlayerKeys(call map[string]interface{})map[string]interface{}  {
	result:=make(map[string]interface{})
	numPlayers := int(call["numPlayers"].(float64))
	keys := make([]string, numPlayers)

	for i := 0; i < 4; i++ {
		key := make([]byte, 16) // AES-256 requires a 32-byte key
		_, err := rand.Read(key)

		if err != nil {
			result=(map[string]interface{}{
				"success": false,
				"message": err.Error(),
			})	
			return result
		}

		keys[i] = hex.EncodeToString(key)
	}
	result=(map[string]interface{}{
		"success": true,
		"message": keys,
	})	
	return result
}

// createDeck creates a standard deck of 52 cards
// func (caller *CallData)CreateDeck() []string {
func (caller *CallData)CreateDeck() map[string]interface{} {
	result:=make(map[string]interface{})
	ranks := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	deck := make([]string, deckSize)

	for i, suit := range suits {
		for j, rank := range ranks {
			card := rank + " of " + suit
			deck[i*len(ranks)+j] = card
		}
	}
	result=(map[string]interface{}{
		// "success": true,
		"deck": deck,
	})	
	return result
}

// shuffleDeck shuffles the given deck of cards
// func (caller *CallData)ShuffleDeck(deck []string) {
func (caller *CallData)ShuffleDeck(call map[string]interface{})[]string {
	// result:=make(map[string]interface{})
	deck := call["deck"].([]string)
	random.Shuffle(len(deck), func(i, j int) {
		deck[i], deck[j] = deck[j], deck[i]
	})
	// result=(map[string]interface{}{
	// 	// "success": true,
	// 	"deck": deck,
	// })	
	return deck
}
// encryptDeck encrypts each card in the deck using the private keys of the players
// func (caller *CallData)EncryptDeck(deck []string, playerKeys []string) []string {
func (caller *CallData)EncryptDeck(deck []string, call map[string]interface{}) map[string]interface{} {
	result:=make(map[string]interface{})
	// deck := call["deck"].([]string)
	playerKeys,_ := call["playerKeys"].([]interface{})
	var arrmap []string
	for _,v := range playerKeys{
		arrmap= append(arrmap,v.(string))
	}


	encryptedDeck := make([]string, len(deck))
	for i, card := range deck {
		for _, priKey := range arrmap {
			encryptedCard:= encryption(card,priKey)
			card =encryptedCard
			encryptedDeck[i] = card
		}	
	}
	result=(map[string]interface{}{
		"success": true,
		"message": encryptedDeck,
	})	
	return result
}
func createCipher(key string) cipher.Block {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Fatalf("Failed to create the AES cipher: %s", err)
	}
	return c
}
func encryption(plainText string,key string) string{
	bytes := []byte(plainText)
	blockCipher := createCipher(key)
	stream := cipher.NewCTR(blockCipher, IV)
	// Buffer for storing decrypted data
	encryptedData := make([]byte, len(bytes))
	stream.XORKeyStream(encryptedData, bytes)
	result:=base64.StdEncoding.EncodeToString(encryptedData)
	return result
}
// func (caller *CallData)DecryptDeck(encrytedDeck []string, playerKeys []string) []string {
func (caller *CallData)DecryptDeck(call map[string]interface{}) map[string]interface{}  {
	fmt.Println("-----------")
	result:=make(map[string]interface{})
	encrytedDeck := call["encrytedDeck"].([]string)
	playerKeys := call["playerKeys"].([]string)
	decryptedDeck := make([]string, len(encrytedDeck))
	for i, encryptedcard := range encrytedDeck {
		for j:=len(playerKeys)-1;j>=0;j--{
				decryptedBlockBytes:= decryption(encryptedcard,playerKeys[j])
				encryptedcard =string(decryptedBlockBytes)
				decryptedDeck[i] = encryptedcard
		}
	}
	result=(map[string]interface{}{
		"success": true,
		"message": decryptedDeck,
	})	
	return result

} 

func decryption(encrypted string,key string) []byte {
	bytes,err:=base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		fmt.Println("Error decoding the encrypted string:", err)
		return nil
	}
	blockCipher := createCipher(key)
	stream := cipher.NewCTR(blockCipher, IV)
	stream.XORKeyStream(bytes, bytes)
	return bytes
}
