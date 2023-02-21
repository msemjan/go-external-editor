# Demo: External Text Editor From Go

This repository shows how to open an External Text Editor from Go application.
Why would this be useful? When editing text, I prefer a well known environment
of my Vim text editor with all of its plugins and my personalized settings.
Achieving a similar experience in a hand-crafted  application would be not
easy.

This app creates a temporary file, fills it with "template" data, and then lets
user to modify it.

Possible use-cases of using an external editor from your Go App:
1. Editing your configuration files from within the app (maybe with
   validation of the settings that were inputted by the user)
2. Create files from templates and let the user modify them
3. For longer inputs from the user, e.g. a [diary with encryption](https://jrnl.sh/en/stable/)
4. When your application just finds/manages files for the user, but their
   editing is delegated to an external editor.

## How to run this example? 

1. Clone this repository
2. Run the app `go run main.go`
