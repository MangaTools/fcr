# Folder creator (fcr)
As the name implies, this program generates folders according to specified parameters

## Functionality

At the moment, the application knows how to create n number of folders with names according to a pattern.

The template includes:
 - plain text
 - template for inserting the number with start index, incrementing step, and padding

## Install

You can [build](#Build) app from source or download precompiled binaries from [releases](https://github.com/ShaDream/folder_creator/releases)

## Build

In order to build an application, you need to install:
 - go 1.18 or higher

Then clone this repository:

`git clone https://github.com/ShaDream/folder_creator`

And in root directory of app execute:

 - Windows

`go build -o fcr.exe`

- Linux

`go build -o fcr && chmod +x fcr`

Then you can place that executable in folder listed in PATH variable

## License

see [`LICENSE`](https://github.com/ShaDream/folder_creator/blob/main/LICENSE)
