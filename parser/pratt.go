package parser

import "github.com/SunA0/monkey/ast"

type (
	// 前缀解析函数
	prefixParseFn func() ast.Expression
	// 中缀解析函数
	infixParseFn func() ast.Expression
)




