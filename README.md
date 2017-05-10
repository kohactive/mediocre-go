# mediocre-go
A re-imaging of mediocre in go. Image hosting and resizing as a service. Proxy.

## Usage

Pass a URL to an image and resizing attributes to get a new image.

**Attributes**
- URL: Full URL to the image, e.g. https://www.domain.com/path/to/image.jpg
- Width: New image width
- Height: New image height
- Crop Type: `fit` or `fill`

https://resizer.domain.com/[WIDTH]/[HEIGHT]/[CROP_TYPE]/[PATH_TO_IMAGE]

**Example**

https://resizer.domain.com/600/400/fit/https://www.domain.com/path/to/image.jpg
