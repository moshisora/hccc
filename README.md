## hccc

command line tools to find undefined or not used classes in html/css file.

## Usage

### Install

```
$ go get github.com/moshisora/hccc
```

### Run and Options

```
$ hccc -h path/to/html_file -c path/to/css_file
```

* `-h` `-html`: Path to the HTML file you want to check
* `-c` `-css`: Path to the CSS file you want to check
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
