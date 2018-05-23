package parse

import (
	"bufio"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// item represents a token or text string returned from the scanner.
type item struct {
	typ itemType // The type of this item.
	pos Pos      // The starting position, in bytes, of this item in the input string.
	val string   // The value of this item.
}

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	case i.typ > itemKeyword:
		return fmt.Sprintf("<%s>", i.val)
	case len(i.val) > 10:
		return fmt.Sprintf("%v:%.10q...", i.typ, i.val)
	}
	return fmt.Sprintf("%v:%q", i.typ, i.val)
}

// itemType identifies the type of lex items.
type itemType int

const (
	itemError        itemType = iota // error occurred; value is text of error
	itemChar                         // printable ASCII character; grab bag for comma etc.
	itemCharConstant                 // character constant
	itemCode                         // @code { any GO code inside }
	itemComplex                      // complex constant (1+2i); imaginary is just a number
	itemColonEquals                  // colon-equals (':=') introducing a declaration
	itemEOF
	itemEndOfLine
	itemField      // alphanumeric identifier starting with '.'
	itemIdentifier // alphanumeric identifier not starting with '.'
	itemLeftDelim  // left action delimiter
	itemLeftParen  // '(' inside action
	itemNumber     // simple number, including imaginary
	itemPipe       // pipe symbol
	itemRawString  // raw quoted string (includes quotes)
	itemRightDelim // right action delimiter
	itemRightParen // ')' inside action
	itemSpace      // run of spaces separating arguments
	itemString     // quoted string (includes quotes)
	itemText       // plain text
	itemVariable   // variable starting with '$', such as '$' or  '$1' or '$hello'
	// Directives appear after all the rest but before keywords.
	itemOpenDirective  // used only to delimit the directives
	itemExtends    // Defines parent/layout template
	itemImport     // Defines Go imports
	itemParams     // Defines input params for the template
	itemInclude    // Include file
	itemCloseDirective // Directive finished
	// Keywords appear after all the rest.
	itemKeyword  // used only to delimit the keywords
	itemDot      // the cursor, spelled '.'
	itemDefine   // define keyword
	itemElse     // else keyword
	itemEnd      // end keyword
	itemIf       // if keyword
	itemNil      // the untyped nil constant, easiest to treat as a keyword
	itemRange    // range keyword
	itemTemplate // template keyword
	itemWith     // with keyword
)

const (
	directiveChar    = "@"
	leftDelim    = "{{"
	rightDelim   = "}}"
	leftComment  = "/*"
	rightComment = "*/"
)

var key = map[string]itemType{
	"else":     itemElse,
	"end":      itemEnd,
	"if":       itemIf,
	"include": itemTemplate,
	"with":     itemWith,
}

var directives = map[string]itemType {
	"extends": itemExtends,
	"import": itemImport,
	"params": itemParams,
	"include": itemInclude,
}

const eof = -1

// Pos represents a byte position in the original input text from which
// this template was parsed.
type Pos int

// stateFn represents the state of the scanner as a function that returns the next state.
type stateFn func(*lexer) stateFn

// lexer holds the state of the scanner.
type lexer struct {
	name       string    // the name of the input; used only for error reports
	input      string    // the string being scanned
	reader     *bufio.Reader
	directive  string
	leftDelim  string    // start of action
	rightDelim string    // end of action
	state      stateFn   // the next lexing function to enter
	inside     stateFn   // Context stack for some functions
	pos        Pos       // current position in the input
	start      Pos       // start position of this item
	width      Pos       // width of last rune read from input
	lastPos    Pos       // position of most recent item returned by nextItem
	items      chan item // channel of scanned items
	parenDepth int       // nesting depth of ( ) exprs

	hasExtends bool
}

func (l *lexer) afterPosHasPrefix(s string) string {
	return strings.HasPrefix(l.input[l.pos:], s)
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isWhitepace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r' || r == '\n'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isAlphaNumeric reports whether r is an alphabetic, digit, or underscore.
func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	return r
}

// peek returns but does not consume the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// backup steps back one rune. Can only be called once per call of next.
func (l *lexer) backup() {
	l.pos -= l.width
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.input[l.start:l.pos]}
	l.start = l.pos
}

// ignore skips over the pending input before this point.
func (l *lexer) ignore() {
	l.start = l.pos
}

// accept consumes the next rune if it's from the valid set.
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set.
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
	}
	l.backup()
}

// lineNumber reports which line we're on, based on the position of
// the previous item returned by nextItem. Doing it this way
// means we don't have to worry about peek double counting.
func (l *lexer) lineNumber() int {
	return 1 + strings.Count(l.input[:l.lastPos], "\n")
}

// errorf returns an error token and terminates the scan by passing
// back a nil pointer that will be the next state, terminating l.nextItem.
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...)}
	return nil
}

// nextItem returns the next item from the input.
func (l *lexer) nextItem() item {
	item := <-l.items
	l.lastPos = item.pos
	return item
}

// lex creates a new scanner for the input string.
func lex(name, input, directive, left, right string) *lexer {
	if directive == "" {
		directive = directiveChar
	}
	if left == "" {
		left = leftDelim
	}
	if right == "" {
		right = rightDelim
	}
	l := &lexer{
		name:       name,
		input:      input,
		directive:  directive,
		leftDelim:  left,
		rightDelim: right,
		items:      make(chan item),
	}
	go l.run()
	return l
}

// run runs the state machine for the lexer.
func (l *lexer) run() {
	for l.state = lexText; l.state != nil; {
		l.state = l.state(l)
	}
}

// lexLeftDelim scans the left delimiter, which is known to be present.
func lexLeftDelim(l *lexer) stateFn {
	l.pos += Pos(len(l.leftDelim))
	if strings.HasPrefix(l.input[l.pos:], leftComment) {
		return lexComment
	}
	l.emit(itemLeftDelim)
	return lexAction
}

func lexAction(l *lexer) stateFn {
	if l.afterPosHasPrefix(leftComment) {
		return lexComment
	}
	r := l.next()
	if r != ' ' {
		return l.errorf("Mandatory space is expected after {{.")
	}
	r = l.next()
	if !isAlphaNumeric(r) {
		return l.errorf("First item of action should be action identifier")
	}
	l.parenDepth = 0
	return lexInsideAction
}

func lexInsideAction(l *lexer) stateFn {
	return nil
}

// lexComment scans a comment. The left comment marker is known to be present.
func lexComment(l *lexer) stateFn {
	l.pos += Pos(len(leftComment))
	i := strings.Index(l.input[l.pos:], rightComment)
	if i < 0 {
		return l.errorf("unclosed comment")
	}
	l.pos += Pos(i + len(rightComment))
	if !strings.HasPrefix(l.input[l.pos:], l.rightDelim) {
		return l.errorf("comment ends before closing delimiter")

	}
	l.pos += Pos(len(l.rightDelim))
	l.ignore()
	return lexText
}

// lexQuote scans a quoted string.
func lexQuote(l *lexer) stateFn {
	Loop:
	for {
		switch l.next() {
		case '\\':
			if r := l.next(); r != eof && r != '\n' {
				break
			}
			fallthrough
		case eof, '\n':
			return l.errorf("unterminated quoted string")
		case '"':
			l.backup()
			break Loop
		}
	}
	l.emit(itemString)
	l.next()
	l.ignore()
	return l.inside
}

func lexExtendsParam(l *lexer) stateFn {
	if r := l.next(); r != '"' {
		l.errorf("Expected %v, got %v", string('"'), string(r))
	}
	l.inside = lexCloseDirective
	l.ignore()
	return lexQuote
}


// lexSpace scans a run of space characters.
// One space has already been seen.
func lexWhitespace(l *lexer) stateFn {
	for isSpace(l.peek()) {
		l.next()
	}
	l.ignore()
	return l.inside
}

func lexInsideDirective(l *lexer) stateFn {
	r := l.peek()
	if isSpace(r) {
		return lexWhitespace
	} else if r == '"' {
		l.next()
		l.ignore()
		return lexQuote
	} else if isAlphaNumeric(r) {
		return lexIdentifier
	} else if isEndOfLine(r) { // TODO: How do we handle \r\n and multiple empty lines?
		l.next()
		l.emit(itemEndOfLine)
	} else if r == ')' {
		l.next()
		l.emit(itemCloseDirective)
		return lexText // TODO: Replace with lexExtended or lexText depends on template type
	} else if r == eof {
		return l.errorf("Unclosed directive")
	}
	return l.errorf("Unexpected token: %v", string(r))
}

// lexIdentifier scans an alphanumeric.
func lexIdentifier(l *lexer) stateFn {
	Loop:
	for {
		switch r := l.next(); {
		case isAlphaNumeric(r):
		// absorb.
		default:
			l.backup()
			word := l.input[l.start:l.pos]
			//if !l.atTerminator() {
			//	return l.errorf("bad character %#U", r)
			//}
			switch {
			case key[word] > itemKeyword:
				l.emit(key[word])
			//case word[0] == '.':
			//g	l.emit(itemField)
			//case word == "true", word == "false":
			//	l.emit(itemBool)
			default:
				l.emit(itemIdentifier)
			}
			break Loop
		}
	}
	return l.inside
}

func lexCloseDirective(l * lexer) stateFn {
	r := l.next()
	if r == ')' {
		l.emit(itemCloseDirective)
		return lexText
	}
	return l.errorf("Unclosed directive. Expected ')', got: %v", string(r))
}

func lexExtended(l *lexer) stateFn {
	/*
	Extended files can hae just following items on the root level:
		- @extends
		- @import
		- @params
		-- {{/* comment * /}}
		- {{ block BLOCK_NAME }}
		- {{ end [block] [BLOCK NAME] }}
	 */
	l.inside = lexExtended // TODO: Do we need it here?
	for {
		r := l.next()
		if r == eof {
			break
		}  else if isWhitepace(r) {
			//absorb
		} else {
			l.ignore()
			afterPosition := l.input[l.pos:]
			if strings.HasPrefix(afterPosition, l.leftDelim) {
				l.backup()
				return lexLeftDelim // TODO: Different for extended and unextended - should we check in lexer or parser?
			}
			return l.errorf("Unexpected char on a root level in extended tempalate: %v", string(r))
		}
	}
	l.emit(itemEOF)
	return nil
}

// lexText scans until an opening action delimiter, "{{" or a directive opening char, "@".
func lexText(l *lexer) stateFn {
	for {
		if strings.HasPrefix(l.input[l.pos:], l.leftDelim) { // If starts with "{{"
			if l.pos > l.start {
				l.emit(itemText)
			}
			return lexLeftDelim
		}
		if strings.HasPrefix(l.input[l.pos:], l.directive) { // If starts with "@" check it is following with a directive keyword and a "(".
			//fmt.Println(l.input[l.start:l.pos])
			directivePos := l.pos
			l.next() // Consume "@"
			Loop:
			for {
				switch r := l.next(); {
				case isAlphaNumeric(r):
					// absorb.
				default:
					if r == eof {
						break Loop
					}
					word := l.input[directivePos+1:l.pos-1]
					if directive, isDirective := directives[word]; isDirective && r == '(' {
						l.pos = directivePos
						if directive == itemExtends {
							if l.hasExtends {
								if directivePos > l.start {
									l.emit(itemText)
								}
								return l.errorf("Duplicate @extends")
							}
							if directivePos > l.start || l.start > 0 {
								return l.errorf("@extends should be the first thing in a file")
							}
							l.hasExtends = true
						}
						if directivePos > l.start {
							l.emit(itemText)
						}
						//fmt.Println(l.input[l.start:l.pos])
						l.pos += 1
						l.ignore()
						l.pos += Pos(len(word))
						l.emit(directive)
						l.next() // Known "("
						l.emit(itemOpenDirective)
						if directive == itemExtends {
							return lexExtendsParam
						}
						l.inside = lexInsideDirective
						return lexInsideDirective

					} else if isDirective && r != '(' && directive == itemExtends && l.start == 0 {
						return l.errorf("@extends should be followed by '('")
					}
					break Loop
				}
			}
		}
		if r := l.next(); r == eof {
			break
		} else {
			//fmt.Println(fmt.Sprintf("%s", string(r)))
		}
	}
	// Correctly reached EOF.
	if l.pos > l.start {
		l.emit(itemText)
	}
	l.emit(itemEOF)
	return nil
}