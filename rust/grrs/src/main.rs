use structopt::StructOpt;
use std::io::prelude::*;
use std::io::BufReader;
use std::fs::File;
// use std::io::{self, BufRead};
/// Search for a pattern in a file and display the lines that contain it.
#[derive(StructOpt)]

struct Cli {
    pattern: String,
    /// The path to the file to read
    #[structopt(parse(from_os_str))]
    path: std::path::PathBuf,
}

fn main() -> std::io::Result<()> {
    // println!("Hello, world!");
    // let pattern = std::env::args().nth(1).expect("no pattern given");
    // let path = std::env::args().nth(2).expect("no path given");
    let args = Cli::from_args();
    // let content = std::fs::read_to_string(&args.path)
    let f = File::open(&args.path)?;
    let mut reader = BufReader::new(f);
    let mut line = String::new();
    
    let content = reader.read_line(&mut line)?;
    //   .expect("reading file faile");
    while content > 0 {
        
        if line.contains(&args.pattern) {
            println!("{}", line);
        }

        line.clear();
        let content = reader.read_line(&mut line)?;
        
    }
    line.clear();
    // for line in content {
    //     if line.contains(&args.pattern) {
    //         println!("{}", line);
    //     }
    // }
    Ok(())
}
