package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// go run . grep => /usr/bin/grep
func main()  {
	// parse args on execute ./go run . <args1> <argsN..>
	arguments := os.Args;

	// args0 == executable file
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return;
	}

	// check if file is within '$PATH' dirs
	path := os.Getenv("PATH");
	pathSplit := filepath.SplitList(path);

	for _, dir := range pathSplit {
		// exercise 2: accept multiple args
		for _, file := range arguments {
			fullpath := filepath.Join(dir, file);
			
			fileInfo, err := os.Stat(fullpath)
			if(err != nil) {
				// skip iter if not exist
				continue
			}
	
			mode := fileInfo.Mode()
			// is regular file?
			if mode.IsRegular() {
				// is executable?
				if mode&0111 != 0 {
					fmt.Println(fullpath)
					// exercise 1: find all occurences ('return' omitted from code example)
					// return
				}
			}
		}
	}

}