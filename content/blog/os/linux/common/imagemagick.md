## Convert image to PDF
You can use the following command to convert images to a PDF file in current directory:

```shell
docker run -v $(pwd):/imgs dpokidov/imagemagick /imgs/sample.jpeg /imgs/sample.pdf
```
