// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package gen // SSQL
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SSQLParser struct {
	*antlr.BaseParser
}

var ssqlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ssqlParserInit() {
	staticData := &ssqlParserStaticData
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
		"start", "selection", "attribute", "aggregate", "percentile", "groupBy",
		"partition", "from", "expression", "tuple", "vector", "or", "and", "predicate",
		"eq", "neq", "gt", "ge", "lt", "le", "in", "between", "contain", "exist",
		"isnull", "boolean", "scalar", "list", "stringList", "doubleList", "intList",
		"orderBy", "order", "limit",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 366, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 1, 0, 1, 0, 3, 0, 73, 8, 0,
		1, 0, 1, 0, 1, 0, 3, 0, 78, 8, 0, 1, 0, 3, 0, 81, 8, 0, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 1, 5, 1, 88, 8, 1, 10, 1, 12, 1, 91, 9, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 98, 8, 1, 10, 1, 12, 1, 101, 9, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 106, 8, 1, 10, 1, 12, 1, 109, 9, 1, 3, 1, 111, 8, 1, 1, 2, 1, 2,
		3, 2, 115, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 122, 8, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 133, 8, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 3, 7, 143, 8, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 5, 8, 149, 8, 8, 10, 8, 12, 8, 152, 9, 8, 1, 9, 1, 9, 1, 9, 3,
		9, 157, 8, 9, 1, 10, 1, 10, 3, 10, 161, 8, 10, 1, 10, 3, 10, 164, 8, 10,
		1, 10, 1, 10, 4, 10, 168, 8, 10, 11, 10, 12, 10, 169, 3, 10, 172, 8, 10,
		1, 10, 1, 10, 1, 11, 1, 11, 4, 11, 178, 8, 11, 11, 11, 12, 11, 179, 1,
		11, 1, 11, 1, 12, 1, 12, 4, 12, 186, 8, 12, 11, 12, 12, 12, 187, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 3, 13, 204, 8, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 3, 21, 244, 8, 21, 1, 21, 1, 21, 1, 21, 3, 21, 249, 8, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 256, 8, 21, 1, 21, 1, 21, 1, 21,
		3, 21, 261, 8, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 3, 21, 271, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23,
		1, 23, 3, 23, 281, 8, 23, 1, 24, 1, 24, 1, 24, 3, 24, 286, 8, 24, 1, 25,
		1, 25, 1, 26, 3, 26, 291, 8, 26, 1, 26, 1, 26, 3, 26, 295, 8, 26, 1, 26,
		4, 26, 298, 8, 26, 11, 26, 12, 26, 299, 1, 26, 1, 26, 3, 26, 304, 8, 26,
		1, 27, 1, 27, 1, 27, 3, 27, 309, 8, 27, 1, 28, 1, 28, 1, 28, 5, 28, 314,
		8, 28, 10, 28, 12, 28, 317, 9, 28, 1, 29, 3, 29, 320, 8, 29, 1, 29, 1,
		29, 1, 29, 3, 29, 325, 8, 29, 1, 29, 5, 29, 328, 8, 29, 10, 29, 12, 29,
		331, 9, 29, 1, 30, 3, 30, 334, 8, 30, 1, 30, 1, 30, 1, 30, 3, 30, 339,
		8, 30, 1, 30, 5, 30, 342, 8, 30, 10, 30, 12, 30, 345, 9, 30, 1, 31, 1,
		31, 1, 31, 1, 31, 5, 31, 351, 8, 31, 10, 31, 12, 31, 354, 9, 31, 1, 32,
		1, 32, 3, 32, 358, 8, 32, 1, 33, 1, 33, 4, 33, 362, 8, 33, 11, 33, 12,
		33, 363, 1, 33, 0, 0, 34, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 0, 3, 1, 0, 9, 13, 1, 0, 27, 28, 1, 0, 35, 36, 389, 0, 68,
		1, 0, 0, 0, 2, 110, 1, 0, 0, 0, 4, 114, 1, 0, 0, 0, 6, 121, 1, 0, 0, 0,
		8, 123, 1, 0, 0, 0, 10, 132, 1, 0, 0, 0, 12, 134, 1, 0, 0, 0, 14, 142,
		1, 0, 0, 0, 16, 146, 1, 0, 0, 0, 18, 156, 1, 0, 0, 0, 20, 158, 1, 0, 0,
		0, 22, 175, 1, 0, 0, 0, 24, 183, 1, 0, 0, 0, 26, 203, 1, 0, 0, 0, 28, 205,
		1, 0, 0, 0, 30, 210, 1, 0, 0, 0, 32, 215, 1, 0, 0, 0, 34, 220, 1, 0, 0,
		0, 36, 225, 1, 0, 0, 0, 38, 230, 1, 0, 0, 0, 40, 235, 1, 0, 0, 0, 42, 270,
		1, 0, 0, 0, 44, 272, 1, 0, 0, 0, 46, 277, 1, 0, 0, 0, 48, 282, 1, 0, 0,
		0, 50, 287, 1, 0, 0, 0, 52, 303, 1, 0, 0, 0, 54, 308, 1, 0, 0, 0, 56, 310,
		1, 0, 0, 0, 58, 319, 1, 0, 0, 0, 60, 333, 1, 0, 0, 0, 62, 346, 1, 0, 0,
		0, 64, 355, 1, 0, 0, 0, 66, 359, 1, 0, 0, 0, 68, 69, 5, 29, 0, 0, 69, 72,
		3, 2, 1, 0, 70, 71, 5, 30, 0, 0, 71, 73, 3, 14, 7, 0, 72, 70, 1, 0, 0,
		0, 72, 73, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 75, 5, 31, 0, 0, 75, 77,
		3, 16, 8, 0, 76, 78, 3, 62, 31, 0, 77, 76, 1, 0, 0, 0, 77, 78, 1, 0, 0,
		0, 78, 80, 1, 0, 0, 0, 79, 81, 3, 66, 33, 0, 80, 79, 1, 0, 0, 0, 80, 81,
		1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 5, 0, 0, 1, 83, 1, 1, 0, 0, 0,
		84, 89, 3, 4, 2, 0, 85, 86, 5, 1, 0, 0, 86, 88, 3, 4, 2, 0, 87, 85, 1,
		0, 0, 0, 88, 91, 1, 0, 0, 0, 89, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90,
		111, 1, 0, 0, 0, 91, 89, 1, 0, 0, 0, 92, 93, 5, 33, 0, 0, 93, 94, 5, 2,
		0, 0, 94, 99, 3, 10, 5, 0, 95, 96, 5, 1, 0, 0, 96, 98, 3, 10, 5, 0, 97,
		95, 1, 0, 0, 0, 98, 101, 1, 0, 0, 0, 99, 97, 1, 0, 0, 0, 99, 100, 1, 0,
		0, 0, 100, 102, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 102, 107, 5, 3, 0, 0,
		103, 104, 5, 1, 0, 0, 104, 106, 3, 6, 3, 0, 105, 103, 1, 0, 0, 0, 106,
		109, 1, 0, 0, 0, 107, 105, 1, 0, 0, 0, 107, 108, 1, 0, 0, 0, 108, 111,
		1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 110, 84, 1, 0, 0, 0, 110, 92, 1, 0, 0,
		0, 111, 3, 1, 0, 0, 0, 112, 115, 5, 45, 0, 0, 113, 115, 3, 6, 3, 0, 114,
		112, 1, 0, 0, 0, 114, 113, 1, 0, 0, 0, 115, 5, 1, 0, 0, 0, 116, 117, 7,
		0, 0, 0, 117, 118, 5, 2, 0, 0, 118, 119, 5, 45, 0, 0, 119, 122, 5, 3, 0,
		0, 120, 122, 3, 8, 4, 0, 121, 116, 1, 0, 0, 0, 121, 120, 1, 0, 0, 0, 122,
		7, 1, 0, 0, 0, 123, 124, 5, 14, 0, 0, 124, 125, 5, 2, 0, 0, 125, 126, 5,
		45, 0, 0, 126, 127, 5, 1, 0, 0, 127, 128, 5, 42, 0, 0, 128, 129, 5, 3,
		0, 0, 129, 9, 1, 0, 0, 0, 130, 133, 5, 45, 0, 0, 131, 133, 3, 12, 6, 0,
		132, 130, 1, 0, 0, 0, 132, 131, 1, 0, 0, 0, 133, 11, 1, 0, 0, 0, 134, 135,
		5, 15, 0, 0, 135, 136, 5, 2, 0, 0, 136, 137, 5, 45, 0, 0, 137, 138, 5,
		1, 0, 0, 138, 139, 5, 41, 0, 0, 139, 140, 5, 3, 0, 0, 140, 13, 1, 0, 0,
		0, 141, 143, 5, 37, 0, 0, 142, 141, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143,
		144, 1, 0, 0, 0, 144, 145, 5, 38, 0, 0, 145, 15, 1, 0, 0, 0, 146, 150,
		3, 18, 9, 0, 147, 149, 3, 18, 9, 0, 148, 147, 1, 0, 0, 0, 149, 152, 1,
		0, 0, 0, 150, 148, 1, 0, 0, 0, 150, 151, 1, 0, 0, 0, 151, 17, 1, 0, 0,
		0, 152, 150, 1, 0, 0, 0, 153, 157, 3, 20, 10, 0, 154, 157, 3, 22, 11, 0,
		155, 157, 3, 24, 12, 0, 156, 153, 1, 0, 0, 0, 156, 154, 1, 0, 0, 0, 156,
		155, 1, 0, 0, 0, 157, 19, 1, 0, 0, 0, 158, 160, 5, 4, 0, 0, 159, 161, 5,
		45, 0, 0, 160, 159, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 163, 1, 0, 0,
		0, 162, 164, 5, 39, 0, 0, 163, 162, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164,
		171, 1, 0, 0, 0, 165, 172, 3, 26, 13, 0, 166, 168, 3, 20, 10, 0, 167, 166,
		1, 0, 0, 0, 168, 169, 1, 0, 0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0,
		0, 0, 170, 172, 1, 0, 0, 0, 171, 165, 1, 0, 0, 0, 171, 167, 1, 0, 0, 0,
		171, 172, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 174, 5, 5, 0, 0, 174,
		21, 1, 0, 0, 0, 175, 177, 5, 6, 0, 0, 176, 178, 3, 18, 9, 0, 177, 176,
		1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 177, 1, 0, 0, 0, 179, 180, 1, 0,
		0, 0, 180, 181, 1, 0, 0, 0, 181, 182, 5, 7, 0, 0, 182, 23, 1, 0, 0, 0,
		183, 185, 5, 8, 0, 0, 184, 186, 3, 18, 9, 0, 185, 184, 1, 0, 0, 0, 186,
		187, 1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 189,
		1, 0, 0, 0, 189, 190, 5, 7, 0, 0, 190, 25, 1, 0, 0, 0, 191, 204, 3, 28,
		14, 0, 192, 204, 3, 30, 15, 0, 193, 204, 3, 32, 16, 0, 194, 204, 3, 34,
		17, 0, 195, 204, 3, 36, 18, 0, 196, 204, 3, 38, 19, 0, 197, 204, 3, 40,
		20, 0, 198, 204, 3, 42, 21, 0, 199, 204, 3, 44, 22, 0, 200, 204, 3, 46,
		23, 0, 201, 204, 3, 48, 24, 0, 202, 204, 3, 50, 25, 0, 203, 191, 1, 0,
		0, 0, 203, 192, 1, 0, 0, 0, 203, 193, 1, 0, 0, 0, 203, 194, 1, 0, 0, 0,
		203, 195, 1, 0, 0, 0, 203, 196, 1, 0, 0, 0, 203, 197, 1, 0, 0, 0, 203,
		198, 1, 0, 0, 0, 203, 199, 1, 0, 0, 0, 203, 200, 1, 0, 0, 0, 203, 201,
		1, 0, 0, 0, 203, 202, 1, 0, 0, 0, 204, 27, 1, 0, 0, 0, 205, 206, 5, 16,
		0, 0, 206, 207, 5, 2, 0, 0, 207, 208, 3, 52, 26, 0, 208, 209, 5, 3, 0,
		0, 209, 29, 1, 0, 0, 0, 210, 211, 5, 17, 0, 0, 211, 212, 5, 2, 0, 0, 212,
		213, 3, 52, 26, 0, 213, 214, 5, 3, 0, 0, 214, 31, 1, 0, 0, 0, 215, 216,
		5, 22, 0, 0, 216, 217, 5, 2, 0, 0, 217, 218, 3, 52, 26, 0, 218, 219, 5,
		3, 0, 0, 219, 33, 1, 0, 0, 0, 220, 221, 5, 21, 0, 0, 221, 222, 5, 2, 0,
		0, 222, 223, 3, 52, 26, 0, 223, 224, 5, 3, 0, 0, 224, 35, 1, 0, 0, 0, 225,
		226, 5, 19, 0, 0, 226, 227, 5, 2, 0, 0, 227, 228, 3, 52, 26, 0, 228, 229,
		5, 3, 0, 0, 229, 37, 1, 0, 0, 0, 230, 231, 5, 20, 0, 0, 231, 232, 5, 2,
		0, 0, 232, 233, 3, 52, 26, 0, 233, 234, 5, 3, 0, 0, 234, 39, 1, 0, 0, 0,
		235, 236, 5, 18, 0, 0, 236, 237, 5, 2, 0, 0, 237, 238, 3, 54, 27, 0, 238,
		239, 5, 3, 0, 0, 239, 41, 1, 0, 0, 0, 240, 241, 5, 23, 0, 0, 241, 243,
		5, 2, 0, 0, 242, 244, 5, 44, 0, 0, 243, 242, 1, 0, 0, 0, 243, 244, 1, 0,
		0, 0, 244, 245, 1, 0, 0, 0, 245, 246, 5, 41, 0, 0, 246, 248, 5, 1, 0, 0,
		247, 249, 5, 44, 0, 0, 248, 247, 1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249,
		250, 1, 0, 0, 0, 250, 251, 5, 41, 0, 0, 251, 271, 5, 3, 0, 0, 252, 253,
		5, 23, 0, 0, 253, 255, 5, 2, 0, 0, 254, 256, 5, 44, 0, 0, 255, 254, 1,
		0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 257, 1, 0, 0, 0, 257, 258, 5, 42, 0,
		0, 258, 260, 5, 1, 0, 0, 259, 261, 5, 44, 0, 0, 260, 259, 1, 0, 0, 0, 260,
		261, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262, 263, 5, 42, 0, 0, 263, 271,
		5, 3, 0, 0, 264, 265, 5, 23, 0, 0, 265, 266, 5, 2, 0, 0, 266, 267, 5, 43,
		0, 0, 267, 268, 5, 1, 0, 0, 268, 269, 5, 43, 0, 0, 269, 271, 5, 3, 0, 0,
		270, 240, 1, 0, 0, 0, 270, 252, 1, 0, 0, 0, 270, 264, 1, 0, 0, 0, 271,
		43, 1, 0, 0, 0, 272, 273, 5, 24, 0, 0, 273, 274, 5, 2, 0, 0, 274, 275,
		5, 40, 0, 0, 275, 276, 5, 3, 0, 0, 276, 45, 1, 0, 0, 0, 277, 280, 5, 25,
		0, 0, 278, 279, 5, 2, 0, 0, 279, 281, 5, 3, 0, 0, 280, 278, 1, 0, 0, 0,
		280, 281, 1, 0, 0, 0, 281, 47, 1, 0, 0, 0, 282, 285, 5, 26, 0, 0, 283,
		284, 5, 2, 0, 0, 284, 286, 5, 3, 0, 0, 285, 283, 1, 0, 0, 0, 285, 286,
		1, 0, 0, 0, 286, 49, 1, 0, 0, 0, 287, 288, 7, 1, 0, 0, 288, 51, 1, 0, 0,
		0, 289, 291, 5, 44, 0, 0, 290, 289, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291,
		292, 1, 0, 0, 0, 292, 304, 5, 42, 0, 0, 293, 295, 5, 44, 0, 0, 294, 293,
		1, 0, 0, 0, 294, 295, 1, 0, 0, 0, 295, 297, 1, 0, 0, 0, 296, 298, 5, 41,
		0, 0, 297, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0,
		299, 300, 1, 0, 0, 0, 300, 304, 1, 0, 0, 0, 301, 304, 5, 43, 0, 0, 302,
		304, 5, 40, 0, 0, 303, 290, 1, 0, 0, 0, 303, 294, 1, 0, 0, 0, 303, 301,
		1, 0, 0, 0, 303, 302, 1, 0, 0, 0, 304, 53, 1, 0, 0, 0, 305, 309, 3, 56,
		28, 0, 306, 309, 3, 58, 29, 0, 307, 309, 3, 60, 30, 0, 308, 305, 1, 0,
		0, 0, 308, 306, 1, 0, 0, 0, 308, 307, 1, 0, 0, 0, 309, 55, 1, 0, 0, 0,
		310, 315, 5, 40, 0, 0, 311, 312, 5, 1, 0, 0, 312, 314, 5, 40, 0, 0, 313,
		311, 1, 0, 0, 0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315, 316,
		1, 0, 0, 0, 316, 57, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 318, 320, 5, 44,
		0, 0, 319, 318, 1, 0, 0, 0, 319, 320, 1, 0, 0, 0, 320, 321, 1, 0, 0, 0,
		321, 329, 5, 42, 0, 0, 322, 324, 5, 1, 0, 0, 323, 325, 5, 44, 0, 0, 324,
		323, 1, 0, 0, 0, 324, 325, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 328,
		5, 42, 0, 0, 327, 322, 1, 0, 0, 0, 328, 331, 1, 0, 0, 0, 329, 327, 1, 0,
		0, 0, 329, 330, 1, 0, 0, 0, 330, 59, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0,
		332, 334, 5, 44, 0, 0, 333, 332, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334,
		335, 1, 0, 0, 0, 335, 343, 5, 41, 0, 0, 336, 338, 5, 1, 0, 0, 337, 339,
		5, 44, 0, 0, 338, 337, 1, 0, 0, 0, 338, 339, 1, 0, 0, 0, 339, 340, 1, 0,
		0, 0, 340, 342, 5, 41, 0, 0, 341, 336, 1, 0, 0, 0, 342, 345, 1, 0, 0, 0,
		343, 341, 1, 0, 0, 0, 343, 344, 1, 0, 0, 0, 344, 61, 1, 0, 0, 0, 345, 343,
		1, 0, 0, 0, 346, 347, 5, 32, 0, 0, 347, 352, 3, 64, 32, 0, 348, 349, 5,
		1, 0, 0, 349, 351, 3, 64, 32, 0, 350, 348, 1, 0, 0, 0, 351, 354, 1, 0,
		0, 0, 352, 350, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353, 63, 1, 0, 0, 0,
		354, 352, 1, 0, 0, 0, 355, 357, 5, 45, 0, 0, 356, 358, 7, 2, 0, 0, 357,
		356, 1, 0, 0, 0, 357, 358, 1, 0, 0, 0, 358, 65, 1, 0, 0, 0, 359, 361, 5,
		34, 0, 0, 360, 362, 5, 41, 0, 0, 361, 360, 1, 0, 0, 0, 362, 363, 1, 0,
		0, 0, 363, 361, 1, 0, 0, 0, 363, 364, 1, 0, 0, 0, 364, 67, 1, 0, 0, 0,
		42, 72, 77, 80, 89, 99, 107, 110, 114, 121, 132, 142, 150, 156, 160, 163,
		169, 171, 179, 187, 203, 243, 248, 255, 260, 270, 280, 285, 290, 294, 299,
		303, 308, 315, 319, 324, 329, 333, 338, 343, 352, 357, 363,
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

// SSQLParserInit initializes any static state used to implement SSQLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSSQLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SSQLParserInit() {
	staticData := &ssqlParserStaticData
	staticData.once.Do(ssqlParserInit)
}

// NewSSQLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSSQLParser(input antlr.TokenStream) *SSQLParser {
	SSQLParserInit()
	this := new(SSQLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ssqlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// SSQLParser tokens.
const (
	SSQLParserEOF               = antlr.TokenEOF
	SSQLParserT__0              = 1
	SSQLParserT__1              = 2
	SSQLParserT__2              = 3
	SSQLParserT__3              = 4
	SSQLParserT__4              = 5
	SSQLParserT__5              = 6
	SSQLParserT__6              = 7
	SSQLParserT__7              = 8
	SSQLParserAVG               = 9
	SSQLParserMAX               = 10
	SSQLParserMIN               = 11
	SSQLParserSUM               = 12
	SSQLParserCOUNT             = 13
	SSQLParserPERCENTILE        = 14
	SSQLParserPARTITION         = 15
	SSQLParserEQ                = 16
	SSQLParserNEQ               = 17
	SSQLParserIN                = 18
	SSQLParserLT                = 19
	SSQLParserLE                = 20
	SSQLParserGE                = 21
	SSQLParserGT                = 22
	SSQLParserBETWEEN           = 23
	SSQLParserCONTAIN           = 24
	SSQLParserEXIST             = 25
	SSQLParserISNULL            = 26
	SSQLParserTRUE              = 27
	SSQLParserFALSE             = 28
	SSQLParserFIND              = 29
	SSQLParserFROM              = 30
	SSQLParserWHERE             = 31
	SSQLParserORDER_BY          = 32
	SSQLParserGROUP_BY          = 33
	SSQLParserLIMIT             = 34
	SSQLParserASC               = 35
	SSQLParserDESC              = 36
	SSQLParserLOCAL             = 37
	SSQLParserNAME              = 38
	SSQLParserPATH              = 39
	SSQLParserSTRING            = 40
	SSQLParserINTEGER           = 41
	SSQLParserREAL_NUMBER       = 42
	SSQLParserISO8601_DATE_TIME = 43
	SSQLParserSIGN              = 44
	SSQLParserIDENTIFIER        = 45
	SSQLParserWS                = 46
)

// SSQLParser rules.
const (
	SSQLParserRULE_start      = 0
	SSQLParserRULE_selection  = 1
	SSQLParserRULE_attribute  = 2
	SSQLParserRULE_aggregate  = 3
	SSQLParserRULE_percentile = 4
	SSQLParserRULE_groupBy    = 5
	SSQLParserRULE_partition  = 6
	SSQLParserRULE_from       = 7
	SSQLParserRULE_expression = 8
	SSQLParserRULE_tuple      = 9
	SSQLParserRULE_vector     = 10
	SSQLParserRULE_or         = 11
	SSQLParserRULE_and        = 12
	SSQLParserRULE_predicate  = 13
	SSQLParserRULE_eq         = 14
	SSQLParserRULE_neq        = 15
	SSQLParserRULE_gt         = 16
	SSQLParserRULE_ge         = 17
	SSQLParserRULE_lt         = 18
	SSQLParserRULE_le         = 19
	SSQLParserRULE_in         = 20
	SSQLParserRULE_between    = 21
	SSQLParserRULE_contain    = 22
	SSQLParserRULE_exist      = 23
	SSQLParserRULE_isnull     = 24
	SSQLParserRULE_boolean    = 25
	SSQLParserRULE_scalar     = 26
	SSQLParserRULE_list       = 27
	SSQLParserRULE_stringList = 28
	SSQLParserRULE_doubleList = 29
	SSQLParserRULE_intList    = 30
	SSQLParserRULE_orderBy    = 31
	SSQLParserRULE_order      = 32
	SSQLParserRULE_limit      = 33
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) FIND() antlr.TerminalNode {
	return s.GetToken(SSQLParserFIND, 0)
}

func (s *StartContext) Selection() ISelectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectionContext)
}

func (s *StartContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SSQLParserWHERE, 0)
}

func (s *StartContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(SSQLParserEOF, 0)
}

func (s *StartContext) FROM() antlr.TerminalNode {
	return s.GetToken(SSQLParserFROM, 0)
}

func (s *StartContext) From() IFromContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFromContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFromContext)
}

func (s *StartContext) OrderBy() IOrderByContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderByContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderByContext)
}

func (s *StartContext) Limit() ILimitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SSQLParserRULE_start)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(68)
		p.Match(SSQLParserFIND)
	}
	{
		p.SetState(69)
		p.Selection()
	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserFROM {
		{
			p.SetState(70)
			p.Match(SSQLParserFROM)
		}
		{
			p.SetState(71)
			p.From()
		}

	}
	{
		p.SetState(74)
		p.Match(SSQLParserWHERE)
	}
	{
		p.SetState(75)
		p.Expression()
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserORDER_BY {
		{
			p.SetState(76)
			p.OrderBy()
		}

	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserLIMIT {
		{
			p.SetState(79)
			p.Limit()
		}

	}
	{
		p.SetState(82)
		p.Match(SSQLParserEOF)
	}

	return localctx
}

// ISelectionContext is an interface to support dynamic dispatch.
type ISelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionContext differentiates from other interfaces.
	IsSelectionContext()
}

type SelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionContext() *SelectionContext {
	var p = new(SelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_selection
	return p
}

func (*SelectionContext) IsSelectionContext() {}

func NewSelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionContext {
	var p = new(SelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_selection

	return p
}

func (s *SelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionContext) AllAttribute() []IAttributeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeContext); ok {
			tst[i] = t.(IAttributeContext)
			i++
		}
	}

	return tst
}

func (s *SelectionContext) Attribute(i int) IAttributeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *SelectionContext) GROUP_BY() antlr.TerminalNode {
	return s.GetToken(SSQLParserGROUP_BY, 0)
}

func (s *SelectionContext) AllGroupBy() []IGroupByContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupByContext); ok {
			len++
		}
	}

	tst := make([]IGroupByContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupByContext); ok {
			tst[i] = t.(IGroupByContext)
			i++
		}
	}

	return tst
}

func (s *SelectionContext) GroupBy(i int) IGroupByContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupByContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupByContext)
}

func (s *SelectionContext) AllAggregate() []IAggregateContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAggregateContext); ok {
			len++
		}
	}

	tst := make([]IAggregateContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAggregateContext); ok {
			tst[i] = t.(IAggregateContext)
			i++
		}
	}

	return tst
}

func (s *SelectionContext) Aggregate(i int) IAggregateContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateContext)
}

func (s *SelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterSelection(s)
	}
}

func (s *SelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitSelection(s)
	}
}

func (s *SelectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitSelection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Selection() (localctx ISelectionContext) {
	this := p
	_ = this

	localctx = NewSelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SSQLParserRULE_selection)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(110)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserAVG, SSQLParserMAX, SSQLParserMIN, SSQLParserSUM, SSQLParserCOUNT, SSQLParserPERCENTILE, SSQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Attribute()
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SSQLParserT__0 {
			{
				p.SetState(85)
				p.Match(SSQLParserT__0)
			}
			{
				p.SetState(86)
				p.Attribute()
			}

			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case SSQLParserGROUP_BY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Match(SSQLParserGROUP_BY)
		}
		{
			p.SetState(93)
			p.Match(SSQLParserT__1)
		}
		{
			p.SetState(94)
			p.GroupBy()
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SSQLParserT__0 {
			{
				p.SetState(95)
				p.Match(SSQLParserT__0)
			}
			{
				p.SetState(96)
				p.GroupBy()
			}

			p.SetState(101)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(102)
			p.Match(SSQLParserT__2)
		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SSQLParserT__0 {
			{
				p.SetState(103)
				p.Match(SSQLParserT__0)
			}
			{
				p.SetState(104)
				p.Aggregate()
			}

			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *AttributeContext) Aggregate() IAggregateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateContext)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (s *AttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Attribute() (localctx IAttributeContext) {
	this := p
	_ = this

	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SSQLParserRULE_attribute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(114)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(112)
			p.Match(SSQLParserIDENTIFIER)
		}

	case SSQLParserAVG, SSQLParserMAX, SSQLParserMIN, SSQLParserSUM, SSQLParserCOUNT, SSQLParserPERCENTILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.Aggregate()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAggregateContext is an interface to support dynamic dispatch.
type IAggregateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAggregateContext differentiates from other interfaces.
	IsAggregateContext()
}

type AggregateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateContext() *AggregateContext {
	var p = new(AggregateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_aggregate
	return p
}

func (*AggregateContext) IsAggregateContext() {}

func NewAggregateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateContext {
	var p = new(AggregateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_aggregate

	return p
}

func (s *AggregateContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *AggregateContext) AVG() antlr.TerminalNode {
	return s.GetToken(SSQLParserAVG, 0)
}

func (s *AggregateContext) MAX() antlr.TerminalNode {
	return s.GetToken(SSQLParserMAX, 0)
}

func (s *AggregateContext) MIN() antlr.TerminalNode {
	return s.GetToken(SSQLParserMIN, 0)
}

func (s *AggregateContext) SUM() antlr.TerminalNode {
	return s.GetToken(SSQLParserSUM, 0)
}

func (s *AggregateContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SSQLParserCOUNT, 0)
}

func (s *AggregateContext) Percentile() IPercentileContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPercentileContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPercentileContext)
}

func (s *AggregateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterAggregate(s)
	}
}

func (s *AggregateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitAggregate(s)
	}
}

func (s *AggregateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitAggregate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Aggregate() (localctx IAggregateContext) {
	this := p
	_ = this

	localctx = NewAggregateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SSQLParserRULE_aggregate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserAVG, SSQLParserMAX, SSQLParserMIN, SSQLParserSUM, SSQLParserCOUNT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&15872) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(117)
			p.Match(SSQLParserT__1)
		}
		{
			p.SetState(118)
			p.Match(SSQLParserIDENTIFIER)
		}
		{
			p.SetState(119)
			p.Match(SSQLParserT__2)
		}

	case SSQLParserPERCENTILE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(120)
			p.Percentile()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPercentileContext is an interface to support dynamic dispatch.
type IPercentileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPercentileContext differentiates from other interfaces.
	IsPercentileContext()
}

type PercentileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPercentileContext() *PercentileContext {
	var p = new(PercentileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_percentile
	return p
}

func (*PercentileContext) IsPercentileContext() {}

func NewPercentileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PercentileContext {
	var p = new(PercentileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_percentile

	return p
}

func (s *PercentileContext) GetParser() antlr.Parser { return s.parser }

func (s *PercentileContext) PERCENTILE() antlr.TerminalNode {
	return s.GetToken(SSQLParserPERCENTILE, 0)
}

func (s *PercentileContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *PercentileContext) REAL_NUMBER() antlr.TerminalNode {
	return s.GetToken(SSQLParserREAL_NUMBER, 0)
}

func (s *PercentileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PercentileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PercentileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterPercentile(s)
	}
}

func (s *PercentileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitPercentile(s)
	}
}

func (s *PercentileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitPercentile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Percentile() (localctx IPercentileContext) {
	this := p
	_ = this

	localctx = NewPercentileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SSQLParserRULE_percentile)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(SSQLParserPERCENTILE)
	}
	{
		p.SetState(124)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(125)
		p.Match(SSQLParserIDENTIFIER)
	}
	{
		p.SetState(126)
		p.Match(SSQLParserT__0)
	}
	{
		p.SetState(127)
		p.Match(SSQLParserREAL_NUMBER)
	}
	{
		p.SetState(128)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IGroupByContext is an interface to support dynamic dispatch.
type IGroupByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupByContext differentiates from other interfaces.
	IsGroupByContext()
}

type GroupByContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByContext() *GroupByContext {
	var p = new(GroupByContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_groupBy
	return p
}

func (*GroupByContext) IsGroupByContext() {}

func NewGroupByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByContext {
	var p = new(GroupByContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_groupBy

	return p
}

func (s *GroupByContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *GroupByContext) Partition() IPartitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPartitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPartitionContext)
}

func (s *GroupByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterGroupBy(s)
	}
}

func (s *GroupByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitGroupBy(s)
	}
}

func (s *GroupByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitGroupBy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) GroupBy() (localctx IGroupByContext) {
	this := p
	_ = this

	localctx = NewGroupByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SSQLParserRULE_groupBy)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(132)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Match(SSQLParserIDENTIFIER)
		}

	case SSQLParserPARTITION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(131)
			p.Partition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPartitionContext is an interface to support dynamic dispatch.
type IPartitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPartitionContext differentiates from other interfaces.
	IsPartitionContext()
}

type PartitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPartitionContext() *PartitionContext {
	var p = new(PartitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_partition
	return p
}

func (*PartitionContext) IsPartitionContext() {}

func NewPartitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PartitionContext {
	var p = new(PartitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_partition

	return p
}

func (s *PartitionContext) GetParser() antlr.Parser { return s.parser }

func (s *PartitionContext) PARTITION() antlr.TerminalNode {
	return s.GetToken(SSQLParserPARTITION, 0)
}

func (s *PartitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *PartitionContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, 0)
}

func (s *PartitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PartitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PartitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterPartition(s)
	}
}

func (s *PartitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitPartition(s)
	}
}

func (s *PartitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitPartition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Partition() (localctx IPartitionContext) {
	this := p
	_ = this

	localctx = NewPartitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SSQLParserRULE_partition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(SSQLParserPARTITION)
	}
	{
		p.SetState(135)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(136)
		p.Match(SSQLParserIDENTIFIER)
	}
	{
		p.SetState(137)
		p.Match(SSQLParserT__0)
	}
	{
		p.SetState(138)
		p.Match(SSQLParserINTEGER)
	}
	{
		p.SetState(139)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IFromContext is an interface to support dynamic dispatch.
type IFromContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFromContext differentiates from other interfaces.
	IsFromContext()
}

type FromContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFromContext() *FromContext {
	var p = new(FromContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_from
	return p
}

func (*FromContext) IsFromContext() {}

func NewFromContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromContext {
	var p = new(FromContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_from

	return p
}

func (s *FromContext) GetParser() antlr.Parser { return s.parser }

func (s *FromContext) NAME() antlr.TerminalNode {
	return s.GetToken(SSQLParserNAME, 0)
}

func (s *FromContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(SSQLParserLOCAL, 0)
}

func (s *FromContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterFrom(s)
	}
}

func (s *FromContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitFrom(s)
	}
}

func (s *FromContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitFrom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) From() (localctx IFromContext) {
	this := p
	_ = this

	localctx = NewFromContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SSQLParserRULE_from)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserLOCAL {
		{
			p.SetState(141)
			p.Match(SSQLParserLOCAL)
		}

	}
	{
		p.SetState(144)
		p.Match(SSQLParserNAME)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllTuple() []ITupleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITupleContext); ok {
			len++
		}
	}

	tst := make([]ITupleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITupleContext); ok {
			tst[i] = t.(ITupleContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Tuple(i int) ITupleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SSQLParserRULE_expression)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Tuple()
	}
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&336) != 0 {
		{
			p.SetState(147)
			p.Tuple()
		}

		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITupleContext is an interface to support dynamic dispatch.
type ITupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTupleContext differentiates from other interfaces.
	IsTupleContext()
}

type TupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTupleContext() *TupleContext {
	var p = new(TupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_tuple
	return p
}

func (*TupleContext) IsTupleContext() {}

func NewTupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TupleContext {
	var p = new(TupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_tuple

	return p
}

func (s *TupleContext) GetParser() antlr.Parser { return s.parser }

func (s *TupleContext) Vector() IVectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorContext)
}

func (s *TupleContext) Or() IOrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrContext)
}

func (s *TupleContext) And() IAndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndContext)
}

func (s *TupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterTuple(s)
	}
}

func (s *TupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitTuple(s)
	}
}

func (s *TupleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitTuple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Tuple() (localctx ITupleContext) {
	this := p
	_ = this

	localctx = NewTupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SSQLParserRULE_tuple)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(156)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Vector()
		}

	case SSQLParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Or()
		}

	case SSQLParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(155)
			p.And()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVectorContext is an interface to support dynamic dispatch.
type IVectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVectorContext differentiates from other interfaces.
	IsVectorContext()
}

type VectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVectorContext() *VectorContext {
	var p = new(VectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_vector
	return p
}

func (*VectorContext) IsVectorContext() {}

func NewVectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectorContext {
	var p = new(VectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_vector

	return p
}

func (s *VectorContext) GetParser() antlr.Parser { return s.parser }

func (s *VectorContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *VectorContext) PATH() antlr.TerminalNode {
	return s.GetToken(SSQLParserPATH, 0)
}

func (s *VectorContext) Predicate() IPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *VectorContext) AllVector() []IVectorContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVectorContext); ok {
			len++
		}
	}

	tst := make([]IVectorContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVectorContext); ok {
			tst[i] = t.(IVectorContext)
			i++
		}
	}

	return tst
}

func (s *VectorContext) Vector(i int) IVectorContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVectorContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVectorContext)
}

func (s *VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterVector(s)
	}
}

func (s *VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitVector(s)
	}
}

func (s *VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitVector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Vector() (localctx IVectorContext) {
	this := p
	_ = this

	localctx = NewVectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SSQLParserRULE_vector)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(SSQLParserT__3)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserIDENTIFIER {
		{
			p.SetState(159)
			p.Match(SSQLParserIDENTIFIER)
		}

	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserPATH {
		{
			p.SetState(162)
			p.Match(SSQLParserPATH)
		}

	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserEQ, SSQLParserNEQ, SSQLParserIN, SSQLParserLT, SSQLParserLE, SSQLParserGE, SSQLParserGT, SSQLParserBETWEEN, SSQLParserCONTAIN, SSQLParserEXIST, SSQLParserISNULL, SSQLParserTRUE, SSQLParserFALSE:
		{
			p.SetState(165)
			p.Predicate()
		}

	case SSQLParserT__3:
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SSQLParserT__3 {
			{
				p.SetState(166)
				p.Vector()
			}

			p.SetState(169)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case SSQLParserT__4:

	default:
	}
	{
		p.SetState(173)
		p.Match(SSQLParserT__4)
	}

	return localctx
}

// IOrContext is an interface to support dynamic dispatch.
type IOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrContext differentiates from other interfaces.
	IsOrContext()
}

type OrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrContext() *OrContext {
	var p = new(OrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_or
	return p
}

func (*OrContext) IsOrContext() {}

func NewOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrContext {
	var p = new(OrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_or

	return p
}

func (s *OrContext) GetParser() antlr.Parser { return s.parser }

func (s *OrContext) AllTuple() []ITupleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITupleContext); ok {
			len++
		}
	}

	tst := make([]ITupleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITupleContext); ok {
			tst[i] = t.(ITupleContext)
			i++
		}
	}

	return tst
}

func (s *OrContext) Tuple(i int) ITupleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitOr(s)
	}
}

func (s *OrContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitOr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Or() (localctx IOrContext) {
	this := p
	_ = this

	localctx = NewOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SSQLParserRULE_or)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(SSQLParserT__5)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&336) != 0 {
		{
			p.SetState(176)
			p.Tuple()
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(181)
		p.Match(SSQLParserT__6)
	}

	return localctx
}

// IAndContext is an interface to support dynamic dispatch.
type IAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAndContext differentiates from other interfaces.
	IsAndContext()
}

type AndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndContext() *AndContext {
	var p = new(AndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_and
	return p
}

func (*AndContext) IsAndContext() {}

func NewAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndContext {
	var p = new(AndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_and

	return p
}

func (s *AndContext) GetParser() antlr.Parser { return s.parser }

func (s *AndContext) AllTuple() []ITupleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITupleContext); ok {
			len++
		}
	}

	tst := make([]ITupleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITupleContext); ok {
			tst[i] = t.(ITupleContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) Tuple(i int) ITupleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITupleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITupleContext)
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (s *AndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitAnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) And() (localctx IAndContext) {
	this := p
	_ = this

	localctx = NewAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SSQLParserRULE_and)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(SSQLParserT__7)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&336) != 0 {
		{
			p.SetState(184)
			p.Tuple()
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(189)
		p.Match(SSQLParserT__6)
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) Eq() IEqContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqContext)
}

func (s *PredicateContext) Neq() INeqContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INeqContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INeqContext)
}

func (s *PredicateContext) Gt() IGtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGtContext)
}

func (s *PredicateContext) Ge() IGeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGeContext)
}

func (s *PredicateContext) Lt() ILtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILtContext)
}

func (s *PredicateContext) Le() ILeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILeContext)
}

func (s *PredicateContext) In() IInContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInContext)
}

func (s *PredicateContext) Between() IBetweenContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBetweenContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBetweenContext)
}

func (s *PredicateContext) Contain() IContainContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContainContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContainContext)
}

func (s *PredicateContext) Exist() IExistContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExistContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExistContext)
}

func (s *PredicateContext) Isnull() IIsnullContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIsnullContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIsnullContext)
}

func (s *PredicateContext) Boolean() IBooleanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (s *PredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Predicate() (localctx IPredicateContext) {
	this := p
	_ = this

	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SSQLParserRULE_predicate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SSQLParserEQ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)
			p.Eq()
		}

	case SSQLParserNEQ:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(192)
			p.Neq()
		}

	case SSQLParserGT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(193)
			p.Gt()
		}

	case SSQLParserGE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(194)
			p.Ge()
		}

	case SSQLParserLT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(195)
			p.Lt()
		}

	case SSQLParserLE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(196)
			p.Le()
		}

	case SSQLParserIN:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(197)
			p.In()
		}

	case SSQLParserBETWEEN:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(198)
			p.Between()
		}

	case SSQLParserCONTAIN:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(199)
			p.Contain()
		}

	case SSQLParserEXIST:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(200)
			p.Exist()
		}

	case SSQLParserISNULL:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(201)
			p.Isnull()
		}

	case SSQLParserTRUE, SSQLParserFALSE:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(202)
			p.Boolean()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEqContext is an interface to support dynamic dispatch.
type IEqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqContext differentiates from other interfaces.
	IsEqContext()
}

type EqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqContext() *EqContext {
	var p = new(EqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_eq
	return p
}

func (*EqContext) IsEqContext() {}

func NewEqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqContext {
	var p = new(EqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_eq

	return p
}

func (s *EqContext) GetParser() antlr.Parser { return s.parser }

func (s *EqContext) EQ() antlr.TerminalNode {
	return s.GetToken(SSQLParserEQ, 0)
}

func (s *EqContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *EqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterEq(s)
	}
}

func (s *EqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitEq(s)
	}
}

func (s *EqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitEq(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Eq() (localctx IEqContext) {
	this := p
	_ = this

	localctx = NewEqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SSQLParserRULE_eq)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(SSQLParserEQ)
	}
	{
		p.SetState(206)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(207)
		p.Scalar()
	}
	{
		p.SetState(208)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// INeqContext is an interface to support dynamic dispatch.
type INeqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNeqContext differentiates from other interfaces.
	IsNeqContext()
}

type NeqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNeqContext() *NeqContext {
	var p = new(NeqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_neq
	return p
}

func (*NeqContext) IsNeqContext() {}

func NewNeqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NeqContext {
	var p = new(NeqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_neq

	return p
}

func (s *NeqContext) GetParser() antlr.Parser { return s.parser }

func (s *NeqContext) NEQ() antlr.TerminalNode {
	return s.GetToken(SSQLParserNEQ, 0)
}

func (s *NeqContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *NeqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NeqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NeqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterNeq(s)
	}
}

func (s *NeqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitNeq(s)
	}
}

func (s *NeqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitNeq(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Neq() (localctx INeqContext) {
	this := p
	_ = this

	localctx = NewNeqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SSQLParserRULE_neq)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(210)
		p.Match(SSQLParserNEQ)
	}
	{
		p.SetState(211)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(212)
		p.Scalar()
	}
	{
		p.SetState(213)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IGtContext is an interface to support dynamic dispatch.
type IGtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGtContext differentiates from other interfaces.
	IsGtContext()
}

type GtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGtContext() *GtContext {
	var p = new(GtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_gt
	return p
}

func (*GtContext) IsGtContext() {}

func NewGtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GtContext {
	var p = new(GtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_gt

	return p
}

func (s *GtContext) GetParser() antlr.Parser { return s.parser }

func (s *GtContext) GT() antlr.TerminalNode {
	return s.GetToken(SSQLParserGT, 0)
}

func (s *GtContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *GtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterGt(s)
	}
}

func (s *GtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitGt(s)
	}
}

func (s *GtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitGt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Gt() (localctx IGtContext) {
	this := p
	_ = this

	localctx = NewGtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SSQLParserRULE_gt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(SSQLParserGT)
	}
	{
		p.SetState(216)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(217)
		p.Scalar()
	}
	{
		p.SetState(218)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IGeContext is an interface to support dynamic dispatch.
type IGeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeContext differentiates from other interfaces.
	IsGeContext()
}

type GeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeContext() *GeContext {
	var p = new(GeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_ge
	return p
}

func (*GeContext) IsGeContext() {}

func NewGeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeContext {
	var p = new(GeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_ge

	return p
}

func (s *GeContext) GetParser() antlr.Parser { return s.parser }

func (s *GeContext) GE() antlr.TerminalNode {
	return s.GetToken(SSQLParserGE, 0)
}

func (s *GeContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *GeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterGe(s)
	}
}

func (s *GeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitGe(s)
	}
}

func (s *GeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitGe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Ge() (localctx IGeContext) {
	this := p
	_ = this

	localctx = NewGeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SSQLParserRULE_ge)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		p.Match(SSQLParserGE)
	}
	{
		p.SetState(221)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(222)
		p.Scalar()
	}
	{
		p.SetState(223)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// ILtContext is an interface to support dynamic dispatch.
type ILtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLtContext differentiates from other interfaces.
	IsLtContext()
}

type LtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLtContext() *LtContext {
	var p = new(LtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_lt
	return p
}

func (*LtContext) IsLtContext() {}

func NewLtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LtContext {
	var p = new(LtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_lt

	return p
}

func (s *LtContext) GetParser() antlr.Parser { return s.parser }

func (s *LtContext) LT() antlr.TerminalNode {
	return s.GetToken(SSQLParserLT, 0)
}

func (s *LtContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *LtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterLt(s)
	}
}

func (s *LtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitLt(s)
	}
}

func (s *LtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitLt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Lt() (localctx ILtContext) {
	this := p
	_ = this

	localctx = NewLtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SSQLParserRULE_lt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(225)
		p.Match(SSQLParserLT)
	}
	{
		p.SetState(226)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(227)
		p.Scalar()
	}
	{
		p.SetState(228)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// ILeContext is an interface to support dynamic dispatch.
type ILeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeContext differentiates from other interfaces.
	IsLeContext()
}

type LeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeContext() *LeContext {
	var p = new(LeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_le
	return p
}

func (*LeContext) IsLeContext() {}

func NewLeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeContext {
	var p = new(LeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_le

	return p
}

func (s *LeContext) GetParser() antlr.Parser { return s.parser }

func (s *LeContext) LE() antlr.TerminalNode {
	return s.GetToken(SSQLParserLE, 0)
}

func (s *LeContext) Scalar() IScalarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IScalarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IScalarContext)
}

func (s *LeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterLe(s)
	}
}

func (s *LeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitLe(s)
	}
}

func (s *LeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitLe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Le() (localctx ILeContext) {
	this := p
	_ = this

	localctx = NewLeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SSQLParserRULE_le)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(SSQLParserLE)
	}
	{
		p.SetState(231)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(232)
		p.Scalar()
	}
	{
		p.SetState(233)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IInContext is an interface to support dynamic dispatch.
type IInContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInContext differentiates from other interfaces.
	IsInContext()
}

type InContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInContext() *InContext {
	var p = new(InContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_in
	return p
}

func (*InContext) IsInContext() {}

func NewInContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InContext {
	var p = new(InContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_in

	return p
}

func (s *InContext) GetParser() antlr.Parser { return s.parser }

func (s *InContext) IN() antlr.TerminalNode {
	return s.GetToken(SSQLParserIN, 0)
}

func (s *InContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *InContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterIn(s)
	}
}

func (s *InContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitIn(s)
	}
}

func (s *InContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitIn(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) In() (localctx IInContext) {
	this := p
	_ = this

	localctx = NewInContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SSQLParserRULE_in)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.Match(SSQLParserIN)
	}
	{
		p.SetState(236)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(237)
		p.List()
	}
	{
		p.SetState(238)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IBetweenContext is an interface to support dynamic dispatch.
type IBetweenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBetweenContext differentiates from other interfaces.
	IsBetweenContext()
}

type BetweenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBetweenContext() *BetweenContext {
	var p = new(BetweenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_between
	return p
}

func (*BetweenContext) IsBetweenContext() {}

func NewBetweenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BetweenContext {
	var p = new(BetweenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_between

	return p
}

func (s *BetweenContext) GetParser() antlr.Parser { return s.parser }

func (s *BetweenContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(SSQLParserBETWEEN, 0)
}

func (s *BetweenContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserINTEGER)
}

func (s *BetweenContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, i)
}

func (s *BetweenContext) AllSIGN() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserSIGN)
}

func (s *BetweenContext) SIGN(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserSIGN, i)
}

func (s *BetweenContext) AllREAL_NUMBER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserREAL_NUMBER)
}

func (s *BetweenContext) REAL_NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserREAL_NUMBER, i)
}

func (s *BetweenContext) AllISO8601_DATE_TIME() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserISO8601_DATE_TIME)
}

func (s *BetweenContext) ISO8601_DATE_TIME(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserISO8601_DATE_TIME, i)
}

func (s *BetweenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BetweenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterBetween(s)
	}
}

func (s *BetweenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitBetween(s)
	}
}

func (s *BetweenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitBetween(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Between() (localctx IBetweenContext) {
	this := p
	_ = this

	localctx = NewBetweenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SSQLParserRULE_between)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)
			p.Match(SSQLParserBETWEEN)
		}
		{
			p.SetState(241)
			p.Match(SSQLParserT__1)
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SSQLParserSIGN {
			{
				p.SetState(242)
				p.Match(SSQLParserSIGN)
			}

		}
		{
			p.SetState(245)
			p.Match(SSQLParserINTEGER)
		}
		{
			p.SetState(246)
			p.Match(SSQLParserT__0)
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SSQLParserSIGN {
			{
				p.SetState(247)
				p.Match(SSQLParserSIGN)
			}

		}
		{
			p.SetState(250)
			p.Match(SSQLParserINTEGER)
		}
		{
			p.SetState(251)
			p.Match(SSQLParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(252)
			p.Match(SSQLParserBETWEEN)
		}
		{
			p.SetState(253)
			p.Match(SSQLParserT__1)
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SSQLParserSIGN {
			{
				p.SetState(254)
				p.Match(SSQLParserSIGN)
			}

		}
		{
			p.SetState(257)
			p.Match(SSQLParserREAL_NUMBER)
		}
		{
			p.SetState(258)
			p.Match(SSQLParserT__0)
		}
		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SSQLParserSIGN {
			{
				p.SetState(259)
				p.Match(SSQLParserSIGN)
			}

		}
		{
			p.SetState(262)
			p.Match(SSQLParserREAL_NUMBER)
		}
		{
			p.SetState(263)
			p.Match(SSQLParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(264)
			p.Match(SSQLParserBETWEEN)
		}
		{
			p.SetState(265)
			p.Match(SSQLParserT__1)
		}
		{
			p.SetState(266)
			p.Match(SSQLParserISO8601_DATE_TIME)
		}
		{
			p.SetState(267)
			p.Match(SSQLParserT__0)
		}
		{
			p.SetState(268)
			p.Match(SSQLParserISO8601_DATE_TIME)
		}
		{
			p.SetState(269)
			p.Match(SSQLParserT__2)
		}

	}

	return localctx
}

// IContainContext is an interface to support dynamic dispatch.
type IContainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContainContext differentiates from other interfaces.
	IsContainContext()
}

type ContainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainContext() *ContainContext {
	var p = new(ContainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_contain
	return p
}

func (*ContainContext) IsContainContext() {}

func NewContainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContainContext {
	var p = new(ContainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_contain

	return p
}

func (s *ContainContext) GetParser() antlr.Parser { return s.parser }

func (s *ContainContext) CONTAIN() antlr.TerminalNode {
	return s.GetToken(SSQLParserCONTAIN, 0)
}

func (s *ContainContext) STRING() antlr.TerminalNode {
	return s.GetToken(SSQLParserSTRING, 0)
}

func (s *ContainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterContain(s)
	}
}

func (s *ContainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitContain(s)
	}
}

func (s *ContainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitContain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Contain() (localctx IContainContext) {
	this := p
	_ = this

	localctx = NewContainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SSQLParserRULE_contain)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(272)
		p.Match(SSQLParserCONTAIN)
	}
	{
		p.SetState(273)
		p.Match(SSQLParserT__1)
	}
	{
		p.SetState(274)
		p.Match(SSQLParserSTRING)
	}
	{
		p.SetState(275)
		p.Match(SSQLParserT__2)
	}

	return localctx
}

// IExistContext is an interface to support dynamic dispatch.
type IExistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExistContext differentiates from other interfaces.
	IsExistContext()
}

type ExistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExistContext() *ExistContext {
	var p = new(ExistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_exist
	return p
}

func (*ExistContext) IsExistContext() {}

func NewExistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExistContext {
	var p = new(ExistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_exist

	return p
}

func (s *ExistContext) GetParser() antlr.Parser { return s.parser }

func (s *ExistContext) EXIST() antlr.TerminalNode {
	return s.GetToken(SSQLParserEXIST, 0)
}

func (s *ExistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterExist(s)
	}
}

func (s *ExistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitExist(s)
	}
}

func (s *ExistContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitExist(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Exist() (localctx IExistContext) {
	this := p
	_ = this

	localctx = NewExistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SSQLParserRULE_exist)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Match(SSQLParserEXIST)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserT__1 {
		{
			p.SetState(278)
			p.Match(SSQLParserT__1)
		}
		{
			p.SetState(279)
			p.Match(SSQLParserT__2)
		}

	}

	return localctx
}

// IIsnullContext is an interface to support dynamic dispatch.
type IIsnullContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsnullContext differentiates from other interfaces.
	IsIsnullContext()
}

type IsnullContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsnullContext() *IsnullContext {
	var p = new(IsnullContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_isnull
	return p
}

func (*IsnullContext) IsIsnullContext() {}

func NewIsnullContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsnullContext {
	var p = new(IsnullContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_isnull

	return p
}

func (s *IsnullContext) GetParser() antlr.Parser { return s.parser }

func (s *IsnullContext) ISNULL() antlr.TerminalNode {
	return s.GetToken(SSQLParserISNULL, 0)
}

func (s *IsnullContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsnullContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsnullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterIsnull(s)
	}
}

func (s *IsnullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitIsnull(s)
	}
}

func (s *IsnullContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitIsnull(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Isnull() (localctx IIsnullContext) {
	this := p
	_ = this

	localctx = NewIsnullContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SSQLParserRULE_isnull)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(282)
		p.Match(SSQLParserISNULL)
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserT__1 {
		{
			p.SetState(283)
			p.Match(SSQLParserT__1)
		}
		{
			p.SetState(284)
			p.Match(SSQLParserT__2)
		}

	}

	return localctx
}

// IBooleanContext is an interface to support dynamic dispatch.
type IBooleanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanContext differentiates from other interfaces.
	IsBooleanContext()
}

type BooleanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanContext() *BooleanContext {
	var p = new(BooleanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_boolean
	return p
}

func (*BooleanContext) IsBooleanContext() {}

func NewBooleanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanContext {
	var p = new(BooleanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_boolean

	return p
}

func (s *BooleanContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SSQLParserTRUE, 0)
}

func (s *BooleanContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SSQLParserFALSE, 0)
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitBoolean(s)
	}
}

func (s *BooleanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitBoolean(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Boolean() (localctx IBooleanContext) {
	this := p
	_ = this

	localctx = NewBooleanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SSQLParserRULE_boolean)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SSQLParserTRUE || _la == SSQLParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IScalarContext is an interface to support dynamic dispatch.
type IScalarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarContext differentiates from other interfaces.
	IsScalarContext()
}

type ScalarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarContext() *ScalarContext {
	var p = new(ScalarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_scalar
	return p
}

func (*ScalarContext) IsScalarContext() {}

func NewScalarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarContext {
	var p = new(ScalarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_scalar

	return p
}

func (s *ScalarContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarContext) REAL_NUMBER() antlr.TerminalNode {
	return s.GetToken(SSQLParserREAL_NUMBER, 0)
}

func (s *ScalarContext) SIGN() antlr.TerminalNode {
	return s.GetToken(SSQLParserSIGN, 0)
}

func (s *ScalarContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserINTEGER)
}

func (s *ScalarContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, i)
}

func (s *ScalarContext) ISO8601_DATE_TIME() antlr.TerminalNode {
	return s.GetToken(SSQLParserISO8601_DATE_TIME, 0)
}

func (s *ScalarContext) STRING() antlr.TerminalNode {
	return s.GetToken(SSQLParserSTRING, 0)
}

func (s *ScalarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterScalar(s)
	}
}

func (s *ScalarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitScalar(s)
	}
}

func (s *ScalarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitScalar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Scalar() (localctx IScalarContext) {
	this := p
	_ = this

	localctx = NewScalarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SSQLParserRULE_scalar)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(303)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SSQLParserSIGN {
			{
				p.SetState(289)
				p.Match(SSQLParserSIGN)
			}

		}
		{
			p.SetState(292)
			p.Match(SSQLParserREAL_NUMBER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SSQLParserSIGN {
			{
				p.SetState(293)
				p.Match(SSQLParserSIGN)
			}

		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SSQLParserINTEGER {
			{
				p.SetState(296)
				p.Match(SSQLParserINTEGER)
			}

			p.SetState(299)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(301)
			p.Match(SSQLParserISO8601_DATE_TIME)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(302)
			p.Match(SSQLParserSTRING)
		}

	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) StringList() IStringListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStringListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStringListContext)
}

func (s *ListContext) DoubleList() IDoubleListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoubleListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDoubleListContext)
}

func (s *ListContext) IntList() IIntListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntListContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) List() (localctx IListContext) {
	this := p
	_ = this

	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SSQLParserRULE_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(305)
			p.StringList()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(306)
			p.DoubleList()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(307)
			p.IntList()
		}

	}

	return localctx
}

// IStringListContext is an interface to support dynamic dispatch.
type IStringListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringListContext differentiates from other interfaces.
	IsStringListContext()
}

type StringListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringListContext() *StringListContext {
	var p = new(StringListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_stringList
	return p
}

func (*StringListContext) IsStringListContext() {}

func NewStringListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringListContext {
	var p = new(StringListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_stringList

	return p
}

func (s *StringListContext) GetParser() antlr.Parser { return s.parser }

func (s *StringListContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserSTRING)
}

func (s *StringListContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserSTRING, i)
}

func (s *StringListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterStringList(s)
	}
}

func (s *StringListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitStringList(s)
	}
}

func (s *StringListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitStringList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) StringList() (localctx IStringListContext) {
	this := p
	_ = this

	localctx = NewStringListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SSQLParserRULE_stringList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(SSQLParserSTRING)
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SSQLParserT__0 {
		{
			p.SetState(311)
			p.Match(SSQLParserT__0)
		}
		{
			p.SetState(312)
			p.Match(SSQLParserSTRING)
		}

		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDoubleListContext is an interface to support dynamic dispatch.
type IDoubleListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoubleListContext differentiates from other interfaces.
	IsDoubleListContext()
}

type DoubleListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoubleListContext() *DoubleListContext {
	var p = new(DoubleListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_doubleList
	return p
}

func (*DoubleListContext) IsDoubleListContext() {}

func NewDoubleListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoubleListContext {
	var p = new(DoubleListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_doubleList

	return p
}

func (s *DoubleListContext) GetParser() antlr.Parser { return s.parser }

func (s *DoubleListContext) AllREAL_NUMBER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserREAL_NUMBER)
}

func (s *DoubleListContext) REAL_NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserREAL_NUMBER, i)
}

func (s *DoubleListContext) AllSIGN() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserSIGN)
}

func (s *DoubleListContext) SIGN(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserSIGN, i)
}

func (s *DoubleListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DoubleListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterDoubleList(s)
	}
}

func (s *DoubleListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitDoubleList(s)
	}
}

func (s *DoubleListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitDoubleList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) DoubleList() (localctx IDoubleListContext) {
	this := p
	_ = this

	localctx = NewDoubleListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SSQLParserRULE_doubleList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserSIGN {
		{
			p.SetState(318)
			p.Match(SSQLParserSIGN)
		}

	}
	{
		p.SetState(321)
		p.Match(SSQLParserREAL_NUMBER)
	}
	p.SetState(329)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SSQLParserT__0 {
		{
			p.SetState(322)
			p.Match(SSQLParserT__0)
		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SSQLParserSIGN {
			{
				p.SetState(323)
				p.Match(SSQLParserSIGN)
			}

		}
		{
			p.SetState(326)
			p.Match(SSQLParserREAL_NUMBER)
		}

		p.SetState(331)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIntListContext is an interface to support dynamic dispatch.
type IIntListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntListContext differentiates from other interfaces.
	IsIntListContext()
}

type IntListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntListContext() *IntListContext {
	var p = new(IntListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_intList
	return p
}

func (*IntListContext) IsIntListContext() {}

func NewIntListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntListContext {
	var p = new(IntListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_intList

	return p
}

func (s *IntListContext) GetParser() antlr.Parser { return s.parser }

func (s *IntListContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserINTEGER)
}

func (s *IntListContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, i)
}

func (s *IntListContext) AllSIGN() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserSIGN)
}

func (s *IntListContext) SIGN(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserSIGN, i)
}

func (s *IntListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterIntList(s)
	}
}

func (s *IntListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitIntList(s)
	}
}

func (s *IntListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitIntList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) IntList() (localctx IIntListContext) {
	this := p
	_ = this

	localctx = NewIntListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SSQLParserRULE_intList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserSIGN {
		{
			p.SetState(332)
			p.Match(SSQLParserSIGN)
		}

	}
	{
		p.SetState(335)
		p.Match(SSQLParserINTEGER)
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SSQLParserT__0 {
		{
			p.SetState(336)
			p.Match(SSQLParserT__0)
		}
		p.SetState(338)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SSQLParserSIGN {
			{
				p.SetState(337)
				p.Match(SSQLParserSIGN)
			}

		}
		{
			p.SetState(340)
			p.Match(SSQLParserINTEGER)
		}

		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrderByContext is an interface to support dynamic dispatch.
type IOrderByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByContext differentiates from other interfaces.
	IsOrderByContext()
}

type OrderByContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByContext() *OrderByContext {
	var p = new(OrderByContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_orderBy
	return p
}

func (*OrderByContext) IsOrderByContext() {}

func NewOrderByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByContext {
	var p = new(OrderByContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_orderBy

	return p
}

func (s *OrderByContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByContext) ORDER_BY() antlr.TerminalNode {
	return s.GetToken(SSQLParserORDER_BY, 0)
}

func (s *OrderByContext) AllOrder() []IOrderContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrderContext); ok {
			len++
		}
	}

	tst := make([]IOrderContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrderContext); ok {
			tst[i] = t.(IOrderContext)
			i++
		}
	}

	return tst
}

func (s *OrderByContext) Order(i int) IOrderContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderContext)
}

func (s *OrderByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterOrderBy(s)
	}
}

func (s *OrderByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitOrderBy(s)
	}
}

func (s *OrderByContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitOrderBy(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) OrderBy() (localctx IOrderByContext) {
	this := p
	_ = this

	localctx = NewOrderByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SSQLParserRULE_orderBy)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		p.Match(SSQLParserORDER_BY)
	}
	{
		p.SetState(347)
		p.Order()
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SSQLParserT__0 {
		{
			p.SetState(348)
			p.Match(SSQLParserT__0)
		}
		{
			p.SetState(349)
			p.Order()
		}

		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrderContext is an interface to support dynamic dispatch.
type IOrderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDir returns the dir token.
	GetDir() antlr.Token

	// SetDir sets the dir token.
	SetDir(antlr.Token)

	// IsOrderContext differentiates from other interfaces.
	IsOrderContext()
}

type OrderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dir    antlr.Token
}

func NewEmptyOrderContext() *OrderContext {
	var p = new(OrderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_order
	return p
}

func (*OrderContext) IsOrderContext() {}

func NewOrderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderContext {
	var p = new(OrderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_order

	return p
}

func (s *OrderContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderContext) GetDir() antlr.Token { return s.dir }

func (s *OrderContext) SetDir(v antlr.Token) { s.dir = v }

func (s *OrderContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SSQLParserIDENTIFIER, 0)
}

func (s *OrderContext) ASC() antlr.TerminalNode {
	return s.GetToken(SSQLParserASC, 0)
}

func (s *OrderContext) DESC() antlr.TerminalNode {
	return s.GetToken(SSQLParserDESC, 0)
}

func (s *OrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterOrder(s)
	}
}

func (s *OrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitOrder(s)
	}
}

func (s *OrderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitOrder(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Order() (localctx IOrderContext) {
	this := p
	_ = this

	localctx = NewOrderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SSQLParserRULE_order)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(SSQLParserIDENTIFIER)
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SSQLParserASC || _la == SSQLParserDESC {
		{
			p.SetState(356)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*OrderContext).dir = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SSQLParserASC || _la == SSQLParserDESC) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*OrderContext).dir = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// ILimitContext is an interface to support dynamic dispatch.
type ILimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitContext differentiates from other interfaces.
	IsLimitContext()
}

type LimitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitContext() *LimitContext {
	var p = new(LimitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SSQLParserRULE_limit
	return p
}

func (*LimitContext) IsLimitContext() {}

func NewLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitContext {
	var p = new(LimitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SSQLParserRULE_limit

	return p
}

func (s *LimitContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SSQLParserLIMIT, 0)
}

func (s *LimitContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(SSQLParserINTEGER)
}

func (s *LimitContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(SSQLParserINTEGER, i)
}

func (s *LimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.EnterLimit(s)
	}
}

func (s *LimitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SSQLListener); ok {
		listenerT.ExitLimit(s)
	}
}

func (s *LimitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SSQLVisitor:
		return t.VisitLimit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SSQLParser) Limit() (localctx ILimitContext) {
	this := p
	_ = this

	localctx = NewLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SSQLParserRULE_limit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(359)
		p.Match(SSQLParserLIMIT)
	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SSQLParserINTEGER {
		{
			p.SetState(360)
			p.Match(SSQLParserINTEGER)
		}

		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
