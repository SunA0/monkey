package repl

import (
	"bufio"
	"fmt"
	"github.com/SunA0/monkey/lexer"
	"github.com/SunA0/monkey/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.NewLexer(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

func Common(in string) (res string) {

	l := lexer.NewLexer(in)
	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {

		res = fmt.Sprintf("%s \n %s", res, tok)

	}
	return res
}
