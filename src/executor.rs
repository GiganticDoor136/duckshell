use std::process::{Command, Stdio};
use std::io::{self, Read};

pub struct Executor;

impl Executor {
    pub fn new() -> Self {
        Executor
    }

    pub fn execute(&self, command: &str, args: &[&str]) -> Result<String, String> {
        let mut cmd = Command::new(command);
        cmd.args(args)
            .stdout(Stdio::piped())
            .stderr(Stdio::piped());

        let mut child = match cmd.spawn() {
            Ok(child) => child,
            Err(e) => return Err(format!("Failed to execute command: {}", e)),
        };

        let mut stdout = String::new();
        let mut stderr = String::new();

        match child.stdout.take().unwrap().read_to_string(&mut stdout) {
            Ok(_) => (),
            Err(e) => return Err(format!("Failed to read stdout: {}", e)),
        }

        match child.stderr.take().unwrap().read_to_string(&mut stderr) {
            Ok(_) => (),
            Err(e) => return Err(format!("Failed to read stderr: {}", e)),
        }

        let status = match child.wait() {
            Ok(status) => status,
            Err(e) => return Err(format!("Failed to wait for command: {}", e)),
        };

        if status.success() {
            Ok(stdout)
        } else {
            Err(format!("Command exited with status {}: {}", status, stderr))
        }
    }
}