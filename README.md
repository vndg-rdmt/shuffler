# Shuffler
> Shuffles randomly chosen chunks and draws generated image

### Installation

##### Install:
```bash
git clone git@github.com:vndg-rdmt/shuffler.git
cd shuffler
make
```
> If installation failes or you want to change installation directory, change `BINARY_DIRECTORY` in the Makefile to prefferd

##### Build
```bash
# while in clonned repo
make build

# to also install `make install`...
```

### Usage

To get config schema, use `shuffler.getconfig`
```bash
$ shuffler.getconfig -h
```

To generate image `shuffler`
```bash
$ shuffler -h

 -- manual -------------------------------------------------------------------
 | 'replicator' generates images by randomly selecting chunks from           |
 | the loaded 'chunks' from the source directory, which are just also images |
 | of the (important) same size, png shoud be used.                          |
 -----------------------------------------------------------------------------

-- usage: shuffler [flags] -------------------------------------------------

  -config string
        Flags formated in a file config
  -height uint
        Output image height (default 2)
  -help
        Print out this help message
  -log
        Include logging to stdout (default true)
  -output string
        Output filename (default "replicator.output")
  -source string
        Source directory with the chunks
  -width uint
        Output image width (default 2)
```

### Example

Put chunks (small images of the same size) to a directory, for example i have colors/test_1 directory.
```
$ leaf
.
└─ test_1/
│   ├─ Frame 2.png (0.17 KB)
│   ├─ Frame 3.png (0.51 KB)
│   ├─ Frame 4.png (0.35 KB)
│   ├─ Frame 5.png (0.35 KB)
│   └─ Frame 6.png (0.35 KB)

total 6: (1.73 KB)
```
[leaf](https://github.com/vndg-rdmt/leaf) utilty.

Now run shuffler on this chunks
```bash
$ shuffler -width 20 -height 10 -source ./colors/test_1/
```

![example output image](https://github.com/vndg-rdmt/shuffler/blob/main/docs/output.png?raw=true)


> This code currently is barely tested (tests will be added later) so you can encourage bugs, ub and etc.
> Current functionality is a bare minimum to generate images, all features will be added later on, like generation of some amount of
> images at once instead of only one, images templating, generation rules and etc.
> 
> If you have any ideas, write them to the repo issues.