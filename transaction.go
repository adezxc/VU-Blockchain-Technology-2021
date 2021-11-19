package main

import (
	"bufio"
	"bytes"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type Transaction struct {
	sender   []byte
	receiver []byte
	amount   int64
}

type User struct {
	name       string
	public_key []byte
	balance    int64
}

func getNames() (names []string) {
	file, err := os.Open("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	return
}

func TransactionToByteArray(trans Transaction) []byte {
	return bytes.Join(
		[][]byte{trans.receiver, trans.sender, []byte(strconv.Itoa(int(trans.amount)))}, []byte{})
}

func generateTransactions(Users []User) (transactions []Transaction) {
	for i := 0; i < 10000; i++ {
		transaction := newRandomTransaction(Users)
		transactions = append(transactions, transaction)
	}
	return transactions
}

func getUser(publickey []byte, Users []User) User {
	for _, user := range Users {
		if bytes.Equal(user.public_key, publickey) {
			return user
		}
	}
	return User{}
}

func generateUsers() (users []User) {
	for i := 0; i < 1000; i++ {
		user := newRandomUser()
		users = append(users, user)
	}

	return
}

func newRandomUser() User {
	names := getNames()
	randomIndex := rand.Intn(len(names))
	publickey := Base58Encode([]byte(names[randomIndex]))
	maxBalance := 1000000
	randomBalance := rand.Intn(maxBalance)
	return User{
		name:       names[randomIndex],
		public_key: publickey,
		balance:    int64(randomBalance),
	}
}

func newRandomTransaction(Users []User) Transaction {
	randomAmount := rand.Intn(1000000)
	randomIndex := rand.Intn(1000)
	sender := Users[randomIndex].public_key
	randomIndexs := rand.Intn(1000)
	receiver := Users[randomIndexs].public_key
	return Transaction{
		sender: sender,
		receiver: receiver,
		amount: int64(randomAmount),
	}
}

func checkTransactions(trans []Transaction, Users []User) []Transaction{
	for index, transaction := range trans {
		if (transaction.amount > getUser(transaction.sender, Users).balance) {
			trans = remove(trans, index)
		}
	}
	return trans
}

func remove(transactions []Transaction, index int) []Transaction {
    return append(transactions[:index], transactions[index+1:]...)
}

