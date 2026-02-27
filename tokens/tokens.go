package tokens

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

	STRING
	BOOL
)

type Token struct {
	Literal string
	Type    TokenType
}

var keywords = map[string]TokenType{
	"set":     SET,
	"env":     ENV,
	"product": PRODUTO,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
