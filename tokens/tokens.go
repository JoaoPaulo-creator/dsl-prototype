package tokens

import "fmt"

type TokenType string

const (
	ILLEGAL = "illegal"
	EOF     = "eof"

	IDENT = "ident"

	// palavras reservadas
	OPTIONS = "option"
	SET     = "set"

	// ingles
	ENV       = "env"
	PRODUCT   = "product"
	PROMISE   = "promise"
	PROPOSAL  = "proposal"
	APPROVE   = "approve"
	SEND_DATA = "send_data"
	FAILENAME = "filename"
	DIR       = "dir"

	DEBUG = "debug"
	SAVE  = "save"

	// portugues
	OPCOES       = "opcoes"
	NOME_ARQUIVO = "nome_arquivo"
	SALVAR       = "salvar"
	PRODUTO      = "produto"
	PROMESSA     = "promessa"
	PROPOSTA     = "proposta"
	VEICULO      = "veiculo"
	AVALISTA     = "avalista"
	// tokens
	L_CURL  = "left_curly_braces"
	R_CURL  = "right_curly_braces"
	L_PAREN = "left_parenthesis"
	R_PAREN = "right_parenthesis"
	L_BRACE = "left_bracket_sign"
	R_BRACE = "right_bracket_sign"
	COMMA   = "comma"

	ASSIGNMENT = "assignment"

	STRING = "string"
	BOOL   = "bool"
)

type Token struct {
	Literal string
	Type    TokenType
}

var keywords = map[string]TokenType{
	"set":       SET,
	"options":   OPTIONS,
	"env":       ENV,
	"product":   PRODUTO,
	"promise":   PROMISE,
	"proposal":  PROPOSAL,
	"approve":   APPROVE,
	"send_data": SEND_DATA,
	"debug":     DEBUG,
	"save":      SAVE,
	"filename":  FAILENAME,
	"dir":       DIR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

func TokenKindString(kind TokenType) string {
	switch kind {
	case EOF:
		return "eof"
	case ENV:
		return "env"
	case SET:
		return "set"
	case PRODUCT:
		return "product"
	case PROMISE:
		return "promises"
	case PROPOSAL:
		return "proposal"
	case APPROVE:
		return "approve"
	case SEND_DATA:
		return "send_data"
	case VEICULO:
		return "veiculo"
	case AVALISTA:
		return "avalista"
	case OPTIONS:
		return "options"
	case DEBUG:
		return "debug"
	case SAVE:
		return "save"
	case L_CURL:
		return "left_curly"
	case R_CURL:
		return "right_curly"
	case COMMA:
		return "comma"
	case L_BRACE:
		return "left_bracket"
	case R_BRACE:
		return "right_bracket"
	case IDENT:
		return "identifier"
	case L_PAREN:
		return "left_parenthesis"
	case R_PAREN:
		return "left_parenthesis"
	case ASSIGNMENT:
		return "assignment"
	default:
		return fmt.Sprintf("unkown(%s)", kind)
	}
}
