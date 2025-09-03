package lexer

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置(現在の文字を指し示す)
	readPosition int  // これからの読み込み位置(現在の文字の次)
	ch           byte // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
