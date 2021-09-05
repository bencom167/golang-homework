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
		count  int64
		index  int64
	)

	// Tim cac so nguyen to nho hon 100,000
	primes, count = findPrimeNumbers(2, 100000)

	// Nhap so nguyen duong N < 100,000
	for {
		N = readNumberFromKeyboard("Nhap so N = ")
		if N > 0 && N < 100000 {
			break
		}
		fmt.Println("  Nhap sai so nguyen duong N < 100000!")
	}

	// Tim vi tri so nguyen to nho nhat > N
	for index = 0; index < count; index++ {
		if N < primes[index] {
			break
		}
	}

	// In day so nguyen to <= N
	var slide []int64 = primes[0:index]
	fmt.Println(slide)
}

// Ham tim so nguyen to giua 2 so
func findPrimeNumbers(min, max int64) (primes []int64, count int64) {
	// Kiem tra tham so dau vao
	if min < 2 || max < 2 {
		fmt.Println("So phai lon hon 2.")
		return []int64{}, 0
	}
	if min > max {
		fmt.Println("Nhap sai so khoang so.")
		return []int64{}, 0
	}

	// Lap tim kiem so nguyen to
	for min <= max {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(min))); i++ {
			if min%int64(i) == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, min)
			count++
		}
		min++
	}

	return primes, count
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
