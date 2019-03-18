package process

import (
	"bufio"
	"bytes"
	"context"
	"image-service/db"
	"image-service/models"
	"image-service/utils"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func ResizeImageFile(payload []byte) {
	buf := bytes.NewBuffer(payload)
	s := bufio.NewScanner(buf)
	s.Scan()
	imageName := s.Text()

	// lenght of the image name and "\n"
	imgStartIndex := len([]byte("\n")) + len(s.Bytes())
	// extract image content from the payload exculding the image name and "\n"
	imgPayload := payload[imgStartIndex:]

	// Decode the payload to png imnage
	imgDecoded, err := png.Decode(bytes.NewReader(imgPayload))
	utils.HandleError(err)

	resizedImage := resize.Resize(100, 100, imgDecoded, resize.Lanczos3)

	outputFile, err := os.Create("imgs/" + imageName)
	utils.HandleError(err)
	defer outputFile.Close()

	// write new image to file
	png.Encode(outputFile, resizedImage)

	ctx := context.Background()
	// store image status to the database
	db.StoreImageRequestState(ctx, models.ImageRequestState{imageName, true})
}
