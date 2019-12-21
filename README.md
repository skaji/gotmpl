# gotmpl [![](https://github.com/skaji/gotmpl/workflows/test/badge.svg)](https://github.com/skaji/gotmpl/actions)

A CLI for golang template.

## Usage

```
❯ cat text.txt
I'm {{ exec "whoami" | trim }}

❯ gotmpl text.txt
I'm skaji
```

## Available funcsions

* https://masterminds.github.io/sprig/
* exec
* fromYaml
* fromYamlMulti
* readFile
* toYaml

## Credit

Some code taken from https://github.com/roboll/helmfile

## License

MIT

## Author

Shoichi Kaji
