use std::fs;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let contents = fs::read_to_string("input.txt")?;
    println!("{}", contents);
    Ok(())
}
