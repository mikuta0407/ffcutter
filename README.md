# easy ffmpeg cutter

## install

### go install

```bash
go install github.com/mikuta0407/ffcutter@latest
```

### Download binary

[Releases page](https://github.com/mikuta0407/ffcutter/releases)

## Usage
```
ffcutter -h
  -i, --input string    input file name
  -o, --output string   output file name
  -s, --start string    start time (1:23:45 or 1h23m45s)
  -e, --end string      end time (1:23:45 or 1h23m45s)
  -a, --audio           audio only mode
      --dryrun          dryrun mode (print command only)
  -h, --help            show help message

```

Examples

```bash
ffcutter -i hoge.mp4 -o hoge-cut.mp4 -s 1:01:01 -e 2h02m02s
```


```bash
ffcutter -i hoge.mp4 -o hoge-cut_audio.mp4 -s 1:00:00 -e 2h02m02s -a
```
