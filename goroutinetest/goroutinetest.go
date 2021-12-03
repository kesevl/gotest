package goroutinetest


func digits(num int, ciln chan int) {
	for num != 0 {
		digit := num %10
		ciln <- digit
		num /= 10
	}
	close(ciln)
}

func pingFang(num int, pingfang chan int) {
	sum := 0
	dig := make(chan int)
	go digits(num, dig)
	for digit := range dig {
		sum += digit * digit
	}
	pingfang <- sum
}

func liFang(num int, lifang chan int) {

	sum := 0
	dig := make(chan int)
	go digits(num, dig)
	for digit := range dig {
		sum += digit * digit * digit
	}
	lifang <- sum
}
