# Linux/Unix File System Navigation

Basic Navigation
* Navigate to a path: cd <path>
* Navigate to Home directory: `cd`, `cd ~`
* Navigate to Previous directory: `cd -`
* Navigate to Parent directory: `cd ..`
* Navigate to Root directory: `cd /`
 
Print working directory: `pwd`
Type of paths:
* Absolute paths: That start from root directory (`/`)
* Relative paths: That start from current working directory.

Absolute paths start with `/`

Relative paths start with one of:
* a directory name present inside the Current working Directory
* `.` : a dot indicating relative to itself
* `..`: double dot indicating parent of the current working directory
(Yeah... to reference a file that is present directly in cwd, we need to do something like `./<filename>` since relative paths can't start with filename)


Tags:
	#Linux #Navigation #Unix

