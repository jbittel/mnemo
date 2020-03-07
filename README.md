# mnemo

Converts integers into mnemonic "words" and back again.

A port of [mnemo](https://github.com/flon-io/mnemo) from C to Go.

## Usage

```bash
$ mnemo -encode 999
chozo
$ mnemo -encode -999
xachozo
$ mnemo -decode chozo
999
$ mnemo -decode xachozo
-999
```

## License

BSD, see [LICENSE](LICENSE)
