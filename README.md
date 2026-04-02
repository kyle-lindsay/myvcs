# myvcs

A very simple version-control system written in the Go programming language. It allows a user to initialise a repository, commit changes, and revert to older versions.


## How To Use

**Initialise a repository with** `myvcs init`

This will create a `.myvcs/` directory to store commit history.

**Create a commit with** `myvcs commit "<message>"`

This will create a snapshot of your project folder in it's current state.

**View commit history with** `myvcs log`

This will display the commit id, timestamp and message of all previous commits, showing the most recent commits first.

**Revert to a previous state with** `myvcs checkout <commit id>`

This will revert your project folder to the snapshot stored in the commit with the specified id.

_**Note: Project must be initialised with**_ `myvcs init` _**before any commits are made, and all files in the initialiesd directory will be tracked.**_