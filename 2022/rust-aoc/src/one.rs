use crate::utils::read::read_lines;

const INPUT: &str = "../input/1.txt";

pub fn main() {
    println!("Hello, world ONE!");
    if let Ok(lines) = read_lines(INPUT) {
        // Consumes the iterator, returns an (Optional) String
        for line in lines {
            if let Ok(l) = line {
                println!("{}", l);
            }
        }
    }

}
