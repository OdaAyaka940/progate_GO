package main

func main() {
    n := 3
    switch n {
        // caseを追加し、nが1の場合、"大吉です"と出力してください
        case 1:
            println("大吉です")
        // caseを追加し、nが2または3の場合、"吉です"と出力してください
        case 2, 3:
            println("吉です")
        // defaultを追加し、nがどのcaseとも一致しなかった場合、"凶です"と出力してください
        default:
            println("凶です")
    }
}