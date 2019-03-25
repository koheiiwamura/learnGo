package main

import (
	"fmt"
)

type MessageTo int

const (
	MESSAGE_TO_KENT    MessageTo = 0
	MESSAGE_TO_YUKIPIZ MessageTo = 1
	MESSAGE_END_ME     MessageTo = 3
	MESSAGE_END_ME_TOO MessageTo = 4
)

func main() {

	ch := make(chan MessageTo)
	done := make(chan bool)
	go kent(ch, done)
	go yukpiz(ch, done)

	// 受信、もしくは close したことを検知するために使われている(今回はcloseされる)
	<-done

	// close したことを検知するために使われている(会話中に受信は行われるので、<-done を手前に置いている)
	<-ch

	fmt.Println("会話完了")

}

func kent(ch chan MessageTo, done chan bool) {
	texts := []string{"あなたの名前を教えてください", "私はけんとです", "生ハム食べに行きましょう"}
	// ch が close されたらループを抜ける
	for v := range ch {
		switch v {
		case MESSAGE_END_ME_TOO:
			// 終了通知が来たらチャネルを閉じる
			close(ch)
			close(done)
		case MESSAGE_END_ME:
			if len(texts) == 0 {
				// セリフがなくなったら完了
				ch <- MESSAGE_END_ME_TOO
				break
			}
			// １つ下のcase文を実行
			fallthrough
		case MESSAGE_TO_KENT:
			if len(texts) > 0 {
				fmt.Println("kent:" + texts[0]) // 次のセリフを発言する
				texts = texts[1:]               // 発言済みのセリフを消す
				ch <- MESSAGE_TO_YUKIPIZ        // 相手にボールを渡す
			} else {
				ch <- MESSAGE_END_ME
			}
		case MESSAGE_TO_YUKIPIZ:
			ch <- MESSAGE_TO_YUKIPIZ
		}
	}
}

func yukpiz(ch chan MessageTo, done chan bool) {
	texts := []string{"私はゆくぴずです", "よろしくおねがいします", "行きましょう"}
	// 最初はkentから始めるため(非同期処理で)
	ch <- MESSAGE_TO_KENT
	// ch が close されたらループを抜ける
	for v := range ch {
		switch v {
		case MESSAGE_END_ME_TOO:
			close(ch)
			close(done)
		case MESSAGE_END_ME:
			if len(texts) == 0 {
				ch <- MESSAGE_END_ME_TOO
				break
			}
			// １つ下のcase文を実行
			fallthrough
		case MESSAGE_TO_KENT:
			ch <- MESSAGE_TO_KENT
		case MESSAGE_TO_YUKIPIZ:
			if len(texts) > 0 {
				fmt.Println("yukpiz:" + texts[0])
				texts = texts[1:]
				ch <- MESSAGE_TO_KENT
			} else {
				ch <- MESSAGE_END_ME
			}
		}
	}
}
