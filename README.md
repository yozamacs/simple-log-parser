# Simple Log parser

## Usage
Run `go build` to create an executable

Run the executable in the directory as follows
`./logParser -startTime <start time> -endTime <end time> -logFiles <logfile_1>,<logfile_2>,<logfile_infinity>`

The start time and end time must be in seconds since epoch

The sample log file is provided and can be run with the command
`./logParser -startTime 1493969101.645 -endTime 1493969101.655 -logFiles log_sample.txt`
with the output as follows:

```
Between time 1.493969101645e+09 and 1.493969101655e+09:
player.vimeo.com returned 33.33333333333333% 5xx errors
vimeo.com returned 33.33333333333333% 5xx errors
```

## Dependencies
This is a go binary with no external dependencies outside of the go standard library.
