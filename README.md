# file-dates-cli
CLI for cleaning up filenames with dates

## Usage

For now the project contains a very simple CLI which prints `Hello, world!`.
It requires Go 1.23 or later. You can run it with:

```bash
go run main.go
```

This is a starting point that will be extended to support more commands.

## goals

commands for
- reading in a file name from STDIN containing a date in a variety of possible formats (`dd-mm-yyyy`, `mm-dd-yyyy`, `yy-mm-dd`, etc) and outputing the corrected filename which replaces the existing date with a date in ISO 8601 format (`yyyy-mm-dd`), or exits with an error
- taking in a list of filenames containing a date in ISO 8601 format and filtering the list by day of the week (i.e. Monday)

docs for
- how to use these commands together with linux commands for batch renaming files
- how to use these commands together with linux commands for organizing files
