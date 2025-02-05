use std::fs;
use std::path::PathBuf;

pub struct History {
    file_path: PathBuf,
    capacity: usize,
    entries: Vec<String>,
}

impl History {
    pub fn new(capacity: usize) -> Self {
        let file_path = Self::get_history_file_path();
        let mut history = History {
            file_path,
            capacity,
            entries: Vec::new(),
        };
        history.load_from_file();
        history
    }

    fn get_history_file_path() -> PathBuf {
        let home_dir = dirs::home_dir().expect("Could not find home directory");
        home_dir.join(".bash_history") 
    }

    fn load_from_file(&mut self) {
        if let Ok(contents) = fs::read_to_string(&self.file_path) {
            self.entries = contents.lines().map(String::from).collect();
        }
    }

    fn save_to_file(&self) {
        let contents = self.entries.join("\n");
        fs::write(&self.file_path, contents).unwrap();
    }

    pub fn add(&mut self, command: &str) {
        if self.entries.len() >= self.capacity {
            self.entries.remove(0); 
        }
        self.entries.push(command.to_string());
        self.save_to_file();
    }

    pub fn get_all(&self) -> Vec<String> {
        self.entries.clone()
    }

    pub fn clear(&mut self) {
        self.entries.clear();
        self.save_to_file();
    }
}