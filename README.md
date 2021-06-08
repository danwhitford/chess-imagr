# Chess Imagr

```
Usage of chess-imagr:
  -glob string
    	Glob for input files (default "*.pgn")
  -output string
    	Directory to output images (default "out")
```

## Example use

The `-glob` argument accepts a [glob](https://en.wikipedia.org/wiki/Glob_(programming))
for input files.  To read from a single file you can just use the literal path

```
./chess-imagr -glob /path/to/game.pgn
```

or use wildcards to read a bunch of files from a folder

```
./chess-imagr -glob /path/to/*.pgn
```


The `-output` argument defines where the images will be saved. Defaults to `./out` but can
be changed to any path. The directory will be created if it does not exist.
