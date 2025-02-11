new things learned from each project:
* numbers
    * both `new Error("msg")` & `fmt.Errorf("msg")` returns msg string.
    * use `complex64` or `complex128` to present scientific notation

* text
    * printf rune (single-quoted value) with `%s` will show int32 form, use %c instead.

* date
    * `time` library parsing time with predefined format (dd/mm/yyyy -> `time.Parse("02 January 2001", dateString)`). dunno why this looks cursed as f.
    * if input date includes `hh:mm` but parser format only has `"02 January 2001"`, it will return err.

* constants
    * set auto-incremental value to enum by assign `iota` to first enum `const (Zero = iota, One, Two)` == `Zero = 0, One = 1, Two = 2`.
    * `iota` is auto-incremental index starts with `0`, meaning it could be computed (`Zero = iota * 2`, then `One = 4` (2 * 1) )

* slices
    * array = fixed size, slice = dynamic size
    * slices (golang's dynamic array) has preserved memory (capacity) that doubles when appended over initial size. it must capped by it's final length (`s = s[0:len(s):len(s)]`) to prevent unnecessary bloats.

* pointers
    * pointer = reference to variable. function can update variable through pointer `func(*param) { *param = newValue }`
    * other variable can modify referenced variable (`newVar = &oldVar; *newVar = 0; // oldVar=0`)

