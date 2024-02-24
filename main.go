package main

import (
	"fmt"
	"math/rand"
)

var win bool = false

func main() {

	fmt.Println("ジャンケンスタート！")
	fmt.Println("CPUと対戦しよう")

	// 勝つまで続ける
	for win == false {

		fmt.Println("あなたはなんの手を出す？")
		fmt.Println("------------")
		fmt.Println("1: グー")
		fmt.Println("2: チョキ")
		fmt.Println("3: パー")
		fmt.Println("------------")
		fmt.Println("> 何を出すか数字で入力してください")

		var userHand int

		// ユーザーの手を決定するロジック
		for userHand < 1 || userHand > 3 {
			fmt.Println("1~3の数字を入力してください")
			fmt.Scan(&userHand)
		}
		userHandString := handSignToString(userHand)

		// CPUの手を決定するロジック
		cpuHand := rand.Intn(3)
		fmt.Println(cpuHand)
		cpuHandString := handSignToString(cpuHand)

		fmt.Println("あなたの手:" + userHandString)
		fmt.Println("CPUの手:" + cpuHandString)

		// 勝敗を決めるロジック
		fmt.Println(judge(userHand, cpuHand))
		if win == false {
			fmt.Println("まだ勝てていません。もう一度やりましょう")
			fmt.Println("Press the Enter Key")
			fmt.Scanln()
		}
	}

}

// 1~3の数字をジャンケンの手に変換する
func handSignToString(num int) string {
	if num == 1 {
		return "グー👊"
	} else if num == 2 {
		return "チョキ✌️"
	} else {
		return "パー✋"
	}
}

// 勝敗を決める
func judge(userNum int, cpuNum int) string {
	if cpuNum-userNum == 1 || cpuNum-userNum == -2 {
		win = true
		return "あなたの勝ち"
	} else if cpuNum == userNum {
		return "引き分け"
	} else {
		return "あなたの負け"
	}
}
