// Generated from /home/rleiwang/workspace/proto/ssql/go/antlr4/SSQL.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SSQLLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, AVG=9, 
		MAX=10, MIN=11, SUM=12, COUNT=13, PERCENTILE=14, PARTITION=15, EQ=16, 
		NEQ=17, IN=18, LT=19, LE=20, GE=21, GT=22, BETWEEN=23, CONTAIN=24, EXIST=25, 
		TIMEFRAME=26, KEY=27, FIND=28, WHERE=29, ORDER_BY=30, GROUP_BY=31, LIMIT=32, 
		ASC=33, DESC=34, NAME=35, PATH=36, STRING=37, INTEGER=38, REAL_NUMBER=39, 
		IDENTIFIER=40, WS=41;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "AVG", 
			"MAX", "MIN", "SUM", "COUNT", "PERCENTILE", "PARTITION", "EQ", "NEQ", 
			"IN", "LT", "LE", "GE", "GT", "BETWEEN", "CONTAIN", "EXIST", "TIMEFRAME", 
			"KEY", "FIND", "WHERE", "ORDER_BY", "GROUP_BY", "LIMIT", "ASC", "DESC", 
			"NAME", "PATH", "STRING", "INTEGER", "REAL_NUMBER", "LETTER", "NON_ZERO_DIGIT", 
			"DIGIT", "EXPONENT", "DQUOTA_STRING", "SQUOTA_STRING", "BQUOTA_STRING", 
			"IDENTIFIER", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "','", "'('", "')'", "'['", "']'", "'{'", "'}'", "'{&'", "'AVG'", 
			"'MAX'", "'MIN'", "'SUM'", "'COUNT'", "'PERCENTILE'", "'PARTITION'", 
			"'EQ'", "'NEQ'", "'IN'", "'LT'", "'LE'", "'GE'", "'GT'", "'BETWEEN'", 
			"'CONTAIN'", "'EXIST'", "'TIMEFRAME'", "'KEY'", "'FIND'", "'WHERE'", 
			"'ORDER-BY'", "'GROUP-BY'", "'LIMIT'", "'ASC'", "'DESC'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, "AVG", "MAX", "MIN", 
			"SUM", "COUNT", "PERCENTILE", "PARTITION", "EQ", "NEQ", "IN", "LT", "LE", 
			"GE", "GT", "BETWEEN", "CONTAIN", "EXIST", "TIMEFRAME", "KEY", "FIND", 
			"WHERE", "ORDER_BY", "GROUP_BY", "LIMIT", "ASC", "DESC", "NAME", "PATH", 
			"STRING", "INTEGER", "REAL_NUMBER", "IDENTIFIER", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public SSQLLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "SSQL.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2+\u019a\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t+\4"+
		",\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\3\2\3\2\3\3\3\3\3\4\3\4\3\5"+
		"\3\5\3\6\3\6\3\7\3\7\3\b\3\b\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\13\3\13\3\13"+
		"\3\13\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3"+
		"\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3"+
		"\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3"+
		"\23\3\23\3\23\3\24\3\24\3\24\3\25\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3"+
		"\27\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3"+
		"\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32\3\33\3\33\3\33\3\33\3\33\3"+
		"\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\35\3\35\3\35\3\35\3\35\3"+
		"\36\3\36\3\36\3\36\3\36\3\36\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3\37\3"+
		"\37\3 \3 \3 \3 \3 \3 \3 \3 \3 \3!\3!\3!\3!\3!\3!\3\"\3\"\3\"\3\"\3#\3"+
		"#\3#\3#\3#\3$\3$\5$\u0108\n$\3$\3$\3$\7$\u010d\n$\f$\16$\u0110\13$\3%"+
		"\3%\3%\3%\3%\7%\u0117\n%\f%\16%\u011a\13%\5%\u011c\n%\3&\3&\3&\5&\u0121"+
		"\n&\3\'\6\'\u0124\n\'\r\'\16\'\u0125\3\'\3\'\7\'\u012a\n\'\f\'\16\'\u012d"+
		"\13\'\5\'\u012f\n\'\3(\6(\u0132\n(\r(\16(\u0133\5(\u0136\n(\3(\3(\6(\u013a"+
		"\n(\r(\16(\u013b\3(\6(\u013f\n(\r(\16(\u0140\3(\3(\3(\3(\6(\u0147\n(\r"+
		"(\16(\u0148\5(\u014b\n(\3(\3(\6(\u014f\n(\r(\16(\u0150\3(\3(\3(\6(\u0156"+
		"\n(\r(\16(\u0157\3(\3(\5(\u015c\n(\3)\3)\3*\3*\3+\3+\3,\3,\5,\u0166\n"+
		",\3,\6,\u0169\n,\r,\16,\u016a\3-\3-\3-\3-\3-\3-\7-\u0173\n-\f-\16-\u0176"+
		"\13-\3-\3-\3.\3.\3.\3.\3.\3.\7.\u0180\n.\f.\16.\u0183\13.\3.\3.\3/\3/"+
		"\3/\3/\3/\3/\7/\u018d\n/\f/\16/\u0190\13/\3/\3/\3\60\3\60\3\60\3\61\3"+
		"\61\3\61\3\61\2\2\62\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27"+
		"\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33"+
		"\65\34\67\359\36;\37= ?!A\"C#E$G%I&K\'M(O)Q\2S\2U\2W\2Y\2[\2]\2_*a+\3"+
		"\2\t\4\2/\60aa\4\2C\\c|\4\2--//\4\2$$^^\4\2))^^\4\2^^bb\5\2\13\f\17\17"+
		"\"\"\2\u01b3\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2"+
		"\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2"+
		"\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2"+
		"\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2"+
		"\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3"+
		"\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\2E\3\2\2"+
		"\2\2G\3\2\2\2\2I\3\2\2\2\2K\3\2\2\2\2M\3\2\2\2\2O\3\2\2\2\2_\3\2\2\2\2"+
		"a\3\2\2\2\3c\3\2\2\2\5e\3\2\2\2\7g\3\2\2\2\ti\3\2\2\2\13k\3\2\2\2\rm\3"+
		"\2\2\2\17o\3\2\2\2\21q\3\2\2\2\23t\3\2\2\2\25x\3\2\2\2\27|\3\2\2\2\31"+
		"\u0080\3\2\2\2\33\u0084\3\2\2\2\35\u008a\3\2\2\2\37\u0095\3\2\2\2!\u009f"+
		"\3\2\2\2#\u00a2\3\2\2\2%\u00a6\3\2\2\2\'\u00a9\3\2\2\2)\u00ac\3\2\2\2"+
		"+\u00af\3\2\2\2-\u00b2\3\2\2\2/\u00b5\3\2\2\2\61\u00bd\3\2\2\2\63\u00c5"+
		"\3\2\2\2\65\u00cb\3\2\2\2\67\u00d5\3\2\2\29\u00d9\3\2\2\2;\u00de\3\2\2"+
		"\2=\u00e4\3\2\2\2?\u00ed\3\2\2\2A\u00f6\3\2\2\2C\u00fc\3\2\2\2E\u0100"+
		"\3\2\2\2G\u0107\3\2\2\2I\u011b\3\2\2\2K\u0120\3\2\2\2M\u012e\3\2\2\2O"+
		"\u015b\3\2\2\2Q\u015d\3\2\2\2S\u015f\3\2\2\2U\u0161\3\2\2\2W\u0163\3\2"+
		"\2\2Y\u016c\3\2\2\2[\u0179\3\2\2\2]\u0186\3\2\2\2_\u0193\3\2\2\2a\u0196"+
		"\3\2\2\2cd\7.\2\2d\4\3\2\2\2ef\7*\2\2f\6\3\2\2\2gh\7+\2\2h\b\3\2\2\2i"+
		"j\7]\2\2j\n\3\2\2\2kl\7_\2\2l\f\3\2\2\2mn\7}\2\2n\16\3\2\2\2op\7\177\2"+
		"\2p\20\3\2\2\2qr\7}\2\2rs\7(\2\2s\22\3\2\2\2tu\7C\2\2uv\7X\2\2vw\7I\2"+
		"\2w\24\3\2\2\2xy\7O\2\2yz\7C\2\2z{\7Z\2\2{\26\3\2\2\2|}\7O\2\2}~\7K\2"+
		"\2~\177\7P\2\2\177\30\3\2\2\2\u0080\u0081\7U\2\2\u0081\u0082\7W\2\2\u0082"+
		"\u0083\7O\2\2\u0083\32\3\2\2\2\u0084\u0085\7E\2\2\u0085\u0086\7Q\2\2\u0086"+
		"\u0087\7W\2\2\u0087\u0088\7P\2\2\u0088\u0089\7V\2\2\u0089\34\3\2\2\2\u008a"+
		"\u008b\7R\2\2\u008b\u008c\7G\2\2\u008c\u008d\7T\2\2\u008d\u008e\7E\2\2"+
		"\u008e\u008f\7G\2\2\u008f\u0090\7P\2\2\u0090\u0091\7V\2\2\u0091\u0092"+
		"\7K\2\2\u0092\u0093\7N\2\2\u0093\u0094\7G\2\2\u0094\36\3\2\2\2\u0095\u0096"+
		"\7R\2\2\u0096\u0097\7C\2\2\u0097\u0098\7T\2\2\u0098\u0099\7V\2\2\u0099"+
		"\u009a\7K\2\2\u009a\u009b\7V\2\2\u009b\u009c\7K\2\2\u009c\u009d\7Q\2\2"+
		"\u009d\u009e\7P\2\2\u009e \3\2\2\2\u009f\u00a0\7G\2\2\u00a0\u00a1\7S\2"+
		"\2\u00a1\"\3\2\2\2\u00a2\u00a3\7P\2\2\u00a3\u00a4\7G\2\2\u00a4\u00a5\7"+
		"S\2\2\u00a5$\3\2\2\2\u00a6\u00a7\7K\2\2\u00a7\u00a8\7P\2\2\u00a8&\3\2"+
		"\2\2\u00a9\u00aa\7N\2\2\u00aa\u00ab\7V\2\2\u00ab(\3\2\2\2\u00ac\u00ad"+
		"\7N\2\2\u00ad\u00ae\7G\2\2\u00ae*\3\2\2\2\u00af\u00b0\7I\2\2\u00b0\u00b1"+
		"\7G\2\2\u00b1,\3\2\2\2\u00b2\u00b3\7I\2\2\u00b3\u00b4\7V\2\2\u00b4.\3"+
		"\2\2\2\u00b5\u00b6\7D\2\2\u00b6\u00b7\7G\2\2\u00b7\u00b8\7V\2\2\u00b8"+
		"\u00b9\7Y\2\2\u00b9\u00ba\7G\2\2\u00ba\u00bb\7G\2\2\u00bb\u00bc\7P\2\2"+
		"\u00bc\60\3\2\2\2\u00bd\u00be\7E\2\2\u00be\u00bf\7Q\2\2\u00bf\u00c0\7"+
		"P\2\2\u00c0\u00c1\7V\2\2\u00c1\u00c2\7C\2\2\u00c2\u00c3\7K\2\2\u00c3\u00c4"+
		"\7P\2\2\u00c4\62\3\2\2\2\u00c5\u00c6\7G\2\2\u00c6\u00c7\7Z\2\2\u00c7\u00c8"+
		"\7K\2\2\u00c8\u00c9\7U\2\2\u00c9\u00ca\7V\2\2\u00ca\64\3\2\2\2\u00cb\u00cc"+
		"\7V\2\2\u00cc\u00cd\7K\2\2\u00cd\u00ce\7O\2\2\u00ce\u00cf\7G\2\2\u00cf"+
		"\u00d0\7H\2\2\u00d0\u00d1\7T\2\2\u00d1\u00d2\7C\2\2\u00d2\u00d3\7O\2\2"+
		"\u00d3\u00d4\7G\2\2\u00d4\66\3\2\2\2\u00d5\u00d6\7M\2\2\u00d6\u00d7\7"+
		"G\2\2\u00d7\u00d8\7[\2\2\u00d88\3\2\2\2\u00d9\u00da\7H\2\2\u00da\u00db"+
		"\7K\2\2\u00db\u00dc\7P\2\2\u00dc\u00dd\7F\2\2\u00dd:\3\2\2\2\u00de\u00df"+
		"\7Y\2\2\u00df\u00e0\7J\2\2\u00e0\u00e1\7G\2\2\u00e1\u00e2\7T\2\2\u00e2"+
		"\u00e3\7G\2\2\u00e3<\3\2\2\2\u00e4\u00e5\7Q\2\2\u00e5\u00e6\7T\2\2\u00e6"+
		"\u00e7\7F\2\2\u00e7\u00e8\7G\2\2\u00e8\u00e9\7T\2\2\u00e9\u00ea\7/\2\2"+
		"\u00ea\u00eb\7D\2\2\u00eb\u00ec\7[\2\2\u00ec>\3\2\2\2\u00ed\u00ee\7I\2"+
		"\2\u00ee\u00ef\7T\2\2\u00ef\u00f0\7Q\2\2\u00f0\u00f1\7W\2\2\u00f1\u00f2"+
		"\7R\2\2\u00f2\u00f3\7/\2\2\u00f3\u00f4\7D\2\2\u00f4\u00f5\7[\2\2\u00f5"+
		"@\3\2\2\2\u00f6\u00f7\7N\2\2\u00f7\u00f8\7K\2\2\u00f8\u00f9\7O\2\2\u00f9"+
		"\u00fa\7K\2\2\u00fa\u00fb\7V\2\2\u00fbB\3\2\2\2\u00fc\u00fd\7C\2\2\u00fd"+
		"\u00fe\7U\2\2\u00fe\u00ff\7E\2\2\u00ffD\3\2\2\2\u0100\u0101\7F\2\2\u0101"+
		"\u0102\7G\2\2\u0102\u0103\7U\2\2\u0103\u0104\7E\2\2\u0104F\3\2\2\2\u0105"+
		"\u0108\5Q)\2\u0106\u0108\7a\2\2\u0107\u0105\3\2\2\2\u0107\u0106\3\2\2"+
		"\2\u0108\u010e\3\2\2\2\u0109\u010d\5Q)\2\u010a\u010d\5U+\2\u010b\u010d"+
		"\t\2\2\2\u010c\u0109\3\2\2\2\u010c\u010a\3\2\2\2\u010c\u010b\3\2\2\2\u010d"+
		"\u0110\3\2\2\2\u010e\u010c\3\2\2\2\u010e\u010f\3\2\2\2\u010fH\3\2\2\2"+
		"\u0110\u010e\3\2\2\2\u0111\u011c\7\61\2\2\u0112\u0113\7\61\2\2\u0113\u0118"+
		"\5G$\2\u0114\u0115\7\61\2\2\u0115\u0117\5G$\2\u0116\u0114\3\2\2\2\u0117"+
		"\u011a\3\2\2\2\u0118\u0116\3\2\2\2\u0118\u0119\3\2\2\2\u0119\u011c\3\2"+
		"\2\2\u011a\u0118\3\2\2\2\u011b\u0111\3\2\2\2\u011b\u0112\3\2\2\2\u011c"+
		"J\3\2\2\2\u011d\u0121\5Y-\2\u011e\u0121\5[.\2\u011f\u0121\5]/\2\u0120"+
		"\u011d\3\2\2\2\u0120\u011e\3\2\2\2\u0120\u011f\3\2\2\2\u0121L\3\2\2\2"+
		"\u0122\u0124\7\62\2\2\u0123\u0122\3\2\2\2\u0124\u0125\3\2\2\2\u0125\u0123"+
		"\3\2\2\2\u0125\u0126\3\2\2\2\u0126\u012f\3\2\2\2\u0127\u012b\5S*\2\u0128"+
		"\u012a\5U+\2\u0129\u0128\3\2\2\2\u012a\u012d\3\2\2\2\u012b\u0129\3\2\2"+
		"\2\u012b\u012c\3\2\2\2\u012c\u012f\3\2\2\2\u012d\u012b\3\2\2\2\u012e\u0123"+
		"\3\2\2\2\u012e\u0127\3\2\2\2\u012fN\3\2\2\2\u0130\u0132\5U+\2\u0131\u0130"+
		"\3\2\2\2\u0132\u0133\3\2\2\2\u0133\u0131\3\2\2\2\u0133\u0134\3\2\2\2\u0134"+
		"\u0136\3\2\2\2\u0135\u0131\3\2\2\2\u0135\u0136\3\2\2\2\u0136\u0137\3\2"+
		"\2\2\u0137\u0139\7\60\2\2\u0138\u013a\5U+\2\u0139\u0138\3\2\2\2\u013a"+
		"\u013b\3\2\2\2\u013b\u0139\3\2\2\2\u013b\u013c\3\2\2\2\u013c\u015c\3\2"+
		"\2\2\u013d\u013f\5U+\2\u013e\u013d\3\2\2\2\u013f\u0140\3\2\2\2\u0140\u013e"+
		"\3\2\2\2\u0140\u0141\3\2\2\2\u0141\u0142\3\2\2\2\u0142\u0143\7\60\2\2"+
		"\u0143\u0144\5W,\2\u0144\u015c\3\2\2\2\u0145\u0147\5U+\2\u0146\u0145\3"+
		"\2\2\2\u0147\u0148\3\2\2\2\u0148\u0146\3\2\2\2\u0148\u0149\3\2\2\2\u0149"+
		"\u014b\3\2\2\2\u014a\u0146\3\2\2\2\u014a\u014b\3\2\2\2\u014b\u014c\3\2"+
		"\2\2\u014c\u014e\7\60\2\2\u014d\u014f\5U+\2\u014e\u014d\3\2\2\2\u014f"+
		"\u0150\3\2\2\2\u0150\u014e\3\2\2\2\u0150\u0151\3\2\2\2\u0151\u0152\3\2"+
		"\2\2\u0152\u0153\5W,\2\u0153\u015c\3\2\2\2\u0154\u0156\5U+\2\u0155\u0154"+
		"\3\2\2\2\u0156\u0157\3\2\2\2\u0157\u0155\3\2\2\2\u0157\u0158\3\2\2\2\u0158"+
		"\u0159\3\2\2\2\u0159\u015a\5W,\2\u015a\u015c\3\2\2\2\u015b\u0135\3\2\2"+
		"\2\u015b\u013e\3\2\2\2\u015b\u014a\3\2\2\2\u015b\u0155\3\2\2\2\u015cP"+
		"\3\2\2\2\u015d\u015e\t\3\2\2\u015eR\3\2\2\2\u015f\u0160\4\63;\2\u0160"+
		"T\3\2\2\2\u0161\u0162\4\62;\2\u0162V\3\2\2\2\u0163\u0165\7G\2\2\u0164"+
		"\u0166\t\4\2\2\u0165\u0164\3\2\2\2\u0165\u0166\3\2\2\2\u0166\u0168\3\2"+
		"\2\2\u0167\u0169\5U+\2\u0168\u0167\3\2\2\2\u0169\u016a\3\2\2\2\u016a\u0168"+
		"\3\2\2\2\u016a\u016b\3\2\2\2\u016bX\3\2\2\2\u016c\u0174\7$\2\2\u016d\u016e"+
		"\7^\2\2\u016e\u0173\13\2\2\2\u016f\u0170\7$\2\2\u0170\u0173\7$\2\2\u0171"+
		"\u0173\n\5\2\2\u0172\u016d\3\2\2\2\u0172\u016f\3\2\2\2\u0172\u0171\3\2"+
		"\2\2\u0173\u0176\3\2\2\2\u0174\u0172\3\2\2\2\u0174\u0175\3\2\2\2\u0175"+
		"\u0177\3\2\2\2\u0176\u0174\3\2\2\2\u0177\u0178\7$\2\2\u0178Z\3\2\2\2\u0179"+
		"\u0181\7)\2\2\u017a\u017b\7^\2\2\u017b\u0180\13\2\2\2\u017c\u017d\7)\2"+
		"\2\u017d\u0180\7)\2\2\u017e\u0180\n\6\2\2\u017f\u017a\3\2\2\2\u017f\u017c"+
		"\3\2\2\2\u017f\u017e\3\2\2\2\u0180\u0183\3\2\2\2\u0181\u017f\3\2\2\2\u0181"+
		"\u0182\3\2\2\2\u0182\u0184\3\2\2\2\u0183\u0181\3\2\2\2\u0184\u0185\7)"+
		"\2\2\u0185\\\3\2\2\2\u0186\u018e\7b\2\2\u0187\u0188\7^\2\2\u0188\u018d"+
		"\13\2\2\2\u0189\u018a\7b\2\2\u018a\u018d\7b\2\2\u018b\u018d\n\7\2\2\u018c"+
		"\u0187\3\2\2\2\u018c\u0189\3\2\2\2\u018c\u018b\3\2\2\2\u018d\u0190\3\2"+
		"\2\2\u018e\u018c\3\2\2\2\u018e\u018f\3\2\2\2\u018f\u0191\3\2\2\2\u0190"+
		"\u018e\3\2\2\2\u0191\u0192\7b\2\2\u0192^\3\2\2\2\u0193\u0194\7&\2\2\u0194"+
		"\u0195\5G$\2\u0195`\3\2\2\2\u0196\u0197\t\b\2\2\u0197\u0198\3\2\2\2\u0198"+
		"\u0199\b\61\2\2\u0199b\3\2\2\2\35\2\u0107\u010c\u010e\u0118\u011b\u0120"+
		"\u0125\u012b\u012e\u0133\u0135\u013b\u0140\u0148\u014a\u0150\u0157\u015b"+
		"\u0165\u016a\u0172\u0174\u017f\u0181\u018c\u018e\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}