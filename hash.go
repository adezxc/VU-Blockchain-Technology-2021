package main

func hash(text []byte) (h []byte) {
	prime := 1297
	var prime2 = nextPrime(prime)
	hash := uint64(prime)
	for i := 0; i < len(text); i++ {
		placeholder := hash
		hash = (hash >> 7) ^ ((placeholder * (uint64(text[i])) >> 3))
	}


	for i := 0; i < 32; i++ {
		value := byte(i*nextPrime(prime2))+byte(hash)*text[(i+1)%len(text)]
		h = append(h, (value*byte(i+1)))
	}

	return h
}