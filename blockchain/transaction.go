package blockchain

type Transaction struct {
	ID []byte
	sender []byte
	receiver []byte
	sum int
}

type User struct {
	name string
	public_key []byte
	balance int
}

//func() SendTransaction(trans Transaction) {



/*func() GenerateTransactions() []Transaction {
	var transactions []Transaction
	for i := 0; i < 10000; i++ {

	}
}

func() GenerateUsers() []User {
	var users []User
	for i := 0; i < 1000; i++ {

		users = append(users, )
	}
}*/