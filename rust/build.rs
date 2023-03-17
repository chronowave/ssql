use std::env;
use std::fs;
use std::io::{Result, Write};
use std::path::Path;

use prost_build;

fn main() -> Result<()> {
    println!("cargo:rerun-if-changed=../ssql.proto");
    let mut config = prost_build::Config::new();

    config
        .btree_map(&["."])
        .compile_protos(&["../ssql.proto"], &["../"])?;

    let output = Path::new(env::var("OUT_DIR").unwrap().as_str()).join("_.rs");

    println!("output dir {:?}", output.as_path());

    let mut file = std::fs::OpenOptions::new()
        .write(true)
        .append(true)
        .open(output.as_path())
        .unwrap();

    write!(
        file,
        r#"
pub mod arrow;"#
    )?;

    fs::copy(output.as_path(), "src/lib.rs")?;

    Ok(())
}
