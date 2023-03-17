#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Aux {
    #[prost(bool, tag="1")]
    pub local_only: bool,
    #[prost(uint32, tag="2")]
    pub tuple_count: u32,
    #[prost(uint32, repeated, tag="3")]
    pub aggregate_attributes: ::prost::alloc::vec::Vec<u32>,
    #[prost(uint32, repeated, tag="4")]
    pub group_attributes: ::prost::alloc::vec::Vec<u32>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Statement {
    #[prost(message, repeated, tag="1")]
    pub find: ::prost::alloc::vec::Vec<Attribute>,
    #[prost(string, tag="2")]
    pub from: ::prost::alloc::string::String,
    #[prost(message, repeated, tag="3")]
    pub r#where: ::prost::alloc::vec::Vec<Expr>,
    #[prost(message, repeated, tag="4")]
    pub order_by: ::prost::alloc::vec::Vec<OrderBy>,
    #[prost(uint32, tag="5")]
    pub limit: u32,
    #[prost(message, optional, tag="6")]
    pub aux: ::core::option::Option<Aux>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Parameter {
    #[prost(oneof="parameter::Value", tags="1, 2")]
    pub value: ::core::option::Option<parameter::Value>,
}
/// Nested message and enum types in `Parameter`.
pub mod parameter {
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Value {
        #[prost(double, tag="1")]
        Double(f64),
        #[prost(int64, tag="2")]
        Int(i64),
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Function {
    #[prost(enumeration="function::Func", tag="1")]
    pub name: i32,
    #[prost(message, repeated, tag="2")]
    pub parameter: ::prost::alloc::vec::Vec<Parameter>,
}
/// Nested message and enum types in `Function`.
pub mod function {
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
    #[repr(i32)]
    pub enum Func {
        Avg = 0,
        Sum = 1,
        Min = 2,
        Max = 3,
        Count = 4,
        Percentile = 5,
        Partition = 6,
    }
    impl Func {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                Func::Avg => "AVG",
                Func::Sum => "SUM",
                Func::Min => "MIN",
                Func::Max => "MAX",
                Func::Count => "COUNT",
                Func::Percentile => "PERCENTILE",
                Func::Partition => "PARTITION",
            }
        }
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Timestamp {
    #[prost(enumeration="timestamp::TimeUnit", tag="1")]
    pub timeunit: i32,
    #[prost(string, tag="2")]
    pub tz: ::prost::alloc::string::String,
}
/// Nested message and enum types in `Timestamp`.
pub mod timestamp {
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
    #[repr(i32)]
    pub enum TimeUnit {
        Second = 0,
        Millisecond = 1,
        Microsecond = 2,
        Nanosecond = 3,
    }
    impl TimeUnit {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                TimeUnit::Second => "Second",
                TimeUnit::Millisecond => "Millisecond",
                TimeUnit::Microsecond => "Microsecond",
                TimeUnit::Nanosecond => "Nanosecond",
            }
        }
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Field {
    #[prost(uint32, tag="1")]
    pub unicode: u32,
    #[prost(string, tag="2")]
    pub name: ::prost::alloc::string::String,
    #[prost(enumeration="Type", tag="3")]
    pub r#type: i32,
    #[prost(message, optional, tag="4")]
    pub timestamp: ::core::option::Option<Timestamp>,
    #[prost(oneof="field::Nested", tags="5, 6")]
    pub nested: ::core::option::Option<field::Nested>,
}
/// Nested message and enum types in `Field`.
pub mod field {
    #[derive(Clone, PartialEq, ::prost::Message)]
    pub struct Struct {
        #[prost(message, repeated, tag="1")]
        pub fields: ::prost::alloc::vec::Vec<super::Field>,
    }
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Nested {
        #[prost(message, tag="5")]
        Struct(Struct),
        #[prost(message, tag="6")]
        List(::prost::alloc::boxed::Box<super::Field>),
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Attribute {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
    #[prost(bool, tag="2")]
    pub group: bool,
    #[prost(message, optional, tag="3")]
    pub func: ::core::option::Option<Function>,
    /// Type type = 4;
    #[prost(message, optional, tag="4")]
    pub arrow: ::core::option::Option<Field>,
    #[prost(uint32, tag="5")]
    pub tuple_id: u32,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Expr {
    #[prost(oneof="expr::Field", tags="1, 2")]
    pub field: ::core::option::Option<expr::Field>,
}
/// Nested message and enum types in `Expr`.
pub mod expr {
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Field {
        #[prost(message, tag="1")]
        Or(super::Or),
        #[prost(message, tag="2")]
        Tuple(super::Tuple),
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Or {
    #[prost(message, repeated, tag="1")]
    pub expr: ::prost::alloc::vec::Vec<Expr>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Operand {
    #[prost(string, tag="1")]
    pub content: ::prost::alloc::string::String,
    #[prost(oneof="operand::Value", tags="2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16")]
    pub value: ::core::option::Option<operand::Value>,
}
/// Nested message and enum types in `Operand`.
pub mod operand {
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Value {
        #[prost(int32, tag="2")]
        Int8(i32),
        #[prost(uint32, tag="3")]
        Uint8(u32),
        #[prost(int32, tag="4")]
        Int16(i32),
        #[prost(uint32, tag="5")]
        Uint16(u32),
        #[prost(int32, tag="6")]
        Int32(i32),
        #[prost(int32, tag="7")]
        Date32(i32),
        #[prost(uint32, tag="8")]
        Uint32(u32),
        #[prost(float, tag="9")]
        Float32(f32),
        #[prost(int64, tag="10")]
        Int64(i64),
        #[prost(int64, tag="11")]
        Date64(i64),
        #[prost(uint64, tag="12")]
        Uint64(u64),
        #[prost(double, tag="13")]
        Float64(f64),
        #[prost(bool, tag="14")]
        Bool(bool),
        #[prost(bool, tag="15")]
        Text(bool),
        #[prost(bool, tag="16")]
        Null(bool),
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Tuple {
    #[prost(uint32, tag="1")]
    pub id: u32,
    #[prost(string, tag="2")]
    pub name: ::prost::alloc::string::String,
    #[prost(string, tag="3")]
    pub path: ::prost::alloc::string::String,
    #[prost(uint32, tag="4")]
    pub unicode: u32,
    #[prost(oneof="tuple::Predicate", tags="5, 6, 7, 8, 9")]
    pub predicate: ::core::option::Option<tuple::Predicate>,
}
/// Nested message and enum types in `Tuple`.
pub mod tuple {
    #[derive(Clone, PartialEq, ::prost::Oneof)]
    pub enum Predicate {
        #[prost(message, tag="5")]
        Nested(super::Nested),
        #[prost(message, tag="6")]
        Binary(super::Binary),
        #[prost(message, tag="7")]
        Unary(super::Unary),
        #[prost(message, tag="8")]
        List(super::List),
        #[prost(enumeration="super::Type", tag="9")]
        SelectOnly(i32),
    }
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Nested {
    #[prost(message, repeated, tag="1")]
    pub expr: ::prost::alloc::vec::Vec<Expr>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Unary {
    #[prost(enumeration="Operator", tag="1")]
    pub op: i32,
    #[prost(message, optional, tag="2")]
    pub operand: ::core::option::Option<Operand>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct Binary {
    #[prost(message, optional, tag="1")]
    pub first: ::core::option::Option<Operand>,
    #[prost(message, optional, tag="2")]
    pub second: ::core::option::Option<Operand>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct List {
    #[prost(enumeration="Type", tag="1")]
    pub r#type: i32,
    #[prost(message, repeated, tag="2")]
    pub operands: ::prost::alloc::vec::Vec<Operand>,
}
#[derive(Clone, PartialEq, ::prost::Message)]
pub struct OrderBy {
    #[prost(string, tag="1")]
    pub name: ::prost::alloc::string::String,
    #[prost(enumeration="order_by::Direction", tag="2")]
    pub direction: i32,
}
/// Nested message and enum types in `OrderBy`.
pub mod order_by {
    #[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
    #[repr(i32)]
    pub enum Direction {
        Asc = 0,
        Desc = 1,
    }
    impl Direction {
        /// String value of the enum field names used in the ProtoBuf definition.
        ///
        /// The values are not transformed in any way and thus are considered stable
        /// (if the ProtoBuf definition does not change) and safe for programmatic use.
        pub fn as_str_name(&self) -> &'static str {
            match self {
                Direction::Asc => "ASC",
                Direction::Desc => "DESC",
            }
        }
    }
}
/// note: ssd runtime contains static convert table to arrow2 data type
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum Type {
    List = 0,
    Struct = 1,
    Int8 = 2,
    Uint8 = 3,
    Int16 = 4,
    Uint16 = 5,
    Int32 = 6,
    Uint32 = 7,
    Date32 = 8,
    Float32 = 9,
    Int64 = 10,
    Uint64 = 11,
    Date64 = 12,
    Float64 = 13,
    Bool = 14,
    Text = 15,
    Null = 16,
}
impl Type {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            Type::List => "list",
            Type::Struct => "struct",
            Type::Int8 => "int8",
            Type::Uint8 => "uint8",
            Type::Int16 => "int16",
            Type::Uint16 => "uint16",
            Type::Int32 => "int32",
            Type::Uint32 => "uint32",
            Type::Date32 => "date32",
            Type::Float32 => "float32",
            Type::Int64 => "int64",
            Type::Uint64 => "uint64",
            Type::Date64 => "date64",
            Type::Float64 => "float64",
            Type::Bool => "bool",
            Type::Text => "text",
            Type::Null => "null",
        }
    }
}
#[derive(Clone, Copy, Debug, PartialEq, Eq, Hash, PartialOrd, Ord, ::prost::Enumeration)]
#[repr(i32)]
pub enum Operator {
    Eq = 0,
    Neq = 1,
    Gt = 2,
    Ge = 3,
    Lt = 4,
    Le = 5,
    Between = 6,
    Contain = 7,
    Exist = 8,
    Isnull = 9,
}
impl Operator {
    /// String value of the enum field names used in the ProtoBuf definition.
    ///
    /// The values are not transformed in any way and thus are considered stable
    /// (if the ProtoBuf definition does not change) and safe for programmatic use.
    pub fn as_str_name(&self) -> &'static str {
        match self {
            Operator::Eq => "EQ",
            Operator::Neq => "NEQ",
            Operator::Gt => "GT",
            Operator::Ge => "GE",
            Operator::Lt => "LT",
            Operator::Le => "LE",
            Operator::Between => "BETWEEN",
            Operator::Contain => "CONTAIN",
            Operator::Exist => "EXIST",
            Operator::Isnull => "ISNULL",
        }
    }
}

pub mod arrow;