package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요.")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요.")

			//키보드 버퍼를 지운다.
			stdin.ReadString('\n')
			continue
		}

		fmt.Printf("입력하신 숫자는 %d 입니다 .\n", number)
		if number%2 == 0 {
			break //짝수 검사를 한다.
		}
	}
	fmt.Println("for 문이 종료되었습니다.")
}
