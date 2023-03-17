/*
 * Copyright 2022 Rock Lei Wang
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Package parser declares an expression parser with support for macro
 * expansion.
 */

grammar SSQL;

// Grammar Rules
// =============

start
    : FIND selection (FROM from)? WHERE expression orderBy? limit? EOF
    ;

selection
    : attribute (',' attribute)*
    | GROUP_BY '(' groupBy (',' groupBy)* ')' (',' aggregate)*
    ;

attribute
    : IDENTIFIER | aggregate
    ;

aggregate
    : (AVG | MAX | MIN | SUM | COUNT) '(' IDENTIFIER ')' | percentile
    ;

percentile
    : PERCENTILE '(' IDENTIFIER ',' REAL_NUMBER ')'
    ;

groupBy
    : IDENTIFIER | partition
    ;

partition
    : PARTITION '(' IDENTIFIER ',' INTEGER ')'
    ;

from
    : LOCAL? NAME
    ;

expression
    : tuple tuple*
    ;

tuple
    : vector | or | and
    ;

vector
    : '[' IDENTIFIER? PATH? (predicate | vector+)? ']'
    ;

or
    : '{' tuple+ '}'
    ;

and
    : '{&' tuple+ '}'
    ;

predicate
    : eq | neq | gt | ge | lt | le | in | between | contain | exist | isnull | boolean
    ;

eq
    : EQ '(' scalar ')'
    ;

neq
    : NEQ '(' scalar ')'
    ;

gt
    : GT '(' scalar ')'
    ;

ge
    : GE '(' scalar ')'
    ;

lt
    : LT '(' scalar ')'
    ;

le
    : LE '(' scalar ')'
    ;

in
    : IN '(' list ')'
    ;

between
    : BETWEEN '(' SIGN? INTEGER ',' SIGN? INTEGER ')'
    | BETWEEN '(' SIGN? REAL_NUMBER ',' SIGN? REAL_NUMBER ')'
    | BETWEEN '(' ISO8601_DATE_TIME ',' ISO8601_DATE_TIME ')'
    ;

contain
    : CONTAIN '(' STRING ')'
    ;

exist
    : EXIST ('(' ')')?
    ;

isnull
    : ISNULL ('(' ')')?
    ;

boolean
    : TRUE | FALSE
    ;

scalar
    : SIGN? REAL_NUMBER
    | SIGN? INTEGER+
    | ISO8601_DATE_TIME
    | STRING
    ;

list
    : stringList | doubleList | intList
    ;

stringList
    : STRING (',' STRING)*
    ;

doubleList
    : SIGN? REAL_NUMBER (',' SIGN? REAL_NUMBER)*
    ;

intList
    : SIGN? INTEGER (',' SIGN? INTEGER)*
    ;

orderBy
    : ORDER_BY order (',' order)*
    ;

order
    : IDENTIFIER dir=(ASC | DESC)?
    ;

limit
    : LIMIT INTEGER+
    ;

// Lexer Rules
// ===========
AVG : A V G;
MAX : M A X;
MIN : M I N;
SUM : S U M;
COUNT : C O U N T;
PERCENTILE : P C T L;
PARTITION : P A R T;

EQ : E Q;
NEQ : N E Q;
IN: I N;
LT : L T;
LE : L E;
GE : G E;
GT : G T;
BETWEEN : B E T W E E N;
CONTAIN : C O N T A I N;
EXIST : E X I S T;
ISNULL : I S N U L L;
TRUE: T R U E;
FALSE: F A L S E;

FIND : F I N D;
FROM : F R O M;
WHERE : W H E R E;
ORDER_BY : O R D E R MINUS B Y;
GROUP_BY : G R O U P MINUS B Y;
LIMIT : L I M I T;
ASC : A S C;
DESC : D E S C;
LOCAL : L O C A L;
NAME : (LETTER | '_') (LETTER | DIGIT | '_' | '.' | '-')*;
PATH : '/' | '/' NAME ('/' NAME)*;

STRING : DQUOTA_STRING | SQUOTA_STRING | BQUOTA_STRING;
INTEGER : '0'+ | NON_ZERO_DIGIT DIGIT*;
REAL_NUMBER
    : (DIGIT+)? '.' DIGIT+
    | DIGIT+ '.' EXPONENT
    | (DIGIT+)? '.' (DIGIT+ EXPONENT)
    | DIGIT+ EXPONENT
    ;
ISO8601_DATE_TIME
    : YEAR '-' MONTH '-' DAY ('T' HOUR ':' MINUTE ':' SECOND ('.' MICROSECOND)? TIMEZONE?)?
    ;
SIGN : ('+'|'-');

fragment LETTER : 'A'..'Z' | 'a'..'z';
fragment NON_ZERO_DIGIT : '1'..'9';
fragment DIGIT : '0'..'9';
fragment EXPONENT : 'E' [-+]? DIGIT+;
fragment SINGLE_QUOTE : '\'';
fragment DOUBLE_QUOTE : '"';
fragment BACK_QUOTE : '`';
fragment MICROSECOND: [0-9] [0-9] [0-9];
fragment TIMEZONE: 'Z' | [+-] HOUR ( ':'? MINUTE )?;  // hour offset, e.g. `+01:20`, or else literal `Z` indicating +0000.
fragment YEAR: [0-9] [0-9] [0-9] [0-9]; // Year in ISO8601:2006 is 4 digits with 0-filling as needed
fragment MONTH: ( [0] [1-9] | [1] [0-2] );  // month in year
fragment DAY: ( [0] [1-9] | [12] [0-9] | [3] [0-1] ); // day in month
fragment HOUR: ( [0] [0-9] | [1] [0-9] | [2] [0-3] ); // hour in 24 hour clock
fragment MINUTE: ([0] [0-9] | [1-5] [0-9]); // minutes
fragment SECOND: ([0] [0-9] | [1-5] [0-9]); // seconds
fragment DQUOTA_STRING: DOUBLE_QUOTE ( '\\'. | '""' | ~('"'| '\\') )* DOUBLE_QUOTE;
fragment SQUOTA_STRING: SINGLE_QUOTE ('\\'. | '\'\'' | ~('\'' | '\\'))* SINGLE_QUOTE;
fragment BQUOTA_STRING: BACK_QUOTE ( '\\'. | '``' | ~('`'|'\\'))* BACK_QUOTE;

fragment A: [aA];
fragment B: [bB];
fragment C: [cC];
fragment D: [dD];
fragment E: [eE];
fragment F: [fF];
fragment G: [gG];
fragment H: [hH];
fragment I: [iI];
fragment J: [jJ];
fragment K: [kK];
fragment L: [lL];
fragment M: [mM];
fragment N: [nN];
fragment O: [oO];
fragment P: [pP];
fragment Q: [qQ];
fragment R: [rR];
fragment S: [sS];
fragment T: [tT];
fragment U: [uU];
fragment V: [vV];
fragment W: [wW];
fragment X: [xX];
fragment Y: [yY];
fragment Z: [zZ];
fragment MINUS: '-';

IDENTIFIER : '$' NAME;
WS : [ \n\t\r] -> skip;