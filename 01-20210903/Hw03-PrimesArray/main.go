/*
	Lập dãy số nguyên tố
	Nhập vào số nguyên dương N < 100,000 hãy trả về mảng các số nguyên tố <= N
*/

package main

import (
	// Thu vien chuan
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Khai bao bien
	var (
		N      int64
		primes []int64
		index  int
	)

	// Tim cac so nguyen to nho hon 100,000
	primes = findPrimeNumbers(100000)

	// Nhap so nguyen duong N < 100,000
	for {
		N = readNumberFromKeyboard("Nhap so N = ")
		if N > 0 && N < 100000 {
			break
		}
		fmt.Println("  Nhap sai so nguyen duong N < 100000!")
	}

	// Tim vi tri so nguyen to nho nhat > N
	for index = 0; index < len(primes); index++ {
		if N < primes[index] {
			break
		}
	}

	// In day so nguyen to <= N
	var slide []int64 = primes[0:index]
	fmt.Println(slide)
}

// Ham tim cac so nguyen to nho hon gia tri dau vao
func findPrimeNumbers(max int64) (primes []int64) {
	// Kiem tra tham so dau vao
	if max < 2 {
		fmt.Println("So phai lon hon 2.")
		return []int64{}
	}

	// Push so nguyen to dau tien la 2
	primes = append(primes, 2)

	// Lap tim kiem so nguyen to, bo qua so chan, chi kiem tra so le, bat dau tu so 3
	var min int64 = 3
	for min <= max {
		isPrime := true
		for i := 3; i <= int(math.Sqrt(float64(min))); i++ { // Chi kiem tra so le nen bo chia 2
			if min%int64(i) == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, min)
		}
		min += 2 // Chi kiem tra so le
	}

	return primes
}

// Ham do so nguyen tu ban phim
func readNumberFromKeyboard(msg string) (result int64) {
	var err error
	reader := bufio.NewReader(os.Stdin)

	// Lap den khi nhap lieu dung thi thoi
	for {
		fmt.Print(msg)
		str, _ := reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")

		if result, err = strconv.ParseInt(str, 10, 64); err == nil {
			break
		}
		fmt.Println(err.Error())
	}

	return result
}
