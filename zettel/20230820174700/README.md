# Basic Linux User Management

Common Linux User Management operations:
1. Creating an account: `useradd`
1. Modifying an account: `usermod`
1. Deleting an account: `userdel`

Creating an account considerations:
* Should we create a home directory for this user? For service accounts, usually we don't need.
* Should we explicitly set the UID? If there are shared volumes across multiple hosts. Same UID across multiple hosts has to be ensured. It is needed then to explicitly set the UID.
* Should we customize the home directory path? If the home directories reside at non-standard location based on some policy.
* Should a group with same username be created?
* Primary and Supplemental groups that the user should belong to.
* Should the user be able to login ? Most Service accounts shouldn't be able to login. Set shell value to /sbin/nologin or /bin/false to prevent login.

Deleting an account considerations: (Security risk)
* Deleting an account doesn't delete the files owned by the account by default.  
  The ownership of these files will still be associated with same UID though no such account exists on the current machine.
* Deleting an account is usually a security risk as when a new account is added, it may reuse an old user id and any files owned by the user iid will now be owned by the new account.

Where will user info be stored ? /etc/passwd (Don't ever edit this file manually. Use one of the commands or a service that makes system calls to modify this)

How do I set password ?
* `passwd` - Set current account password
* `sudo passwd foo` - Set password of account named, foo.

One of the columns in /etc/passwd is for password but it is marked as `x`.  
* Linux usually stores hashed passwords in a shadow file called /etc/shadow

#linux #user #management #security #risk  

