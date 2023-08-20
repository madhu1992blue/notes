# Linux File System Navigation and Permissions

Listing files in current directory: `ls`  
List directory/file name `foo`: `ls foo`  
Long Listing: `ls -l` (Shows file type, permissions, last modified timestamps, etc)  
Show hidden files: `ls -a` (Shows hidden files/dirs too along with others. Hidden files/dirs are files/dir names that begin with `.`)


Common permissions
1. Read(r): For file, can read contents. On Directory, allows list the file names. 
1. Write(w): For file, can write contents. On Directory, allows add/remove entries in the directory.
1. Execute(x): Can run the file. On Directory, allows navigate to the directory.  
   For a file contents to be read, all it's parent directories should be executable.

What are the identities that can be granted these permissions?
1. user/owner : `u`
1. group: `g`
1. others(rest of the world): `o`

Assume `who` to be any of the identity (user, group and others) : `u`,`g`,`o`, `ug`, `go`, `uo`, `ugo`  
Assume `p` to be any of the permission combination: `r`,`w`,`x`, `rw`, `rx`, `wx`, `rwx`.  
Granting permission combination `p`: `chmod <who>+<p> <path>`  
Removing permission combination `p`: `chmod <who>-<p> <path>`  
Set to absolute permission `p`: `chmod <who>=<p> <path>`  (Ex: If p is just `r`, write and execute will be cleared) 

Combining chmod on a path: Comma seperated ACL management can be done.   
Ex: chmod ug+x,o-rwx foo  
If <who> is not specified, it will assume who as `ugo` 


Isn't there an Absolute/octal notation?   
`chmod <octal> <path>` : Yes, there is but don't use it as it's easy to get it wrong :) 

How do I recursively apply a permission on a directory? Use `-R` or `-r` for recursive.

#linux #basic #basics #permissions 
