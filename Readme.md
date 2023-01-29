# Ffuf: Post processing
Unfortunately - despite its "-ac" flag, ffuf tends to produce a lot of irrelevant entries. 
This is why I've created a post-processing tool to filter out those entries. This tool has to be
run after ffuf has finished. Additionally, the initial ffuf command should be run with the following flags:

```
-o /folder/to/results.json
-od /folder/to/bodies
-of json (default)
```

This forces ffuf to write a summary file in json format as well as bodies of the responses to disk. 
Adding "-od" is not mandatory but recommended.

I highly recommend NOT to use the "-ac" flag - especially if you don't want to miss cool stuff and want to use this
post-processing tool.

## Usage

```
Usage of ./ffufPostprocessing:
  -result-file string
        Path to the original ffuf result file
  -bodies-folder string
        Path to the ffuf bodies folder (optional, if set results will be better)
  -new-result-file string
        Path to the new ffuf result file (optional)
  -delete-bodies
        Delete unnecessary body files after filtering (optional)
  -delete-bodies
        Delete all body files after filtering (optional)
  -overwrite-result-file
        Overwrite original result file (optional)
  -verbose
        Verbose mode (Shows currently filtered results) (optional)
```

## Details

Especially when -od is set, which means we have all http headers and bodies for each requested url - this tool will initially
analyse all bodies and enrich the initial results json file with the following data points:

- count of all headers
- domain of redirect if applicable
- amount of parameters in redirect if applicable
- length and words of page title (if existent)
- count of detected css files
- count of detected js files
- count of tags in html/xml/json (calculation is wild)

Afterwards it will scan the entire new results file and keep only those entries which are unique based on known metadata types.
If it turns out that one of those values is always different (e.g. the title of pages can vary very much) - this metadata type 
will be skipped for the uniqueness check.

In general this tool will always keep a small amount of entries which are _not_ unique. For example, if the results json file
contains 300x http status 403 (with size, length, ... identical) and 2 unique http status 200 responses, it won't drop all 300 http status 403 entries. 
It will keep X of them in the data set.

## Install

ffufPostprocessing requires golang 1.19

### Build from source

```
cd ffufPostprocessing
go build -o dist/ffufPostprocessing main.go
```

## License

I don't care. Do whatever you want with this tool.