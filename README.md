# Go File Auto SorT &rarr; GoFAST

A tool for automatically sorting my downloads (or other) folder currently in development. 
The main idea is a daemon process that can be configured via json, allowing the user to create advanced sorting patterns with regex.

At the moment it's just a rough proof of concept, a lot is about to change. Here is a list of things I'm definitely adding. (Also subject to change as I'll probably have new ideas on the Go.)

## Todo
- [ ] Tests
- [x] Check regex patterns and directories
- [ ] File overwrite option
- [ ] File size limit option
- [ ] Create target directory option
- [ ] Graceful shutdown, signal handling
- [ ] Error handling: the program shouldn't exit on every error
- [ ] Structured logging
- [ ] Ensure everything is cross-platform
- [ ] Usage docs
- [ ] Install script
- [ ] Maybe a simple server for configuration via gui in the browser, so that my parents can also use it :)
