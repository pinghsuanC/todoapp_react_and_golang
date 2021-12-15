# todoapp_react_and_golang

- Credit: Following [this youtube tutorial](https://www.youtube.com/channel/UC_XQE5LEqCdgC-_gdlqsqfQ); Thank you Steve-kaufman :).
- Goal is to get familiar with golang syntax and its code structure; later I'd like to use golang for personal website.
- The backend will be "overly-complicated" and focus on extensibility

# Setting up backend with go:

1. command: `go mod init github.com/pinghsuanC/todoapp_react_and_golang/backend`
   - note that it's usually named by github directory even if it's not going to be pushed to the github
   - $GOPATH is another story...
2. project structure
   - `usecases`: the package that orchestrates everything; knows where the resources are and where to send them
   - `test files`
   - `entities`: store the types; public fields have to be capitalized, else they won't be exported
   - access the types by entites.'type' (e.g. entities.Todo)
3. other notes

   - in Go, errors are plain values. There is no throw/catch types of things in Java
   - note useful packages:
     - reflect
     - testing
     - http package
     - http testing package
   - need to checkout golang errors

# TODO：

- goLang: read http and httptest package
- connect react and go backend
- transfer data from go to react
- rightnow don't need to use server, local memory is fine
