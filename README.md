# escapecmd

Expand a set of arguments into a properly escaped bash string or json array.

```
$ escapecmd cat *.txt
"cat foo.txt \"a test.txt\""
$ escapecmd -json cat *.txt
"[\"cat\",\"another.txt\",\"a test.txt\"]"
```

This is useful for commands that need to take multiple subcommands as arguments,
without forcing the end user to mess around with escape strings.

```
$ dosomething -subcmd1 $(escapecmd foo *.txt) -subcmd2 $(escapecmd find ~/bar)
```