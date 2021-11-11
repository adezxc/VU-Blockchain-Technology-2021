package hashfunction 

func isPrime(n int) bool {
    // Corner cases
    if (n <= 1) {
	    return false
    }
    if (n <= 3)  {
	    return true
    }
    if (n%2 == 0 || n%3 == 0) {
	    return false
}
   
    for i := 5; i*i<=n; i+=6 {
	if (n%i == 0 || n%(i+2) == 0) {
		return false
	}
    }
   
    return true
}
 
// Function to return the smallest
// prime number greater than N
func nextPrime(n int) int {
 
    // Base case
    if (n <= 1) {
	return 2
    }

    prime := n
    found := false
 
    // Loop continuously until isPrime returns
    // true for a number greater than n
    for !found {
        prime++;
 
        if (isPrime(prime)) {
		found = true
	}
    }
 
    return prime
}