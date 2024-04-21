# What is this?
yyt (pronounced yeet) is a small program that basically keeps track of files you wish to copy or move at a later time.  
This project was created as a way for me to get acquainted with a programming language as this project covers a lot of  
areas in programming languages including file systems, File IO, a programming language's standard library and so on.

# How to use?
`yyt add <file1> <file2> ...` adds specific file(s) to the buffer without affecting the original files  
`yyt ls` lists all files currently in the buffer  
`yyt mv <file1> <file2> ...` moves specific file(s) from one place to the other  
`yyt cp <file1> <file2> ... <dir>` copies specified file(s) to the directory. Another way to achieve this  
is by omitting the `cp` flag altogether
`yyt rm <file1> <file2> ...` removes specified file(s) from the buffer without deleting the original files.  
`yyt clean` removes all invalid or missing files from the cache
