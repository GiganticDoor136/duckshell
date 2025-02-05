use std::env;
use std::path::PathBuf;

pub fn get_home_dir() -> Option<PathBuf> {
    match env::var("HOME") {
        Ok(path) => Some(PathBuf::from(path)),
        Err(_) => None,
    }
}

pub fn get_config_file_path(filename: &str) -> Option<PathBuf> {
    if let Ok(config_dir) = env::var("DSH_CONFIG_DIR") {
        let path = PathBuf::from(config_dir).join(filename);
        if path.exists() {
            return Some(path);
        }
    }

    if let Some(home_dir) = get_home_dir() {
        let path = home_dir.join(".config").join("dsh").join(filename);
        if path.exists() {
            return Some(path);
        }
    }

    let path = PathBuf::from("dsh").join(filename);
    if path.exists() {
        return Some(path);
    }

    None
}

pub fn read_file_to_string(path: &PathBuf) -> Result<String, std::io::Error> {
    std::fs::read_to_string(path)
}