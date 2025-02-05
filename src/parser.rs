use std::str::SplitWhitespace;

#[derive(Debug)]
pub enum Command {
    Ld { path: String, long_format: bool },
    Mkfile { path: String },
    Sysinfo,
    Help,
    Ver,
    Enable { function: String },
    Disable { function: String },
    Ctl { function: String, settings: String },
    Custom { name: String, args: Vec<String> }, 
    Unknown(String),
}

pub fn parse_command(input: &str) -> Command {
    let mut parts = input.split_whitespace();
    let command_name = parts.next().unwrap_or("");

    match command_name {
        "ld" => {
            let mut long_format = false;
            let mut path = String::new();
            for part in parts {
                if part == "-l" {
                    long_format = true;
                } else {
                    path = part.to_string();
                }
            }
            Command::Ld { path, long_format }
        }
        "mkfile" => {
            let path = parts.next().unwrap_or("").to_string();
            Command::Mkfile { path }
        }
        "sysinfo" => Command::Sysinfo,
        "help" | "-h" | "--help" => Command::Help,
        "ver" | "--ver" => Command::Ver,
        "enable" => {
            let function = parts.next().unwrap_or("").to_string();
            Command::Enable { function }
        }
        "disable" => {
            let function = parts.next().unwrap_or("").to_string();
            Command::Disable { function }
        }
        "ctl" => {
            let function = parts.next().unwrap_or("").to_string();
            let settings = parts.next().unwrap_or("").to_string();
            Command::Ctl { function, settings }
        }
        name if is_custom_command(name) => {
            let args = parts.map(|s| s.to_string()).collect();
            Command::Custom { name: name.to_string(), args }
        }
        "" => Command::Unknown(""), 
        _ => Command::Unknown(command_name.to_string()),
    }
}

fn is_custom_command(name: &str) -> bool {
    false 
}