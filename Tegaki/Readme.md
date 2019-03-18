# README

## System Requirements
( Golang:1.11, Docker )
# Run
### `$ docker-compose up`
##### End points:
- `http://localhost:8080/` Welcome
* `http://localhost:8080/image` upload an image for processing:
  - Example ` curl -i -X POST -F 'img=@your_image.png' http://localhost:8080/image`

- `http://localhost:8080/image/thumbnail/:id` disply the thumbnail