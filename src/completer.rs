use std::path::Path;
use std::fs;

use rustyline::completion::{Completer, Pair};
use rustyline::Context;

pub struct DuckCompleter;

impl Completer for DuckCompleter {
    type Candidate = Pair;

    fn complete(&self, line_part: &str, cursor_position: usize, _context: &Context<'_>) -> rustyline::Result<(usize, Vec<Pair>)> {
        let (start, word) = self.find_word_boundary(line_part, cursor_position);
        let candidates = self.find_candidates(word);
        Ok((start, candidates))
    }
}

impl DuckCompleter {
    fn find_word_boundary(&self, line: &str, position: usize) -> (usize, &str) {
        let mut start = position;
        while start > 0 && !line.as_bytes()[start - 1].is_ascii_whitespace() {
            start -= 1;
        }
        (&line[start..position], start)
    }

    fn find_candidates(&self, word: &str) -> Vec<Pair> {
        let mut candidates = Vec::new();

        // Поиск команд
        let paths = env::var_os("PATH").unwrap_or_default();
        for path in env::split_paths(&paths) {
            if let Ok(entries) = fs::read_dir(path) {
                for entry in entries {
                    if let Ok(entry) = entry {
                        let file_name = entry.file_name();
                        let name = file_name.to_string_lossy();
                        if name.starts_with(word) {
                            candidates.push(Pair {
                                display: name.to_string(),
                                replacement: name.to_string(),
                            });
                        }
                    }
                }
            }
        }

        // Поиск файлов в текущем каталоге
        if let Ok(entries) = fs::read_dir(".") {
            for entry in entries {
                if let Ok(entry) = entry {
                    let file_name = entry.file_name();
                    let name = file_name.to_string_lossy();
                    if name.starts_with(word) {
                        candidates.push(Pair {
                            display: name.to_string(),
                            replacement: name.to_string(),
                        });
                    }
                }
            }
        }

        candidates
    }
}