use std::collections::HashMap;
use std::env;
use std::fs;
use std::io::{self, Write};
use std::path::Path;

type BuiltinFn = fn(&[String]) -> Result<(), String>;

pub struct Builtins {
    functions: HashMap<String, BuiltinFn>,
}

impl Builtins {
    pub fn new() -> Self {
        Builtins {
            functions: HashMap::new(),
        }
    }

    pub fn add(&mut self, name: &str, func: BuiltinFn) {
        self.functions.insert(name.to_string(), func);
    }

    pub fn execute(&self, args: &[String]) -> Result<(), String> {
        if args.is_empty() {
            return Err("No command specified".to_string());
        }

        let name = &args[0];
        let func = self.functions.get(name);

        match func {
            Some(f) => f(&args[1..]),
            None => Err(format!("Command not found: {}", name)),
        }
    }

    pub fn help(&self) {
        println!("Available built-in commands:");
        for (name, _) in &self.functions {
            println!("  {}", name);
        }
    }
}

fn cd(args: &[String]) -> Result<(), String> {
    let path = match args.len() {
        0 => env::var("HOME").unwrap_or_else(|_| "/".to_string()),
        1 => args[0].clone(),
        _ => return Err("cd: too many arguments".to_string()),
    };

    if let Err(e) = env::set_current_dir(path) {
        return Err(format!("cd: {}", e));
    }

    Ok(())
}

fn pwd(_args: &[String]) -> Result<(), String> {
    let path = env::current_dir().map_err(|e| format!("pwd: {}", e))?;
    println!("{}", path.display());
    Ok(())
}

fn echo(args: &[String]) -> Result<(), String> {
    for arg in args {
        print!("{} ", arg);
    }
    println!();
    Ok(())
}

fn ls(args: &[String]) -> Result<(), String> {
    let path = match args.len() {
        0 => ".",
        1 => &args[0],
        _ => return Err("ls: too many arguments".to_string()),
    };

    let entries = fs::read_dir(path).map_err(|e| format!("ls: {}", e))?;

    for entry in entries {
        let entry = entry.map_err(|e| format!("ls: {}", e))?;
        println!("{}", entry.file_name().to_string_lossy());
    }

    Ok(())
}

fn cat(args: &[String]) -> Result<(), String> {
    if args.is_empty() {
        return Err("cat: missing operand".to_string());
    }

    for arg in args {
        let path = Path::new(arg);
        let contents = fs::read_to_string(path).map_err(|e| format!("cat: {}: {}", arg, e))?;
        print!("{}", contents);
    }

    Ok(())
}

fn help(_args: &[String]) -> Result<(), String> {
    println!("Type `help` to see available commands.");
    Ok(())
}

pub fn register_builtins(builtins: &mut Builtins) {
    builtins.add("cd", cd);
    builtins.add("pwd", pwd);
    builtins.add("echo", echo);
    builtins.add("ls", ls);
    builtins.add("cat", cat);
    builtins.add("help", help);
}