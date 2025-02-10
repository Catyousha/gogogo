new things learned from each project:
* which
    * `os.getEnv()` can be used to get device env
    * arg0 is current executable path
    * `continue` to skip loop
    * `mode&0111 != 0` == executable file
* phonebook
    * if non-void func is nullable, pointer must be added on returned data type (`func search(key string) *Entry`)
    * open file with `os.OpenFile()`, append to existing file and create if unexist with `os.O_APPEND|os.O_CREATE|os.O_WRONLY` flags (no more usual if-elses)
