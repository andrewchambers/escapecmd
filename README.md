# escapecmd

Expand a set of arguments into a properly escaped bash string or json array.

```
$ ./escapecmd -json "a b" b c
"[\"a b\",\"b\",\"c\"]"
$ ./escapecmd -json echo "a b" b c
"[\"echo\",\"a b\",\"b\",\"c\"]"
$ ./escapecmd  echo "a b" b c
```


This is useful for commands that need to take multiple subcommands as arguments,
without forcing the end user to mess around with escape strings.

```
dosomething -subcmd1 $(escapecmd foo "a b") -subcmd2 $(escapecmd c "d e" f)
```