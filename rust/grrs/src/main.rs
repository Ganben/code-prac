use structopt::StructOpt;
use std::io::prelude::*;
use std::io::BufReader;
use std::fs::File;
use log::{info, warn};
use std::time::Instant;

use indicatif::{HumanDuration, ProgressBar};

// use std::io::{self, BufRead};
/// Search for a pattern in a file and display the lines that contain it.
#[derive(StructOpt)]

struct Cli {
    pattern: String,
    /// The path to the file to read
    #[structopt(parse(from_os_str))]
    path: std::path::PathBuf,
}



fn main() -> Result<(), Box<dyn std::error::Error>> {
    // println!("Hello, world!");
    // let pattern = std::env::args().nth(1).expect("no pattern given");
    // let path = std::env::args().nth(2).expect("no path given");
    let args = Cli::from_args();
    // let content = std::fs::read_to_string(&args.path)
    let f = File::open(&args.path)?;
    let mut reader = BufReader::new(f);
    let mut line = String::new();
    

    // for line in content {
    //     if line.contains(&args.pattern) {
    //         println!("{}", line);
    //     }
    // }
    let pb = indicatif::ProgressBar::new(100);
    for i in 0..100 {
        // do_hard_work();
        let content = reader.read_line(&mut line)?;
    //   .expect("reading file faile");
        while line != "" {
        
        if line.contains(&args.pattern) {
            println!("{}", line);
        }

        line.clear();
        let content = reader.read_line(&mut line)?;
        
    }
    line.clear();
        pb.println(format!("[+] finished #{}", i));
        pb.inc(1);
    }
    pb.finish_with_message("done");

    Ok(())
}

fn find_matches(content: &str, pattern: &str, mut writer: impl std::io::Write) {
    for line in content.lines() {
        if line.contains(pattern) {
            writeln!(writer, "{}", line);
        }
    }
}

#[test]
fn find_a_match() {
    let mut result = Vec::new();
    find_matches("lorem ipsum\ndolor sit amet", "lorem", &mut result);
    assert_eq!(result, b"lorem ipsum\n");
}