https://gobyexample.com/command-line-subcommands 

Some command-line tools, like the go tool or git have many subcommands, each with its own set of flags. For example, go build and go get are two different subcommands of the go tool. The flag package lets us easily define simple subcommands that have their own flags.

```bash

go build command-line-subcommands.go 



First invoke the foo subcommand.

$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]
Now try bar.

$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]
But bar won’t accept foo’s flags.

$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
        level
```