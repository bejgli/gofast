# Go File Auto SorT &rarr; GoFAST

A tool for automatically sorting my downloads (or other) folder currently in development. 
The main idea is a daemon process that can be configured via json, allowing the user to create advanced sorting patterns with regex.

At the moment it's just a rough proof of concept, a lot is about to change. Here is a list of things I'm definitely adding. (Also subject to changing as I'll probably have new ideas on the Go.)

## Todo
- [ ] Tests
- [ ] Check registered target paths when reading the config file
- [ ] File overwrite option
- [ ] Error handling: the program shouldn't exit on every error
- [ ] Ensure everything is cross-platform
- [ ] Install script