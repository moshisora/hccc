## hccc

command line tool to find undefined or not used classes in html/css file.

## Usage

### Install

```
$ go get github.com/moshisora/hccc
```

### Run and Options

```
$ hccc -h path/to/html -c path/to/css
```

* `-h` `-html`: Path to the HTML files you want to check
    - Select a file or directory
	- If a directory is selected, hccc use included files recursively
* `-c` `-css`: Path to the CSS files you want to check
    - Select a file or directory
	- If a directory is selected, hccc use included files recursively
* `-v`: show version
* `-help`: show help

## Example

### input

```html
<!-- sample.html -->
<html>
    <head>...</head>
    <body>
        <div class="class1 class2"></div>
        <div class="class2"></div>
    </nody>
</html>
```

```css
/* sample.css */
.class2 {...}
div.class3 {...}
```

```
$ hccc -h sample.html -c sample.css
```

### output

```
classes used in html but undefined in css
class1
classes defined in css but not used in html
class3
```

## License

The MIT License (MIT)

Copyright 2016~ moshisora
