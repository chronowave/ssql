use std::string::ToString;

use arrow2::datatypes::{DataType, Field, IntervalUnit, TimeUnit};

// note: should be same sequence in Type ssql.proto and supports primitive type only
pub const ARROW_TIME_UNIT: [TimeUnit; 4] = [
    TimeUnit::Second,
    TimeUnit::Millisecond,
    TimeUnit::Microsecond,
    TimeUnit::Nanosecond,
];

// note: should be same sequence in Type ssql.proto and supports primitive type only
pub const ARROW_DATA_TYPE: [DataType; 17] = [
    // list = 0;
    DataType::Null,
    //struct = 1;
    DataType::Null,
    // int8 = 2;
    DataType::Int8,
    //uint8 = 3;
    DataType::UInt8,
    //int16 = 4;
    DataType::Int16,
    //uint16 = 5;
    DataType::UInt16,
    //int32 = 6;
    DataType::Int32,
    //uint32 = 7;
    DataType::UInt32,
    //date32 = 8;
    DataType::Date32,
    //float32 = 9;
    DataType::Float32,
    //int64 = 10;
    DataType::Int64,
    //uint64 = 11;
    DataType::UInt64,
    //date64 = 12;
    DataType::Date64,
    //float64 = 13;
    DataType::Float64,
    //bool = 14;
    DataType::Boolean,
    //text = 15;
    DataType::Utf8,
    //null = 16;
    DataType::Null,
];

#[cfg(test)]
mod test {
    use arrow2::datatypes::DataType;

    use crate::arrow::{ARROW_DATA_TYPE, ARROW_TIME_UNIT};
    use crate::arrow::TimeUnit;

    #[test]
    fn assert_datatype() {
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Int8 as usize], DataType::Int8);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Uint8 as usize], DataType::UInt8);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Int16 as usize], DataType::Int16);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Uint16 as usize], DataType::UInt16);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Int32 as usize], DataType::Int32);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Uint32 as usize], DataType::UInt32);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Date32 as usize], DataType::Date32);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Float32 as usize], DataType::Float32);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Int64 as usize], DataType::Int64);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Uint64 as usize], DataType::UInt64);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Date64 as usize], DataType::Date64);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Float64 as usize], DataType::Float64);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Bool as usize], DataType::Boolean);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Text as usize], DataType::Utf8);
        assert_eq!(ARROW_DATA_TYPE[crate::Type::Null as usize], DataType::Null);
    }

    #[test]
    fn assert_timeunit() {
        assert_eq!(ARROW_TIME_UNIT[crate::timestamp::TimeUnit::Second as usize], TimeUnit::Second);
        assert_eq!(ARROW_TIME_UNIT[crate::timestamp::TimeUnit::Millisecond as usize], TimeUnit::Millisecond);
        assert_eq!(ARROW_TIME_UNIT[crate::timestamp::TimeUnit::Microsecond as usize], TimeUnit::Microsecond);
        assert_eq!(ARROW_TIME_UNIT[crate::timestamp::TimeUnit::Nanosecond as usize], TimeUnit::Nanosecond);
    }
}