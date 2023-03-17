// Generated from /home/rleiwang/workspace/proto/ssql/go/antlr4/SSQL.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class SSQLParser extends Parser {
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
	public static final int
		RULE_start = 0, RULE_selection = 1, RULE_attribute = 2, RULE_aggregate = 3, 
		RULE_percentile = 4, RULE_groupBy = 5, RULE_partition = 6, RULE_expression = 7, 
		RULE_tuple = 8, RULE_vector = 9, RULE_or = 10, RULE_and = 11, RULE_predicate = 12, 
		RULE_eq = 13, RULE_neq = 14, RULE_gt = 15, RULE_ge = 16, RULE_lt = 17, 
		RULE_le = 18, RULE_in = 19, RULE_between = 20, RULE_contain = 21, RULE_exist = 22, 
		RULE_timeframe = 23, RULE_key = 24, RULE_scalar = 25, RULE_list = 26, 
		RULE_stringList = 27, RULE_doubleList = 28, RULE_intList = 29, RULE_orderBy = 30, 
		RULE_order = 31, RULE_limit = 32;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "selection", "attribute", "aggregate", "percentile", "groupBy", 
			"partition", "expression", "tuple", "vector", "or", "and", "predicate", 
			"eq", "neq", "gt", "ge", "lt", "le", "in", "between", "contain", "exist", 
			"timeframe", "key", "scalar", "list", "stringList", "doubleList", "intList", 
			"orderBy", "order", "limit"
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

	@Override
	public String getGrammarFileName() { return "SSQL.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public SSQLParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public TerminalNode FIND() { return getToken(SSQLParser.FIND, 0); }
		public SelectionContext selection() {
			return getRuleContext(SelectionContext.class,0);
		}
		public TerminalNode WHERE() { return getToken(SSQLParser.WHERE, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode EOF() { return getToken(SSQLParser.EOF, 0); }
		public OrderByContext orderBy() {
			return getRuleContext(OrderByContext.class,0);
		}
		public LimitContext limit() {
			return getRuleContext(LimitContext.class,0);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			match(FIND);
			setState(67);
			selection();
			setState(68);
			match(WHERE);
			setState(69);
			expression();
			setState(71);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ORDER_BY) {
				{
				setState(70);
				orderBy();
				}
			}

			setState(74);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==LIMIT) {
				{
				setState(73);
				limit();
				}
			}

			setState(76);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SelectionContext extends ParserRuleContext {
		public List<AttributeContext> attribute() {
			return getRuleContexts(AttributeContext.class);
		}
		public AttributeContext attribute(int i) {
			return getRuleContext(AttributeContext.class,i);
		}
		public TerminalNode GROUP_BY() { return getToken(SSQLParser.GROUP_BY, 0); }
		public List<GroupByContext> groupBy() {
			return getRuleContexts(GroupByContext.class);
		}
		public GroupByContext groupBy(int i) {
			return getRuleContext(GroupByContext.class,i);
		}
		public List<AggregateContext> aggregate() {
			return getRuleContexts(AggregateContext.class);
		}
		public AggregateContext aggregate(int i) {
			return getRuleContext(AggregateContext.class,i);
		}
		public SelectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_selection; }
	}

	public final SelectionContext selection() throws RecognitionException {
		SelectionContext _localctx = new SelectionContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_selection);
		int _la;
		try {
			setState(104);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case AVG:
			case MAX:
			case MIN:
			case SUM:
			case COUNT:
			case PERCENTILE:
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(78);
				attribute();
				setState(83);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__0) {
					{
					{
					setState(79);
					match(T__0);
					setState(80);
					attribute();
					}
					}
					setState(85);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case GROUP_BY:
				enterOuterAlt(_localctx, 2);
				{
				setState(86);
				match(GROUP_BY);
				setState(87);
				match(T__1);
				setState(88);
				groupBy();
				setState(93);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__0) {
					{
					{
					setState(89);
					match(T__0);
					setState(90);
					groupBy();
					}
					}
					setState(95);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(96);
				match(T__2);
				setState(101);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__0) {
					{
					{
					setState(97);
					match(T__0);
					setState(98);
					aggregate();
					}
					}
					setState(103);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AttributeContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(SSQLParser.IDENTIFIER, 0); }
		public AggregateContext aggregate() {
			return getRuleContext(AggregateContext.class,0);
		}
		public AttributeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attribute; }
	}

	public final AttributeContext attribute() throws RecognitionException {
		AttributeContext _localctx = new AttributeContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_attribute);
		try {
			setState(108);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(106);
				match(IDENTIFIER);
				}
				break;
			case AVG:
			case MAX:
			case MIN:
			case SUM:
			case COUNT:
			case PERCENTILE:
				enterOuterAlt(_localctx, 2);
				{
				setState(107);
				aggregate();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AggregateContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(SSQLParser.IDENTIFIER, 0); }
		public TerminalNode AVG() { return getToken(SSQLParser.AVG, 0); }
		public TerminalNode MAX() { return getToken(SSQLParser.MAX, 0); }
		public TerminalNode MIN() { return getToken(SSQLParser.MIN, 0); }
		public TerminalNode SUM() { return getToken(SSQLParser.SUM, 0); }
		public TerminalNode COUNT() { return getToken(SSQLParser.COUNT, 0); }
		public PercentileContext percentile() {
			return getRuleContext(PercentileContext.class,0);
		}
		public AggregateContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_aggregate; }
	}

	public final AggregateContext aggregate() throws RecognitionException {
		AggregateContext _localctx = new AggregateContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_aggregate);
		int _la;
		try {
			setState(115);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case AVG:
			case MAX:
			case MIN:
			case SUM:
			case COUNT:
				enterOuterAlt(_localctx, 1);
				{
				setState(110);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << AVG) | (1L << MAX) | (1L << MIN) | (1L << SUM) | (1L << COUNT))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(111);
				match(T__1);
				setState(112);
				match(IDENTIFIER);
				setState(113);
				match(T__2);
				}
				break;
			case PERCENTILE:
				enterOuterAlt(_localctx, 2);
				{
				setState(114);
				percentile();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PercentileContext extends ParserRuleContext {
		public TerminalNode PERCENTILE() { return getToken(SSQLParser.PERCENTILE, 0); }
		public TerminalNode IDENTIFIER() { return getToken(SSQLParser.IDENTIFIER, 0); }
		public TerminalNode REAL_NUMBER() { return getToken(SSQLParser.REAL_NUMBER, 0); }
		public PercentileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_percentile; }
	}

	public final PercentileContext percentile() throws RecognitionException {
		PercentileContext _localctx = new PercentileContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_percentile);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(117);
			match(PERCENTILE);
			setState(118);
			match(T__1);
			setState(119);
			match(IDENTIFIER);
			setState(120);
			match(T__0);
			setState(121);
			match(REAL_NUMBER);
			setState(122);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GroupByContext extends ParserRuleContext {
		public TerminalNode IDENTIFIER() { return getToken(SSQLParser.IDENTIFIER, 0); }
		public PartitionContext partition() {
			return getRuleContext(PartitionContext.class,0);
		}
		public GroupByContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_groupBy; }
	}

	public final GroupByContext groupBy() throws RecognitionException {
		GroupByContext _localctx = new GroupByContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_groupBy);
		try {
			setState(126);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IDENTIFIER:
				enterOuterAlt(_localctx, 1);
				{
				setState(124);
				match(IDENTIFIER);
				}
				break;
			case PARTITION:
				enterOuterAlt(_localctx, 2);
				{
				setState(125);
				partition();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PartitionContext extends ParserRuleContext {
		public TerminalNode PARTITION() { return getToken(SSQLParser.PARTITION, 0); }
		public TerminalNode IDENTIFIER() { return getToken(SSQLParser.IDENTIFIER, 0); }
		public TerminalNode INTEGER() { return getToken(SSQLParser.INTEGER, 0); }
		public PartitionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_partition; }
	}

	public final PartitionContext partition() throws RecognitionException {
		PartitionContext _localctx = new PartitionContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_partition);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			match(PARTITION);
			setState(129);
			match(T__1);
			setState(130);
			match(IDENTIFIER);
			setState(131);
			match(T__0);
			setState(132);
			match(INTEGER);
			setState(133);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public List<TupleContext> tuple() {
			return getRuleContexts(TupleContext.class);
		}
		public TupleContext tuple(int i) {
			return getRuleContext(TupleContext.class,i);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_expression);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			tuple();
			setState(139);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__3) | (1L << T__5) | (1L << T__7))) != 0)) {
				{
				{
				setState(136);
				tuple();
				}
				}
				setState(141);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TupleContext extends ParserRuleContext {
		public VectorContext vector() {
			return getRuleContext(VectorContext.class,0);
		}
		public OrContext or() {
			return getRuleContext(OrContext.class,0);
		}
		public AndContext and() {
			return getRuleContext(AndContext.class,0);
		}
		public TupleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tuple; }
	}

	public final TupleContext tuple() throws RecognitionException {
		TupleContext _localctx = new TupleContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_tuple);
		try {
			setState(145);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__3:
				enterOuterAlt(_localctx, 1);
				{
				setState(142);
				vector();
				}
				break;
			case T__5:
				enterOuterAlt(_localctx, 2);
				{
				setState(143);
				or();
				}
				break;
			case T__7:
				enterOuterAlt(_localctx, 3);
				{
				setState(144);
				and();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VectorContext extends ParserRuleContext {
		public TerminalNode PATH() { return getToken(SSQLParser.PATH, 0); }
		public TerminalNode IDENTIFIER() { return getToken(SSQLParser.IDENTIFIER, 0); }
		public PredicateContext predicate() {
			return getRuleContext(PredicateContext.class,0);
		}
		public List<VectorContext> vector() {
			return getRuleContexts(VectorContext.class);
		}
		public VectorContext vector(int i) {
			return getRuleContext(VectorContext.class,i);
		}
		public VectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vector; }
	}

	public final VectorContext vector() throws RecognitionException {
		VectorContext _localctx = new VectorContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_vector);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			match(T__3);
			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==IDENTIFIER) {
				{
				setState(148);
				match(IDENTIFIER);
				}
			}

			setState(151);
			match(PATH);
			setState(158);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case EQ:
			case NEQ:
			case IN:
			case LT:
			case LE:
			case GE:
			case GT:
			case BETWEEN:
			case CONTAIN:
			case EXIST:
			case TIMEFRAME:
			case KEY:
				{
				setState(152);
				predicate();
				}
				break;
			case T__3:
				{
				setState(154); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(153);
					vector();
					}
					}
					setState(156); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==T__3 );
				}
				break;
			case T__4:
				break;
			default:
				break;
			}
			setState(160);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OrContext extends ParserRuleContext {
		public List<TupleContext> tuple() {
			return getRuleContexts(TupleContext.class);
		}
		public TupleContext tuple(int i) {
			return getRuleContext(TupleContext.class,i);
		}
		public OrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_or; }
	}

	public final OrContext or() throws RecognitionException {
		OrContext _localctx = new OrContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_or);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(T__5);
			setState(164); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(163);
				tuple();
				}
				}
				setState(166); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__3) | (1L << T__5) | (1L << T__7))) != 0) );
			setState(168);
			match(T__6);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AndContext extends ParserRuleContext {
		public List<TupleContext> tuple() {
			return getRuleContexts(TupleContext.class);
		}
		public TupleContext tuple(int i) {
			return getRuleContext(TupleContext.class,i);
		}
		public AndContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_and; }
	}

	public final AndContext and() throws RecognitionException {
		AndContext _localctx = new AndContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_and);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(170);
			match(T__7);
			setState(172); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(171);
				tuple();
				}
				}
				setState(174); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( (((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__3) | (1L << T__5) | (1L << T__7))) != 0) );
			setState(176);
			match(T__6);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PredicateContext extends ParserRuleContext {
		public EqContext eq() {
			return getRuleContext(EqContext.class,0);
		}
		public NeqContext neq() {
			return getRuleContext(NeqContext.class,0);
		}
		public GtContext gt() {
			return getRuleContext(GtContext.class,0);
		}
		public GeContext ge() {
			return getRuleContext(GeContext.class,0);
		}
		public LtContext lt() {
			return getRuleContext(LtContext.class,0);
		}
		public LeContext le() {
			return getRuleContext(LeContext.class,0);
		}
		public InContext in() {
			return getRuleContext(InContext.class,0);
		}
		public BetweenContext between() {
			return getRuleContext(BetweenContext.class,0);
		}
		public ContainContext contain() {
			return getRuleContext(ContainContext.class,0);
		}
		public ExistContext exist() {
			return getRuleContext(ExistContext.class,0);
		}
		public TimeframeContext timeframe() {
			return getRuleContext(TimeframeContext.class,0);
		}
		public KeyContext key() {
			return getRuleContext(KeyContext.class,0);
		}
		public PredicateContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_predicate; }
	}

	public final PredicateContext predicate() throws RecognitionException {
		PredicateContext _localctx = new PredicateContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_predicate);
		try {
			setState(190);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case EQ:
				enterOuterAlt(_localctx, 1);
				{
				setState(178);
				eq();
				}
				break;
			case NEQ:
				enterOuterAlt(_localctx, 2);
				{
				setState(179);
				neq();
				}
				break;
			case GT:
				enterOuterAlt(_localctx, 3);
				{
				setState(180);
				gt();
				}
				break;
			case GE:
				enterOuterAlt(_localctx, 4);
				{
				setState(181);
				ge();
				}
				break;
			case LT:
				enterOuterAlt(_localctx, 5);
				{
				setState(182);
				lt();
				}
				break;
			case LE:
				enterOuterAlt(_localctx, 6);
				{
				setState(183);
				le();
				}
				break;
			case IN:
				enterOuterAlt(_localctx, 7);
				{
				setState(184);
				in();
				}
				break;
			case BETWEEN:
				enterOuterAlt(_localctx, 8);
				{
				setState(185);
				between();
				}
				break;
			case CONTAIN:
				enterOuterAlt(_localctx, 9);
				{
				setState(186);
				contain();
				}
				break;
			case EXIST:
				enterOuterAlt(_localctx, 10);
				{
				setState(187);
				exist();
				}
				break;
			case TIMEFRAME:
				enterOuterAlt(_localctx, 11);
				{
				setState(188);
				timeframe();
				}
				break;
			case KEY:
				enterOuterAlt(_localctx, 12);
				{
				setState(189);
				key();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class EqContext extends ParserRuleContext {
		public TerminalNode EQ() { return getToken(SSQLParser.EQ, 0); }
		public ScalarContext scalar() {
			return getRuleContext(ScalarContext.class,0);
		}
		public EqContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_eq; }
	}

	public final EqContext eq() throws RecognitionException {
		EqContext _localctx = new EqContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_eq);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			match(EQ);
			setState(193);
			match(T__1);
			setState(194);
			scalar();
			setState(195);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NeqContext extends ParserRuleContext {
		public TerminalNode NEQ() { return getToken(SSQLParser.NEQ, 0); }
		public ScalarContext scalar() {
			return getRuleContext(ScalarContext.class,0);
		}
		public NeqContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_neq; }
	}

	public final NeqContext neq() throws RecognitionException {
		NeqContext _localctx = new NeqContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_neq);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(197);
			match(NEQ);
			setState(198);
			match(T__1);
			setState(199);
			scalar();
			setState(200);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GtContext extends ParserRuleContext {
		public TerminalNode GT() { return getToken(SSQLParser.GT, 0); }
		public ScalarContext scalar() {
			return getRuleContext(ScalarContext.class,0);
		}
		public GtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_gt; }
	}

	public final GtContext gt() throws RecognitionException {
		GtContext _localctx = new GtContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_gt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(202);
			match(GT);
			setState(203);
			match(T__1);
			setState(204);
			scalar();
			setState(205);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class GeContext extends ParserRuleContext {
		public TerminalNode GE() { return getToken(SSQLParser.GE, 0); }
		public ScalarContext scalar() {
			return getRuleContext(ScalarContext.class,0);
		}
		public GeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ge; }
	}

	public final GeContext ge() throws RecognitionException {
		GeContext _localctx = new GeContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_ge);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(207);
			match(GE);
			setState(208);
			match(T__1);
			setState(209);
			scalar();
			setState(210);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LtContext extends ParserRuleContext {
		public TerminalNode LT() { return getToken(SSQLParser.LT, 0); }
		public ScalarContext scalar() {
			return getRuleContext(ScalarContext.class,0);
		}
		public LtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lt; }
	}

	public final LtContext lt() throws RecognitionException {
		LtContext _localctx = new LtContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_lt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(212);
			match(LT);
			setState(213);
			match(T__1);
			setState(214);
			scalar();
			setState(215);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LeContext extends ParserRuleContext {
		public TerminalNode LE() { return getToken(SSQLParser.LE, 0); }
		public ScalarContext scalar() {
			return getRuleContext(ScalarContext.class,0);
		}
		public LeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_le; }
	}

	public final LeContext le() throws RecognitionException {
		LeContext _localctx = new LeContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_le);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(217);
			match(LE);
			setState(218);
			match(T__1);
			setState(219);
			scalar();
			setState(220);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InContext extends ParserRuleContext {
		public TerminalNode IN() { return getToken(SSQLParser.IN, 0); }
		public ListContext list() {
			return getRuleContext(ListContext.class,0);
		}
		public InContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_in; }
	}

	public final InContext in() throws RecognitionException {
		InContext _localctx = new InContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_in);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(222);
			match(IN);
			setState(223);
			match(T__1);
			setState(224);
			list();
			setState(225);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BetweenContext extends ParserRuleContext {
		public TerminalNode BETWEEN() { return getToken(SSQLParser.BETWEEN, 0); }
		public List<TerminalNode> INTEGER() { return getTokens(SSQLParser.INTEGER); }
		public TerminalNode INTEGER(int i) {
			return getToken(SSQLParser.INTEGER, i);
		}
		public List<TerminalNode> REAL_NUMBER() { return getTokens(SSQLParser.REAL_NUMBER); }
		public TerminalNode REAL_NUMBER(int i) {
			return getToken(SSQLParser.REAL_NUMBER, i);
		}
		public BetweenContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_between; }
	}

	public final BetweenContext between() throws RecognitionException {
		BetweenContext _localctx = new BetweenContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_between);
		try {
			setState(239);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(227);
				match(BETWEEN);
				setState(228);
				match(T__1);
				setState(229);
				match(INTEGER);
				setState(230);
				match(T__0);
				setState(231);
				match(INTEGER);
				setState(232);
				match(T__2);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(233);
				match(BETWEEN);
				setState(234);
				match(T__1);
				setState(235);
				match(REAL_NUMBER);
				setState(236);
				match(T__0);
				setState(237);
				match(REAL_NUMBER);
				setState(238);
				match(T__2);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ContainContext extends ParserRuleContext {
		public TerminalNode CONTAIN() { return getToken(SSQLParser.CONTAIN, 0); }
		public TerminalNode STRING() { return getToken(SSQLParser.STRING, 0); }
		public ContainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_contain; }
	}

	public final ContainContext contain() throws RecognitionException {
		ContainContext _localctx = new ContainContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_contain);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(241);
			match(CONTAIN);
			setState(242);
			match(T__1);
			setState(243);
			match(STRING);
			setState(244);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExistContext extends ParserRuleContext {
		public TerminalNode EXIST() { return getToken(SSQLParser.EXIST, 0); }
		public ExistContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exist; }
	}

	public final ExistContext exist() throws RecognitionException {
		ExistContext _localctx = new ExistContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_exist);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(246);
			match(EXIST);
			setState(249);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__1) {
				{
				setState(247);
				match(T__1);
				setState(248);
				match(T__2);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TimeframeContext extends ParserRuleContext {
		public TerminalNode TIMEFRAME() { return getToken(SSQLParser.TIMEFRAME, 0); }
		public List<TerminalNode> INTEGER() { return getTokens(SSQLParser.INTEGER); }
		public TerminalNode INTEGER(int i) {
			return getToken(SSQLParser.INTEGER, i);
		}
		public TimeframeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_timeframe; }
	}

	public final TimeframeContext timeframe() throws RecognitionException {
		TimeframeContext _localctx = new TimeframeContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_timeframe);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(251);
			match(TIMEFRAME);
			setState(252);
			match(T__1);
			setState(253);
			match(INTEGER);
			setState(254);
			match(T__0);
			setState(255);
			match(INTEGER);
			setState(256);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class KeyContext extends ParserRuleContext {
		public TerminalNode KEY() { return getToken(SSQLParser.KEY, 0); }
		public TerminalNode INTEGER() { return getToken(SSQLParser.INTEGER, 0); }
		public TerminalNode STRING() { return getToken(SSQLParser.STRING, 0); }
		public KeyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_key; }
	}

	public final KeyContext key() throws RecognitionException {
		KeyContext _localctx = new KeyContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_key);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(258);
			match(KEY);
			setState(259);
			match(T__1);
			setState(260);
			_la = _input.LA(1);
			if ( !(_la==STRING || _la==INTEGER) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(261);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ScalarContext extends ParserRuleContext {
		public TerminalNode REAL_NUMBER() { return getToken(SSQLParser.REAL_NUMBER, 0); }
		public List<TerminalNode> INTEGER() { return getTokens(SSQLParser.INTEGER); }
		public TerminalNode INTEGER(int i) {
			return getToken(SSQLParser.INTEGER, i);
		}
		public ScalarContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_scalar; }
	}

	public final ScalarContext scalar() throws RecognitionException {
		ScalarContext _localctx = new ScalarContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_scalar);
		int _la;
		try {
			setState(269);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case REAL_NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(263);
				match(REAL_NUMBER);
				}
				break;
			case INTEGER:
				enterOuterAlt(_localctx, 2);
				{
				setState(265); 
				_errHandler.sync(this);
				_la = _input.LA(1);
				do {
					{
					{
					setState(264);
					match(INTEGER);
					}
					}
					setState(267); 
					_errHandler.sync(this);
					_la = _input.LA(1);
				} while ( _la==INTEGER );
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListContext extends ParserRuleContext {
		public StringListContext stringList() {
			return getRuleContext(StringListContext.class,0);
		}
		public DoubleListContext doubleList() {
			return getRuleContext(DoubleListContext.class,0);
		}
		public IntListContext intList() {
			return getRuleContext(IntListContext.class,0);
		}
		public ListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list; }
	}

	public final ListContext list() throws RecognitionException {
		ListContext _localctx = new ListContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_list);
		try {
			setState(274);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(271);
				stringList();
				}
				break;
			case REAL_NUMBER:
				enterOuterAlt(_localctx, 2);
				{
				setState(272);
				doubleList();
				}
				break;
			case INTEGER:
				enterOuterAlt(_localctx, 3);
				{
				setState(273);
				intList();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StringListContext extends ParserRuleContext {
		public List<TerminalNode> STRING() { return getTokens(SSQLParser.STRING); }
		public TerminalNode STRING(int i) {
			return getToken(SSQLParser.STRING, i);
		}
		public StringListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringList; }
	}

	public final StringListContext stringList() throws RecognitionException {
		StringListContext _localctx = new StringListContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_stringList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(276);
			match(STRING);
			setState(281);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0) {
				{
				{
				setState(277);
				match(T__0);
				setState(278);
				match(STRING);
				}
				}
				setState(283);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DoubleListContext extends ParserRuleContext {
		public List<TerminalNode> REAL_NUMBER() { return getTokens(SSQLParser.REAL_NUMBER); }
		public TerminalNode REAL_NUMBER(int i) {
			return getToken(SSQLParser.REAL_NUMBER, i);
		}
		public DoubleListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_doubleList; }
	}

	public final DoubleListContext doubleList() throws RecognitionException {
		DoubleListContext _localctx = new DoubleListContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_doubleList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(284);
			match(REAL_NUMBER);
			setState(289);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0) {
				{
				{
				setState(285);
				match(T__0);
				setState(286);
				match(REAL_NUMBER);
				}
				}
				setState(291);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IntListContext extends ParserRuleContext {
		public List<TerminalNode> INTEGER() { return getTokens(SSQLParser.INTEGER); }
		public TerminalNode INTEGER(int i) {
			return getToken(SSQLParser.INTEGER, i);
		}
		public IntListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_intList; }
	}

	public final IntListContext intList() throws RecognitionException {
		IntListContext _localctx = new IntListContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_intList);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(292);
			match(INTEGER);
			setState(297);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0) {
				{
				{
				setState(293);
				match(T__0);
				setState(294);
				match(INTEGER);
				}
				}
				setState(299);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OrderByContext extends ParserRuleContext {
		public TerminalNode ORDER_BY() { return getToken(SSQLParser.ORDER_BY, 0); }
		public List<OrderContext> order() {
			return getRuleContexts(OrderContext.class);
		}
		public OrderContext order(int i) {
			return getRuleContext(OrderContext.class,i);
		}
		public OrderByContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_orderBy; }
	}

	public final OrderByContext orderBy() throws RecognitionException {
		OrderByContext _localctx = new OrderByContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_orderBy);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(300);
			match(ORDER_BY);
			setState(301);
			order();
			setState(306);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0) {
				{
				{
				setState(302);
				match(T__0);
				setState(303);
				order();
				}
				}
				setState(308);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class OrderContext extends ParserRuleContext {
		public Token dir;
		public TerminalNode IDENTIFIER() { return getToken(SSQLParser.IDENTIFIER, 0); }
		public TerminalNode ASC() { return getToken(SSQLParser.ASC, 0); }
		public TerminalNode DESC() { return getToken(SSQLParser.DESC, 0); }
		public OrderContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_order; }
	}

	public final OrderContext order() throws RecognitionException {
		OrderContext _localctx = new OrderContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_order);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(309);
			match(IDENTIFIER);
			setState(311);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ASC || _la==DESC) {
				{
				setState(310);
				((OrderContext)_localctx).dir = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==ASC || _la==DESC) ) {
					((OrderContext)_localctx).dir = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LimitContext extends ParserRuleContext {
		public TerminalNode LIMIT() { return getToken(SSQLParser.LIMIT, 0); }
		public List<TerminalNode> INTEGER() { return getTokens(SSQLParser.INTEGER); }
		public TerminalNode INTEGER(int i) {
			return getToken(SSQLParser.INTEGER, i);
		}
		public LimitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_limit; }
	}

	public final LimitContext limit() throws RecognitionException {
		LimitContext _localctx = new LimitContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_limit);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(313);
			match(LIMIT);
			setState(315); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(314);
				match(INTEGER);
				}
				}
				setState(317); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==INTEGER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3+\u0142\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\3\2\3\2\3\2\3\2\3\2\5\2J\n\2\3\2\5\2M\n\2\3\2\3\2\3\3\3\3"+
		"\3\3\7\3T\n\3\f\3\16\3W\13\3\3\3\3\3\3\3\3\3\3\3\7\3^\n\3\f\3\16\3a\13"+
		"\3\3\3\3\3\3\3\7\3f\n\3\f\3\16\3i\13\3\5\3k\n\3\3\4\3\4\5\4o\n\4\3\5\3"+
		"\5\3\5\3\5\3\5\5\5v\n\5\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\7\3\7\5\7\u0081"+
		"\n\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\7\t\u008c\n\t\f\t\16\t\u008f"+
		"\13\t\3\n\3\n\3\n\5\n\u0094\n\n\3\13\3\13\5\13\u0098\n\13\3\13\3\13\3"+
		"\13\6\13\u009d\n\13\r\13\16\13\u009e\5\13\u00a1\n\13\3\13\3\13\3\f\3\f"+
		"\6\f\u00a7\n\f\r\f\16\f\u00a8\3\f\3\f\3\r\3\r\6\r\u00af\n\r\r\r\16\r\u00b0"+
		"\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5"+
		"\16\u00c1\n\16\3\17\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\21"+
		"\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23\3\23"+
		"\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\26\3\26\3\26\3\26\5\26\u00f2\n\26\3\27\3\27\3\27"+
		"\3\27\3\27\3\30\3\30\3\30\5\30\u00fc\n\30\3\31\3\31\3\31\3\31\3\31\3\31"+
		"\3\31\3\32\3\32\3\32\3\32\3\32\3\33\3\33\6\33\u010c\n\33\r\33\16\33\u010d"+
		"\5\33\u0110\n\33\3\34\3\34\3\34\5\34\u0115\n\34\3\35\3\35\3\35\7\35\u011a"+
		"\n\35\f\35\16\35\u011d\13\35\3\36\3\36\3\36\7\36\u0122\n\36\f\36\16\36"+
		"\u0125\13\36\3\37\3\37\3\37\7\37\u012a\n\37\f\37\16\37\u012d\13\37\3 "+
		"\3 \3 \3 \7 \u0133\n \f \16 \u0136\13 \3!\3!\5!\u013a\n!\3\"\3\"\6\"\u013e"+
		"\n\"\r\"\16\"\u013f\3\"\2\2#\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \""+
		"$&(*,.\60\62\64\668:<>@B\2\5\3\2\13\17\3\2\'(\3\2#$\2\u0149\2D\3\2\2\2"+
		"\4j\3\2\2\2\6n\3\2\2\2\bu\3\2\2\2\nw\3\2\2\2\f\u0080\3\2\2\2\16\u0082"+
		"\3\2\2\2\20\u0089\3\2\2\2\22\u0093\3\2\2\2\24\u0095\3\2\2\2\26\u00a4\3"+
		"\2\2\2\30\u00ac\3\2\2\2\32\u00c0\3\2\2\2\34\u00c2\3\2\2\2\36\u00c7\3\2"+
		"\2\2 \u00cc\3\2\2\2\"\u00d1\3\2\2\2$\u00d6\3\2\2\2&\u00db\3\2\2\2(\u00e0"+
		"\3\2\2\2*\u00f1\3\2\2\2,\u00f3\3\2\2\2.\u00f8\3\2\2\2\60\u00fd\3\2\2\2"+
		"\62\u0104\3\2\2\2\64\u010f\3\2\2\2\66\u0114\3\2\2\28\u0116\3\2\2\2:\u011e"+
		"\3\2\2\2<\u0126\3\2\2\2>\u012e\3\2\2\2@\u0137\3\2\2\2B\u013b\3\2\2\2D"+
		"E\7\36\2\2EF\5\4\3\2FG\7\37\2\2GI\5\20\t\2HJ\5> \2IH\3\2\2\2IJ\3\2\2\2"+
		"JL\3\2\2\2KM\5B\"\2LK\3\2\2\2LM\3\2\2\2MN\3\2\2\2NO\7\2\2\3O\3\3\2\2\2"+
		"PU\5\6\4\2QR\7\3\2\2RT\5\6\4\2SQ\3\2\2\2TW\3\2\2\2US\3\2\2\2UV\3\2\2\2"+
		"Vk\3\2\2\2WU\3\2\2\2XY\7!\2\2YZ\7\4\2\2Z_\5\f\7\2[\\\7\3\2\2\\^\5\f\7"+
		"\2][\3\2\2\2^a\3\2\2\2_]\3\2\2\2_`\3\2\2\2`b\3\2\2\2a_\3\2\2\2bg\7\5\2"+
		"\2cd\7\3\2\2df\5\b\5\2ec\3\2\2\2fi\3\2\2\2ge\3\2\2\2gh\3\2\2\2hk\3\2\2"+
		"\2ig\3\2\2\2jP\3\2\2\2jX\3\2\2\2k\5\3\2\2\2lo\7*\2\2mo\5\b\5\2nl\3\2\2"+
		"\2nm\3\2\2\2o\7\3\2\2\2pq\t\2\2\2qr\7\4\2\2rs\7*\2\2sv\7\5\2\2tv\5\n\6"+
		"\2up\3\2\2\2ut\3\2\2\2v\t\3\2\2\2wx\7\20\2\2xy\7\4\2\2yz\7*\2\2z{\7\3"+
		"\2\2{|\7)\2\2|}\7\5\2\2}\13\3\2\2\2~\u0081\7*\2\2\177\u0081\5\16\b\2\u0080"+
		"~\3\2\2\2\u0080\177\3\2\2\2\u0081\r\3\2\2\2\u0082\u0083\7\21\2\2\u0083"+
		"\u0084\7\4\2\2\u0084\u0085\7*\2\2\u0085\u0086\7\3\2\2\u0086\u0087\7(\2"+
		"\2\u0087\u0088\7\5\2\2\u0088\17\3\2\2\2\u0089\u008d\5\22\n\2\u008a\u008c"+
		"\5\22\n\2\u008b\u008a\3\2\2\2\u008c\u008f\3\2\2\2\u008d\u008b\3\2\2\2"+
		"\u008d\u008e\3\2\2\2\u008e\21\3\2\2\2\u008f\u008d\3\2\2\2\u0090\u0094"+
		"\5\24\13\2\u0091\u0094\5\26\f\2\u0092\u0094\5\30\r\2\u0093\u0090\3\2\2"+
		"\2\u0093\u0091\3\2\2\2\u0093\u0092\3\2\2\2\u0094\23\3\2\2\2\u0095\u0097"+
		"\7\6\2\2\u0096\u0098\7*\2\2\u0097\u0096\3\2\2\2\u0097\u0098\3\2\2\2\u0098"+
		"\u0099\3\2\2\2\u0099\u00a0\7&\2\2\u009a\u00a1\5\32\16\2\u009b\u009d\5"+
		"\24\13\2\u009c\u009b\3\2\2\2\u009d\u009e\3\2\2\2\u009e\u009c\3\2\2\2\u009e"+
		"\u009f\3\2\2\2\u009f\u00a1\3\2\2\2\u00a0\u009a\3\2\2\2\u00a0\u009c\3\2"+
		"\2\2\u00a0\u00a1\3\2\2\2\u00a1\u00a2\3\2\2\2\u00a2\u00a3\7\7\2\2\u00a3"+
		"\25\3\2\2\2\u00a4\u00a6\7\b\2\2\u00a5\u00a7\5\22\n\2\u00a6\u00a5\3\2\2"+
		"\2\u00a7\u00a8\3\2\2\2\u00a8\u00a6\3\2\2\2\u00a8\u00a9\3\2\2\2\u00a9\u00aa"+
		"\3\2\2\2\u00aa\u00ab\7\t\2\2\u00ab\27\3\2\2\2\u00ac\u00ae\7\n\2\2\u00ad"+
		"\u00af\5\22\n\2\u00ae\u00ad\3\2\2\2\u00af\u00b0\3\2\2\2\u00b0\u00ae\3"+
		"\2\2\2\u00b0\u00b1\3\2\2\2\u00b1\u00b2\3\2\2\2\u00b2\u00b3\7\t\2\2\u00b3"+
		"\31\3\2\2\2\u00b4\u00c1\5\34\17\2\u00b5\u00c1\5\36\20\2\u00b6\u00c1\5"+
		" \21\2\u00b7\u00c1\5\"\22\2\u00b8\u00c1\5$\23\2\u00b9\u00c1\5&\24\2\u00ba"+
		"\u00c1\5(\25\2\u00bb\u00c1\5*\26\2\u00bc\u00c1\5,\27\2\u00bd\u00c1\5."+
		"\30\2\u00be\u00c1\5\60\31\2\u00bf\u00c1\5\62\32\2\u00c0\u00b4\3\2\2\2"+
		"\u00c0\u00b5\3\2\2\2\u00c0\u00b6\3\2\2\2\u00c0\u00b7\3\2\2\2\u00c0\u00b8"+
		"\3\2\2\2\u00c0\u00b9\3\2\2\2\u00c0\u00ba\3\2\2\2\u00c0\u00bb\3\2\2\2\u00c0"+
		"\u00bc\3\2\2\2\u00c0\u00bd\3\2\2\2\u00c0\u00be\3\2\2\2\u00c0\u00bf\3\2"+
		"\2\2\u00c1\33\3\2\2\2\u00c2\u00c3\7\22\2\2\u00c3\u00c4\7\4\2\2\u00c4\u00c5"+
		"\5\64\33\2\u00c5\u00c6\7\5\2\2\u00c6\35\3\2\2\2\u00c7\u00c8\7\23\2\2\u00c8"+
		"\u00c9\7\4\2\2\u00c9\u00ca\5\64\33\2\u00ca\u00cb\7\5\2\2\u00cb\37\3\2"+
		"\2\2\u00cc\u00cd\7\30\2\2\u00cd\u00ce\7\4\2\2\u00ce\u00cf\5\64\33\2\u00cf"+
		"\u00d0\7\5\2\2\u00d0!\3\2\2\2\u00d1\u00d2\7\27\2\2\u00d2\u00d3\7\4\2\2"+
		"\u00d3\u00d4\5\64\33\2\u00d4\u00d5\7\5\2\2\u00d5#\3\2\2\2\u00d6\u00d7"+
		"\7\25\2\2\u00d7\u00d8\7\4\2\2\u00d8\u00d9\5\64\33\2\u00d9\u00da\7\5\2"+
		"\2\u00da%\3\2\2\2\u00db\u00dc\7\26\2\2\u00dc\u00dd\7\4\2\2\u00dd\u00de"+
		"\5\64\33\2\u00de\u00df\7\5\2\2\u00df\'\3\2\2\2\u00e0\u00e1\7\24\2\2\u00e1"+
		"\u00e2\7\4\2\2\u00e2\u00e3\5\66\34\2\u00e3\u00e4\7\5\2\2\u00e4)\3\2\2"+
		"\2\u00e5\u00e6\7\31\2\2\u00e6\u00e7\7\4\2\2\u00e7\u00e8\7(\2\2\u00e8\u00e9"+
		"\7\3\2\2\u00e9\u00ea\7(\2\2\u00ea\u00f2\7\5\2\2\u00eb\u00ec\7\31\2\2\u00ec"+
		"\u00ed\7\4\2\2\u00ed\u00ee\7)\2\2\u00ee\u00ef\7\3\2\2\u00ef\u00f0\7)\2"+
		"\2\u00f0\u00f2\7\5\2\2\u00f1\u00e5\3\2\2\2\u00f1\u00eb\3\2\2\2\u00f2+"+
		"\3\2\2\2\u00f3\u00f4\7\32\2\2\u00f4\u00f5\7\4\2\2\u00f5\u00f6\7\'\2\2"+
		"\u00f6\u00f7\7\5\2\2\u00f7-\3\2\2\2\u00f8\u00fb\7\33\2\2\u00f9\u00fa\7"+
		"\4\2\2\u00fa\u00fc\7\5\2\2\u00fb\u00f9\3\2\2\2\u00fb\u00fc\3\2\2\2\u00fc"+
		"/\3\2\2\2\u00fd\u00fe\7\34\2\2\u00fe\u00ff\7\4\2\2\u00ff\u0100\7(\2\2"+
		"\u0100\u0101\7\3\2\2\u0101\u0102\7(\2\2\u0102\u0103\7\5\2\2\u0103\61\3"+
		"\2\2\2\u0104\u0105\7\35\2\2\u0105\u0106\7\4\2\2\u0106\u0107\t\3\2\2\u0107"+
		"\u0108\7\5\2\2\u0108\63\3\2\2\2\u0109\u0110\7)\2\2\u010a\u010c\7(\2\2"+
		"\u010b\u010a\3\2\2\2\u010c\u010d\3\2\2\2\u010d\u010b\3\2\2\2\u010d\u010e"+
		"\3\2\2\2\u010e\u0110\3\2\2\2\u010f\u0109\3\2\2\2\u010f\u010b\3\2\2\2\u0110"+
		"\65\3\2\2\2\u0111\u0115\58\35\2\u0112\u0115\5:\36\2\u0113\u0115\5<\37"+
		"\2\u0114\u0111\3\2\2\2\u0114\u0112\3\2\2\2\u0114\u0113\3\2\2\2\u0115\67"+
		"\3\2\2\2\u0116\u011b\7\'\2\2\u0117\u0118\7\3\2\2\u0118\u011a\7\'\2\2\u0119"+
		"\u0117\3\2\2\2\u011a\u011d\3\2\2\2\u011b\u0119\3\2\2\2\u011b\u011c\3\2"+
		"\2\2\u011c9\3\2\2\2\u011d\u011b\3\2\2\2\u011e\u0123\7)\2\2\u011f\u0120"+
		"\7\3\2\2\u0120\u0122\7)\2\2\u0121\u011f\3\2\2\2\u0122\u0125\3\2\2\2\u0123"+
		"\u0121\3\2\2\2\u0123\u0124\3\2\2\2\u0124;\3\2\2\2\u0125\u0123\3\2\2\2"+
		"\u0126\u012b\7(\2\2\u0127\u0128\7\3\2\2\u0128\u012a\7(\2\2\u0129\u0127"+
		"\3\2\2\2\u012a\u012d\3\2\2\2\u012b\u0129\3\2\2\2\u012b\u012c\3\2\2\2\u012c"+
		"=\3\2\2\2\u012d\u012b\3\2\2\2\u012e\u012f\7 \2\2\u012f\u0134\5@!\2\u0130"+
		"\u0131\7\3\2\2\u0131\u0133\5@!\2\u0132\u0130\3\2\2\2\u0133\u0136\3\2\2"+
		"\2\u0134\u0132\3\2\2\2\u0134\u0135\3\2\2\2\u0135?\3\2\2\2\u0136\u0134"+
		"\3\2\2\2\u0137\u0139\7*\2\2\u0138\u013a\t\4\2\2\u0139\u0138\3\2\2\2\u0139"+
		"\u013a\3\2\2\2\u013aA\3\2\2\2\u013b\u013d\7\"\2\2\u013c\u013e\7(\2\2\u013d"+
		"\u013c\3\2\2\2\u013e\u013f\3\2\2\2\u013f\u013d\3\2\2\2\u013f\u0140\3\2"+
		"\2\2\u0140C\3\2\2\2\36ILU_gjnu\u0080\u008d\u0093\u0097\u009e\u00a0\u00a8"+
		"\u00b0\u00c0\u00f1\u00fb\u010d\u010f\u0114\u011b\u0123\u012b\u0134\u0139"+
		"\u013f";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}