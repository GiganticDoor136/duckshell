use std::process::Command;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    let output = Command::new("./modules/dsh/cmd/dsh/dsh")
        .args(&args[1..]) 
        .output();

    match output {
        Ok(output) => {
            if output.status.success() {
                print!("{}", String::from_utf8_lossy(&output.stdout));
            } else {
                eprintln!("Error: {}", String::from_utf8_lossy(&output.stderr));
                eprintln!("Exit code: {}", output.status);
            }
        }
        Err(e) => {
            eprintln!("Failed to execute dsh: {}", e);
        }
    }
}