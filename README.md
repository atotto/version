# version utility

install:
```
$ go install github.com/atotto/version@latest
```

## version up

update patch version:

```
$ echo 0.1.0 > VERSION
$ version up ./VERSION
0.1.1
```

write result to file instead of stdout:

```
$ version up -w ./VERSION
$ cat ./VERSION
0.1.1
```

update minor version:

```
$ version up -minor ./VERSION
0.2.0
```

update major version

```
$ version up -major ./VERSION
1.0.0
```
