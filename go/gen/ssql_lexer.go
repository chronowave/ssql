// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package gen

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type SSQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var ssqllexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ssqllexerLexerInit() {
	staticData := &ssqllexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "','", "'('", "')'", "'['", "']'", "'{'", "'}'", "'{&'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "AVG", "MAX", "MIN", "SUM", "COUNT",
		"PERCENTILE", "PARTITION", "EQ", "NEQ", "IN", "LT", "LE", "GE", "GT",
		"BETWEEN", "CONTAIN", "EXIST", "ISNULL", "TRUE", "FALSE", "FIND", "FROM",
		"WHERE", "ORDER_BY", "GROUP_BY", "LIMIT", "ASC", "DESC", "LOCAL", "NAME",
		"PATH", "STRING", "INTEGER", "REAL_NUMBER", "ISO8601_DATE_TIME", "SIGN",
		"IDENTIFIER", "WS",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "AVG",
		"MAX", "MIN", "SUM", "COUNT", "PERCENTILE", "PARTITION", "EQ", "NEQ",
		"IN", "LT", "LE", "GE", "GT", "BETWEEN", "CONTAIN", "EXIST", "ISNULL",
		"TRUE", "FALSE", "FIND", "FROM", "WHERE", "ORDER_BY", "GROUP_BY", "LIMIT",
		"ASC", "DESC", "LOCAL", "NAME", "PATH", "STRING", "INTEGER", "REAL_NUMBER",
		"ISO8601_DATE_TIME", "SIGN", "LETTER", "NON_ZERO_DIGIT", "DIGIT", "EXPONENT",
		"SINGLE_QUOTE", "DOUBLE_QUOTE", "BACK_QUOTE", "MICROSECOND", "TIMEZONE",
		"YEAR", "MONTH", "DAY", "HOUR", "MINUTE", "SECOND", "DQUOTA_STRING",
		"SQUOTA_STRING", "BQUOTA_STRING", "A", "B", "C", "D", "E", "F", "G",
		"H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z", "MINUS", "IDENTIFIER", "WS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 46, 634, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 2, 85, 7, 85, 2, 86, 7, 86, 2, 87, 7, 87, 2, 88, 7, 88,
		2, 89, 7, 89, 2, 90, 7, 90, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18,
		1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1,
		35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 37,
		1, 37, 3, 37, 352, 8, 37, 1, 37, 1, 37, 1, 37, 5, 37, 357, 8, 37, 10, 37,
		12, 37, 360, 9, 37, 1, 38, 1, 38, 1, 38, 1, 38, 1, 38, 5, 38, 367, 8, 38,
		10, 38, 12, 38, 370, 9, 38, 3, 38, 372, 8, 38, 1, 39, 1, 39, 1, 39, 3,
		39, 377, 8, 39, 1, 40, 4, 40, 380, 8, 40, 11, 40, 12, 40, 381, 1, 40, 1,
		40, 5, 40, 386, 8, 40, 10, 40, 12, 40, 389, 9, 40, 3, 40, 391, 8, 40, 1,
		41, 4, 41, 394, 8, 41, 11, 41, 12, 41, 395, 3, 41, 398, 8, 41, 1, 41, 1,
		41, 4, 41, 402, 8, 41, 11, 41, 12, 41, 403, 1, 41, 4, 41, 407, 8, 41, 11,
		41, 12, 41, 408, 1, 41, 1, 41, 1, 41, 1, 41, 4, 41, 415, 8, 41, 11, 41,
		12, 41, 416, 3, 41, 419, 8, 41, 1, 41, 1, 41, 4, 41, 423, 8, 41, 11, 41,
		12, 41, 424, 1, 41, 1, 41, 1, 41, 4, 41, 430, 8, 41, 11, 41, 12, 41, 431,
		1, 41, 1, 41, 3, 41, 436, 8, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1,
		42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42, 3, 42, 451, 8, 42,
		1, 42, 3, 42, 454, 8, 42, 3, 42, 456, 8, 42, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 3, 47, 468, 8, 47, 1, 47, 4,
		47, 471, 8, 47, 11, 47, 12, 47, 472, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50,
		1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52, 489,
		8, 52, 1, 52, 3, 52, 492, 8, 52, 3, 52, 494, 8, 52, 1, 53, 1, 53, 1, 53,
		1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 3, 54, 505, 8, 54, 1, 55, 1,
		55, 1, 55, 1, 55, 1, 55, 1, 55, 3, 55, 513, 8, 55, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 56, 1, 56, 3, 56, 521, 8, 56, 1, 57, 1, 57, 1, 57, 1, 57, 3,
		57, 527, 8, 57, 1, 58, 1, 58, 1, 58, 1, 58, 3, 58, 533, 8, 58, 1, 59, 1,
		59, 1, 59, 1, 59, 1, 59, 1, 59, 5, 59, 541, 8, 59, 10, 59, 12, 59, 544,
		9, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 5, 60, 554,
		8, 60, 10, 60, 12, 60, 557, 9, 60, 1, 60, 1, 60, 1, 61, 1, 61, 1, 61, 1,
		61, 1, 61, 1, 61, 5, 61, 567, 8, 61, 10, 61, 12, 61, 570, 9, 61, 1, 61,
		1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65, 1, 65, 1, 66, 1,
		66, 1, 67, 1, 67, 1, 68, 1, 68, 1, 69, 1, 69, 1, 70, 1, 70, 1, 71, 1, 71,
		1, 72, 1, 72, 1, 73, 1, 73, 1, 74, 1, 74, 1, 75, 1, 75, 1, 76, 1, 76, 1,
		77, 1, 77, 1, 78, 1, 78, 1, 79, 1, 79, 1, 80, 1, 80, 1, 81, 1, 81, 1, 82,
		1, 82, 1, 83, 1, 83, 1, 84, 1, 84, 1, 85, 1, 85, 1, 86, 1, 86, 1, 87, 1,
		87, 1, 88, 1, 88, 1, 89, 1, 89, 1, 89, 1, 90, 1, 90, 1, 90, 1, 90, 0, 0,
		91, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
		11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39,
		20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57,
		29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75,
		38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 0, 91, 0, 93, 0,
		95, 0, 97, 0, 99, 0, 101, 0, 103, 0, 105, 0, 107, 0, 109, 0, 111, 0, 113,
		0, 115, 0, 117, 0, 119, 0, 121, 0, 123, 0, 125, 0, 127, 0, 129, 0, 131,
		0, 133, 0, 135, 0, 137, 0, 139, 0, 141, 0, 143, 0, 145, 0, 147, 0, 149,
		0, 151, 0, 153, 0, 155, 0, 157, 0, 159, 0, 161, 0, 163, 0, 165, 0, 167,
		0, 169, 0, 171, 0, 173, 0, 175, 0, 177, 0, 179, 45, 181, 46, 1, 0, 44,
		2, 0, 45, 46, 95, 95, 2, 0, 43, 43, 45, 45, 2, 0, 65, 90, 97, 122, 1, 0,
		48, 57, 1, 0, 48, 48, 1, 0, 49, 57, 1, 0, 49, 49, 1, 0, 48, 50, 1, 0, 49,
		50, 1, 0, 51, 51, 1, 0, 48, 49, 1, 0, 50, 50, 1, 0, 48, 51, 1, 0, 49, 53,
		2, 0, 34, 34, 92, 92, 2, 0, 39, 39, 92, 92, 2, 0, 92, 92, 96, 96, 2, 0,
		65, 65, 97, 97, 2, 0, 66, 66, 98, 98, 2, 0, 67, 67, 99, 99, 2, 0, 68, 68,
		100, 100, 2, 0, 69, 69, 101, 101, 2, 0, 70, 70, 102, 102, 2, 0, 71, 71,
		103, 103, 2, 0, 72, 72, 104, 104, 2, 0, 73, 73, 105, 105, 2, 0, 74, 74,
		106, 106, 2, 0, 75, 75, 107, 107, 2, 0, 76, 76, 108, 108, 2, 0, 77, 77,
		109, 109, 2, 0, 78, 78, 110, 110, 2, 0, 79, 79, 111, 111, 2, 0, 80, 80,
		112, 112, 2, 0, 81, 81, 113, 113, 2, 0, 82, 82, 114, 114, 2, 0, 83, 83,
		115, 115, 2, 0, 84, 84, 116, 116, 2, 0, 85, 85, 117, 117, 2, 0, 86, 86,
		118, 118, 2, 0, 87, 87, 119, 119, 2, 0, 88, 88, 120, 120, 2, 0, 89, 89,
		121, 121, 2, 0, 90, 90, 122, 122, 3, 0, 9, 10, 13, 13, 32, 32, 634, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0,
		0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
		0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1,
		0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47,
		1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0,
		55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0,
		0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0,
		0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0,
		0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1,
		0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 179, 1, 0, 0, 0, 0, 181, 1, 0, 0, 0, 1,
		183, 1, 0, 0, 0, 3, 185, 1, 0, 0, 0, 5, 187, 1, 0, 0, 0, 7, 189, 1, 0,
		0, 0, 9, 191, 1, 0, 0, 0, 11, 193, 1, 0, 0, 0, 13, 195, 1, 0, 0, 0, 15,
		197, 1, 0, 0, 0, 17, 200, 1, 0, 0, 0, 19, 204, 1, 0, 0, 0, 21, 208, 1,
		0, 0, 0, 23, 212, 1, 0, 0, 0, 25, 216, 1, 0, 0, 0, 27, 222, 1, 0, 0, 0,
		29, 227, 1, 0, 0, 0, 31, 232, 1, 0, 0, 0, 33, 235, 1, 0, 0, 0, 35, 239,
		1, 0, 0, 0, 37, 242, 1, 0, 0, 0, 39, 245, 1, 0, 0, 0, 41, 248, 1, 0, 0,
		0, 43, 251, 1, 0, 0, 0, 45, 254, 1, 0, 0, 0, 47, 262, 1, 0, 0, 0, 49, 270,
		1, 0, 0, 0, 51, 276, 1, 0, 0, 0, 53, 283, 1, 0, 0, 0, 55, 288, 1, 0, 0,
		0, 57, 294, 1, 0, 0, 0, 59, 299, 1, 0, 0, 0, 61, 304, 1, 0, 0, 0, 63, 310,
		1, 0, 0, 0, 65, 319, 1, 0, 0, 0, 67, 328, 1, 0, 0, 0, 69, 334, 1, 0, 0,
		0, 71, 338, 1, 0, 0, 0, 73, 343, 1, 0, 0, 0, 75, 351, 1, 0, 0, 0, 77, 371,
		1, 0, 0, 0, 79, 376, 1, 0, 0, 0, 81, 390, 1, 0, 0, 0, 83, 435, 1, 0, 0,
		0, 85, 437, 1, 0, 0, 0, 87, 457, 1, 0, 0, 0, 89, 459, 1, 0, 0, 0, 91, 461,
		1, 0, 0, 0, 93, 463, 1, 0, 0, 0, 95, 465, 1, 0, 0, 0, 97, 474, 1, 0, 0,
		0, 99, 476, 1, 0, 0, 0, 101, 478, 1, 0, 0, 0, 103, 480, 1, 0, 0, 0, 105,
		493, 1, 0, 0, 0, 107, 495, 1, 0, 0, 0, 109, 504, 1, 0, 0, 0, 111, 512,
		1, 0, 0, 0, 113, 520, 1, 0, 0, 0, 115, 526, 1, 0, 0, 0, 117, 532, 1, 0,
		0, 0, 119, 534, 1, 0, 0, 0, 121, 547, 1, 0, 0, 0, 123, 560, 1, 0, 0, 0,
		125, 573, 1, 0, 0, 0, 127, 575, 1, 0, 0, 0, 129, 577, 1, 0, 0, 0, 131,
		579, 1, 0, 0, 0, 133, 581, 1, 0, 0, 0, 135, 583, 1, 0, 0, 0, 137, 585,
		1, 0, 0, 0, 139, 587, 1, 0, 0, 0, 141, 589, 1, 0, 0, 0, 143, 591, 1, 0,
		0, 0, 145, 593, 1, 0, 0, 0, 147, 595, 1, 0, 0, 0, 149, 597, 1, 0, 0, 0,
		151, 599, 1, 0, 0, 0, 153, 601, 1, 0, 0, 0, 155, 603, 1, 0, 0, 0, 157,
		605, 1, 0, 0, 0, 159, 607, 1, 0, 0, 0, 161, 609, 1, 0, 0, 0, 163, 611,
		1, 0, 0, 0, 165, 613, 1, 0, 0, 0, 167, 615, 1, 0, 0, 0, 169, 617, 1, 0,
		0, 0, 171, 619, 1, 0, 0, 0, 173, 621, 1, 0, 0, 0, 175, 623, 1, 0, 0, 0,
		177, 625, 1, 0, 0, 0, 179, 627, 1, 0, 0, 0, 181, 630, 1, 0, 0, 0, 183,
		184, 5, 44, 0, 0, 184, 2, 1, 0, 0, 0, 185, 186, 5, 40, 0, 0, 186, 4, 1,
		0, 0, 0, 187, 188, 5, 41, 0, 0, 188, 6, 1, 0, 0, 0, 189, 190, 5, 91, 0,
		0, 190, 8, 1, 0, 0, 0, 191, 192, 5, 93, 0, 0, 192, 10, 1, 0, 0, 0, 193,
		194, 5, 123, 0, 0, 194, 12, 1, 0, 0, 0, 195, 196, 5, 125, 0, 0, 196, 14,
		1, 0, 0, 0, 197, 198, 5, 123, 0, 0, 198, 199, 5, 38, 0, 0, 199, 16, 1,
		0, 0, 0, 200, 201, 3, 125, 62, 0, 201, 202, 3, 167, 83, 0, 202, 203, 3,
		137, 68, 0, 203, 18, 1, 0, 0, 0, 204, 205, 3, 149, 74, 0, 205, 206, 3,
		125, 62, 0, 206, 207, 3, 171, 85, 0, 207, 20, 1, 0, 0, 0, 208, 209, 3,
		149, 74, 0, 209, 210, 3, 141, 70, 0, 210, 211, 3, 151, 75, 0, 211, 22,
		1, 0, 0, 0, 212, 213, 3, 161, 80, 0, 213, 214, 3, 165, 82, 0, 214, 215,
		3, 149, 74, 0, 215, 24, 1, 0, 0, 0, 216, 217, 3, 129, 64, 0, 217, 218,
		3, 153, 76, 0, 218, 219, 3, 165, 82, 0, 219, 220, 3, 151, 75, 0, 220, 221,
		3, 163, 81, 0, 221, 26, 1, 0, 0, 0, 222, 223, 3, 155, 77, 0, 223, 224,
		3, 129, 64, 0, 224, 225, 3, 163, 81, 0, 225, 226, 3, 147, 73, 0, 226, 28,
		1, 0, 0, 0, 227, 228, 3, 155, 77, 0, 228, 229, 3, 125, 62, 0, 229, 230,
		3, 159, 79, 0, 230, 231, 3, 163, 81, 0, 231, 30, 1, 0, 0, 0, 232, 233,
		3, 133, 66, 0, 233, 234, 3, 157, 78, 0, 234, 32, 1, 0, 0, 0, 235, 236,
		3, 151, 75, 0, 236, 237, 3, 133, 66, 0, 237, 238, 3, 157, 78, 0, 238, 34,
		1, 0, 0, 0, 239, 240, 3, 141, 70, 0, 240, 241, 3, 151, 75, 0, 241, 36,
		1, 0, 0, 0, 242, 243, 3, 147, 73, 0, 243, 244, 3, 163, 81, 0, 244, 38,
		1, 0, 0, 0, 245, 246, 3, 147, 73, 0, 246, 247, 3, 133, 66, 0, 247, 40,
		1, 0, 0, 0, 248, 249, 3, 137, 68, 0, 249, 250, 3, 133, 66, 0, 250, 42,
		1, 0, 0, 0, 251, 252, 3, 137, 68, 0, 252, 253, 3, 163, 81, 0, 253, 44,
		1, 0, 0, 0, 254, 255, 3, 127, 63, 0, 255, 256, 3, 133, 66, 0, 256, 257,
		3, 163, 81, 0, 257, 258, 3, 169, 84, 0, 258, 259, 3, 133, 66, 0, 259, 260,
		3, 133, 66, 0, 260, 261, 3, 151, 75, 0, 261, 46, 1, 0, 0, 0, 262, 263,
		3, 129, 64, 0, 263, 264, 3, 153, 76, 0, 264, 265, 3, 151, 75, 0, 265, 266,
		3, 163, 81, 0, 266, 267, 3, 125, 62, 0, 267, 268, 3, 141, 70, 0, 268, 269,
		3, 151, 75, 0, 269, 48, 1, 0, 0, 0, 270, 271, 3, 133, 66, 0, 271, 272,
		3, 171, 85, 0, 272, 273, 3, 141, 70, 0, 273, 274, 3, 161, 80, 0, 274, 275,
		3, 163, 81, 0, 275, 50, 1, 0, 0, 0, 276, 277, 3, 141, 70, 0, 277, 278,
		3, 161, 80, 0, 278, 279, 3, 151, 75, 0, 279, 280, 3, 165, 82, 0, 280, 281,
		3, 147, 73, 0, 281, 282, 3, 147, 73, 0, 282, 52, 1, 0, 0, 0, 283, 284,
		3, 163, 81, 0, 284, 285, 3, 159, 79, 0, 285, 286, 3, 165, 82, 0, 286, 287,
		3, 133, 66, 0, 287, 54, 1, 0, 0, 0, 288, 289, 3, 135, 67, 0, 289, 290,
		3, 125, 62, 0, 290, 291, 3, 147, 73, 0, 291, 292, 3, 161, 80, 0, 292, 293,
		3, 133, 66, 0, 293, 56, 1, 0, 0, 0, 294, 295, 3, 135, 67, 0, 295, 296,
		3, 141, 70, 0, 296, 297, 3, 151, 75, 0, 297, 298, 3, 131, 65, 0, 298, 58,
		1, 0, 0, 0, 299, 300, 3, 135, 67, 0, 300, 301, 3, 159, 79, 0, 301, 302,
		3, 153, 76, 0, 302, 303, 3, 149, 74, 0, 303, 60, 1, 0, 0, 0, 304, 305,
		3, 169, 84, 0, 305, 306, 3, 139, 69, 0, 306, 307, 3, 133, 66, 0, 307, 308,
		3, 159, 79, 0, 308, 309, 3, 133, 66, 0, 309, 62, 1, 0, 0, 0, 310, 311,
		3, 153, 76, 0, 311, 312, 3, 159, 79, 0, 312, 313, 3, 131, 65, 0, 313, 314,
		3, 133, 66, 0, 314, 315, 3, 159, 79, 0, 315, 316, 3, 177, 88, 0, 316, 317,
		3, 127, 63, 0, 317, 318, 3, 173, 86, 0, 318, 64, 1, 0, 0, 0, 319, 320,
		3, 137, 68, 0, 320, 321, 3, 159, 79, 0, 321, 322, 3, 153, 76, 0, 322, 323,
		3, 165, 82, 0, 323, 324, 3, 155, 77, 0, 324, 325, 3, 177, 88, 0, 325, 326,
		3, 127, 63, 0, 326, 327, 3, 173, 86, 0, 327, 66, 1, 0, 0, 0, 328, 329,
		3, 147, 73, 0, 329, 330, 3, 141, 70, 0, 330, 331, 3, 149, 74, 0, 331, 332,
		3, 141, 70, 0, 332, 333, 3, 163, 81, 0, 333, 68, 1, 0, 0, 0, 334, 335,
		3, 125, 62, 0, 335, 336, 3, 161, 80, 0, 336, 337, 3, 129, 64, 0, 337, 70,
		1, 0, 0, 0, 338, 339, 3, 131, 65, 0, 339, 340, 3, 133, 66, 0, 340, 341,
		3, 161, 80, 0, 341, 342, 3, 129, 64, 0, 342, 72, 1, 0, 0, 0, 343, 344,
		3, 147, 73, 0, 344, 345, 3, 153, 76, 0, 345, 346, 3, 129, 64, 0, 346, 347,
		3, 125, 62, 0, 347, 348, 3, 147, 73, 0, 348, 74, 1, 0, 0, 0, 349, 352,
		3, 89, 44, 0, 350, 352, 5, 95, 0, 0, 351, 349, 1, 0, 0, 0, 351, 350, 1,
		0, 0, 0, 352, 358, 1, 0, 0, 0, 353, 357, 3, 89, 44, 0, 354, 357, 3, 93,
		46, 0, 355, 357, 7, 0, 0, 0, 356, 353, 1, 0, 0, 0, 356, 354, 1, 0, 0, 0,
		356, 355, 1, 0, 0, 0, 357, 360, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0, 358,
		359, 1, 0, 0, 0, 359, 76, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 361, 372, 5,
		47, 0, 0, 362, 363, 5, 47, 0, 0, 363, 368, 3, 75, 37, 0, 364, 365, 5, 47,
		0, 0, 365, 367, 3, 75, 37, 0, 366, 364, 1, 0, 0, 0, 367, 370, 1, 0, 0,
		0, 368, 366, 1, 0, 0, 0, 368, 369, 1, 0, 0, 0, 369, 372, 1, 0, 0, 0, 370,
		368, 1, 0, 0, 0, 371, 361, 1, 0, 0, 0, 371, 362, 1, 0, 0, 0, 372, 78, 1,
		0, 0, 0, 373, 377, 3, 119, 59, 0, 374, 377, 3, 121, 60, 0, 375, 377, 3,
		123, 61, 0, 376, 373, 1, 0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 375, 1, 0,
		0, 0, 377, 80, 1, 0, 0, 0, 378, 380, 5, 48, 0, 0, 379, 378, 1, 0, 0, 0,
		380, 381, 1, 0, 0, 0, 381, 379, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382,
		391, 1, 0, 0, 0, 383, 387, 3, 91, 45, 0, 384, 386, 3, 93, 46, 0, 385, 384,
		1, 0, 0, 0, 386, 389, 1, 0, 0, 0, 387, 385, 1, 0, 0, 0, 387, 388, 1, 0,
		0, 0, 388, 391, 1, 0, 0, 0, 389, 387, 1, 0, 0, 0, 390, 379, 1, 0, 0, 0,
		390, 383, 1, 0, 0, 0, 391, 82, 1, 0, 0, 0, 392, 394, 3, 93, 46, 0, 393,
		392, 1, 0, 0, 0, 394, 395, 1, 0, 0, 0, 395, 393, 1, 0, 0, 0, 395, 396,
		1, 0, 0, 0, 396, 398, 1, 0, 0, 0, 397, 393, 1, 0, 0, 0, 397, 398, 1, 0,
		0, 0, 398, 399, 1, 0, 0, 0, 399, 401, 5, 46, 0, 0, 400, 402, 3, 93, 46,
		0, 401, 400, 1, 0, 0, 0, 402, 403, 1, 0, 0, 0, 403, 401, 1, 0, 0, 0, 403,
		404, 1, 0, 0, 0, 404, 436, 1, 0, 0, 0, 405, 407, 3, 93, 46, 0, 406, 405,
		1, 0, 0, 0, 407, 408, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 408, 409, 1, 0,
		0, 0, 409, 410, 1, 0, 0, 0, 410, 411, 5, 46, 0, 0, 411, 412, 3, 95, 47,
		0, 412, 436, 1, 0, 0, 0, 413, 415, 3, 93, 46, 0, 414, 413, 1, 0, 0, 0,
		415, 416, 1, 0, 0, 0, 416, 414, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417,
		419, 1, 0, 0, 0, 418, 414, 1, 0, 0, 0, 418, 419, 1, 0, 0, 0, 419, 420,
		1, 0, 0, 0, 420, 422, 5, 46, 0, 0, 421, 423, 3, 93, 46, 0, 422, 421, 1,
		0, 0, 0, 423, 424, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 424, 425, 1, 0, 0,
		0, 425, 426, 1, 0, 0, 0, 426, 427, 3, 95, 47, 0, 427, 436, 1, 0, 0, 0,
		428, 430, 3, 93, 46, 0, 429, 428, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431,
		429, 1, 0, 0, 0, 431, 432, 1, 0, 0, 0, 432, 433, 1, 0, 0, 0, 433, 434,
		3, 95, 47, 0, 434, 436, 1, 0, 0, 0, 435, 397, 1, 0, 0, 0, 435, 406, 1,
		0, 0, 0, 435, 418, 1, 0, 0, 0, 435, 429, 1, 0, 0, 0, 436, 84, 1, 0, 0,
		0, 437, 438, 3, 107, 53, 0, 438, 439, 5, 45, 0, 0, 439, 440, 3, 109, 54,
		0, 440, 441, 5, 45, 0, 0, 441, 455, 3, 111, 55, 0, 442, 443, 5, 84, 0,
		0, 443, 444, 3, 113, 56, 0, 444, 445, 5, 58, 0, 0, 445, 446, 3, 115, 57,
		0, 446, 447, 5, 58, 0, 0, 447, 450, 3, 117, 58, 0, 448, 449, 5, 46, 0,
		0, 449, 451, 3, 103, 51, 0, 450, 448, 1, 0, 0, 0, 450, 451, 1, 0, 0, 0,
		451, 453, 1, 0, 0, 0, 452, 454, 3, 105, 52, 0, 453, 452, 1, 0, 0, 0, 453,
		454, 1, 0, 0, 0, 454, 456, 1, 0, 0, 0, 455, 442, 1, 0, 0, 0, 455, 456,
		1, 0, 0, 0, 456, 86, 1, 0, 0, 0, 457, 458, 7, 1, 0, 0, 458, 88, 1, 0, 0,
		0, 459, 460, 7, 2, 0, 0, 460, 90, 1, 0, 0, 0, 461, 462, 2, 49, 57, 0, 462,
		92, 1, 0, 0, 0, 463, 464, 2, 48, 57, 0, 464, 94, 1, 0, 0, 0, 465, 467,
		5, 69, 0, 0, 466, 468, 7, 1, 0, 0, 467, 466, 1, 0, 0, 0, 467, 468, 1, 0,
		0, 0, 468, 470, 1, 0, 0, 0, 469, 471, 3, 93, 46, 0, 470, 469, 1, 0, 0,
		0, 471, 472, 1, 0, 0, 0, 472, 470, 1, 0, 0, 0, 472, 473, 1, 0, 0, 0, 473,
		96, 1, 0, 0, 0, 474, 475, 5, 39, 0, 0, 475, 98, 1, 0, 0, 0, 476, 477, 5,
		34, 0, 0, 477, 100, 1, 0, 0, 0, 478, 479, 5, 96, 0, 0, 479, 102, 1, 0,
		0, 0, 480, 481, 7, 3, 0, 0, 481, 482, 7, 3, 0, 0, 482, 483, 7, 3, 0, 0,
		483, 104, 1, 0, 0, 0, 484, 494, 5, 90, 0, 0, 485, 486, 7, 1, 0, 0, 486,
		491, 3, 113, 56, 0, 487, 489, 5, 58, 0, 0, 488, 487, 1, 0, 0, 0, 488, 489,
		1, 0, 0, 0, 489, 490, 1, 0, 0, 0, 490, 492, 3, 115, 57, 0, 491, 488, 1,
		0, 0, 0, 491, 492, 1, 0, 0, 0, 492, 494, 1, 0, 0, 0, 493, 484, 1, 0, 0,
		0, 493, 485, 1, 0, 0, 0, 494, 106, 1, 0, 0, 0, 495, 496, 7, 3, 0, 0, 496,
		497, 7, 3, 0, 0, 497, 498, 7, 3, 0, 0, 498, 499, 7, 3, 0, 0, 499, 108,
		1, 0, 0, 0, 500, 501, 7, 4, 0, 0, 501, 505, 7, 5, 0, 0, 502, 503, 7, 6,
		0, 0, 503, 505, 7, 7, 0, 0, 504, 500, 1, 0, 0, 0, 504, 502, 1, 0, 0, 0,
		505, 110, 1, 0, 0, 0, 506, 507, 7, 4, 0, 0, 507, 513, 7, 5, 0, 0, 508,
		509, 7, 8, 0, 0, 509, 513, 7, 3, 0, 0, 510, 511, 7, 9, 0, 0, 511, 513,
		7, 10, 0, 0, 512, 506, 1, 0, 0, 0, 512, 508, 1, 0, 0, 0, 512, 510, 1, 0,
		0, 0, 513, 112, 1, 0, 0, 0, 514, 515, 7, 4, 0, 0, 515, 521, 7, 3, 0, 0,
		516, 517, 7, 6, 0, 0, 517, 521, 7, 3, 0, 0, 518, 519, 7, 11, 0, 0, 519,
		521, 7, 12, 0, 0, 520, 514, 1, 0, 0, 0, 520, 516, 1, 0, 0, 0, 520, 518,
		1, 0, 0, 0, 521, 114, 1, 0, 0, 0, 522, 523, 7, 4, 0, 0, 523, 527, 7, 3,
		0, 0, 524, 525, 7, 13, 0, 0, 525, 527, 7, 3, 0, 0, 526, 522, 1, 0, 0, 0,
		526, 524, 1, 0, 0, 0, 527, 116, 1, 0, 0, 0, 528, 529, 7, 4, 0, 0, 529,
		533, 7, 3, 0, 0, 530, 531, 7, 13, 0, 0, 531, 533, 7, 3, 0, 0, 532, 528,
		1, 0, 0, 0, 532, 530, 1, 0, 0, 0, 533, 118, 1, 0, 0, 0, 534, 542, 3, 99,
		49, 0, 535, 536, 5, 92, 0, 0, 536, 541, 9, 0, 0, 0, 537, 538, 5, 34, 0,
		0, 538, 541, 5, 34, 0, 0, 539, 541, 8, 14, 0, 0, 540, 535, 1, 0, 0, 0,
		540, 537, 1, 0, 0, 0, 540, 539, 1, 0, 0, 0, 541, 544, 1, 0, 0, 0, 542,
		540, 1, 0, 0, 0, 542, 543, 1, 0, 0, 0, 543, 545, 1, 0, 0, 0, 544, 542,
		1, 0, 0, 0, 545, 546, 3, 99, 49, 0, 546, 120, 1, 0, 0, 0, 547, 555, 3,
		97, 48, 0, 548, 549, 5, 92, 0, 0, 549, 554, 9, 0, 0, 0, 550, 551, 5, 39,
		0, 0, 551, 554, 5, 39, 0, 0, 552, 554, 8, 15, 0, 0, 553, 548, 1, 0, 0,
		0, 553, 550, 1, 0, 0, 0, 553, 552, 1, 0, 0, 0, 554, 557, 1, 0, 0, 0, 555,
		553, 1, 0, 0, 0, 555, 556, 1, 0, 0, 0, 556, 558, 1, 0, 0, 0, 557, 555,
		1, 0, 0, 0, 558, 559, 3, 97, 48, 0, 559, 122, 1, 0, 0, 0, 560, 568, 3,
		101, 50, 0, 561, 562, 5, 92, 0, 0, 562, 567, 9, 0, 0, 0, 563, 564, 5, 96,
		0, 0, 564, 567, 5, 96, 0, 0, 565, 567, 8, 16, 0, 0, 566, 561, 1, 0, 0,
		0, 566, 563, 1, 0, 0, 0, 566, 565, 1, 0, 0, 0, 567, 570, 1, 0, 0, 0, 568,
		566, 1, 0, 0, 0, 568, 569, 1, 0, 0, 0, 569, 571, 1, 0, 0, 0, 570, 568,
		1, 0, 0, 0, 571, 572, 3, 101, 50, 0, 572, 124, 1, 0, 0, 0, 573, 574, 7,
		17, 0, 0, 574, 126, 1, 0, 0, 0, 575, 576, 7, 18, 0, 0, 576, 128, 1, 0,
		0, 0, 577, 578, 7, 19, 0, 0, 578, 130, 1, 0, 0, 0, 579, 580, 7, 20, 0,
		0, 580, 132, 1, 0, 0, 0, 581, 582, 7, 21, 0, 0, 582, 134, 1, 0, 0, 0, 583,
		584, 7, 22, 0, 0, 584, 136, 1, 0, 0, 0, 585, 586, 7, 23, 0, 0, 586, 138,
		1, 0, 0, 0, 587, 588, 7, 24, 0, 0, 588, 140, 1, 0, 0, 0, 589, 590, 7, 25,
		0, 0, 590, 142, 1, 0, 0, 0, 591, 592, 7, 26, 0, 0, 592, 144, 1, 0, 0, 0,
		593, 594, 7, 27, 0, 0, 594, 146, 1, 0, 0, 0, 595, 596, 7, 28, 0, 0, 596,
		148, 1, 0, 0, 0, 597, 598, 7, 29, 0, 0, 598, 150, 1, 0, 0, 0, 599, 600,
		7, 30, 0, 0, 600, 152, 1, 0, 0, 0, 601, 602, 7, 31, 0, 0, 602, 154, 1,
		0, 0, 0, 603, 604, 7, 32, 0, 0, 604, 156, 1, 0, 0, 0, 605, 606, 7, 33,
		0, 0, 606, 158, 1, 0, 0, 0, 607, 608, 7, 34, 0, 0, 608, 160, 1, 0, 0, 0,
		609, 610, 7, 35, 0, 0, 610, 162, 1, 0, 0, 0, 611, 612, 7, 36, 0, 0, 612,
		164, 1, 0, 0, 0, 613, 614, 7, 37, 0, 0, 614, 166, 1, 0, 0, 0, 615, 616,
		7, 38, 0, 0, 616, 168, 1, 0, 0, 0, 617, 618, 7, 39, 0, 0, 618, 170, 1,
		0, 0, 0, 619, 620, 7, 40, 0, 0, 620, 172, 1, 0, 0, 0, 621, 622, 7, 41,
		0, 0, 622, 174, 1, 0, 0, 0, 623, 624, 7, 42, 0, 0, 624, 176, 1, 0, 0, 0,
		625, 626, 5, 45, 0, 0, 626, 178, 1, 0, 0, 0, 627, 628, 5, 36, 0, 0, 628,
		629, 3, 75, 37, 0, 629, 180, 1, 0, 0, 0, 630, 631, 7, 43, 0, 0, 631, 632,
		1, 0, 0, 0, 632, 633, 6, 90, 0, 0, 633, 182, 1, 0, 0, 0, 38, 0, 351, 356,
		358, 368, 371, 376, 381, 387, 390, 395, 397, 403, 408, 416, 418, 424, 431,
		435, 450, 453, 455, 467, 472, 488, 491, 493, 504, 512, 520, 526, 532, 540,
		542, 553, 555, 566, 568, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SSQLLexerInit initializes any static state used to implement SSQLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewSSQLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func SSQLLexerInit() {
	staticData := &ssqllexerLexerStaticData
	staticData.once.Do(ssqllexerLexerInit)
}

// NewSSQLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewSSQLLexer(input antlr.CharStream) *SSQLLexer {
	SSQLLexerInit()
	l := new(SSQLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &ssqllexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "SSQL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SSQLLexer tokens.
const (
	SSQLLexerT__0              = 1
	SSQLLexerT__1              = 2
	SSQLLexerT__2              = 3
	SSQLLexerT__3              = 4
	SSQLLexerT__4              = 5
	SSQLLexerT__5              = 6
	SSQLLexerT__6              = 7
	SSQLLexerT__7              = 8
	SSQLLexerAVG               = 9
	SSQLLexerMAX               = 10
	SSQLLexerMIN               = 11
	SSQLLexerSUM               = 12
	SSQLLexerCOUNT             = 13
	SSQLLexerPERCENTILE        = 14
	SSQLLexerPARTITION         = 15
	SSQLLexerEQ                = 16
	SSQLLexerNEQ               = 17
	SSQLLexerIN                = 18
	SSQLLexerLT                = 19
	SSQLLexerLE                = 20
	SSQLLexerGE                = 21
	SSQLLexerGT                = 22
	SSQLLexerBETWEEN           = 23
	SSQLLexerCONTAIN           = 24
	SSQLLexerEXIST             = 25
	SSQLLexerISNULL            = 26
	SSQLLexerTRUE              = 27
	SSQLLexerFALSE             = 28
	SSQLLexerFIND              = 29
	SSQLLexerFROM              = 30
	SSQLLexerWHERE             = 31
	SSQLLexerORDER_BY          = 32
	SSQLLexerGROUP_BY          = 33
	SSQLLexerLIMIT             = 34
	SSQLLexerASC               = 35
	SSQLLexerDESC              = 36
	SSQLLexerLOCAL             = 37
	SSQLLexerNAME              = 38
	SSQLLexerPATH              = 39
	SSQLLexerSTRING            = 40
	SSQLLexerINTEGER           = 41
	SSQLLexerREAL_NUMBER       = 42
	SSQLLexerISO8601_DATE_TIME = 43
	SSQLLexerSIGN              = 44
	SSQLLexerIDENTIFIER        = 45
	SSQLLexerWS                = 46
)
