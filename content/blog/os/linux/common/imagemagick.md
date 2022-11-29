## Convert image to PDF
You can use the following command to convert images to a PDF file in current directory:

```shell
docker run --rm -v $(pwd):/imgs dpokidov/imagemagick /imgs/sample.jpeg /imgs/sample.pdf

# or convert multiple images into one PDF file
docker run --rm -v $(pwd):/imgs dpokidov/imagemagick '/imgs/sample*.jpeg' /imgs/sample.pdf
```
