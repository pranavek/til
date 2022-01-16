# Double dash in commands 
I created a file named `--test` by mistake and tried to delete it using `rm --test`. This won't work as `rm` commands treats the `--test` as command option. Use `--` to signal end of command option.
```bash
$ rm -- --test
```

## Resources and References
[baeldung.com](https://www.baeldung.com/linux/double-dash-in-shell-commands)
