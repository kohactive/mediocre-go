package main

import "net/http"
import "strings"
import "strconv"
//import "log"
import "io/ioutil"
import "github.com/go-martini/martini"
import "github.com/gographics/imagick/imagick"

func main() {
  m := martini.Classic()
  
  imagick.Initialize()
  defer imagick.Terminate()
 
  m.Get("/:width/:height/:mode/**", func(params martini.Params, res http.ResponseWriter){
    
    mode := params["mode"]

    if (!strings.EqualFold(mode, "fit") && !strings.EqualFold(mode, "fill")) {
      res.WriteHeader(http.StatusNotAcceptable)
      res.Write([]byte("You must pass mode as 'fit' or 'fill', /width/height/[fit or fill]/url of image"))
      return
    }

    resp, _ := http.Get(params["_1"])
    body, _ := ioutil.ReadAll(resp.Body)

    
    twidth, _ := strconv.ParseUint(params["width"], 10, 32)
    theight, _ := strconv.ParseUint(params["height"], 10, 32)

    width := uint(twidth)
    height := uint(theight)

    mw := imagick.NewMagickWand()
    mw.ReadImageBlob(body)
    
    // calculate current aspect ratio of image
    ow, oh, _, _, _ := mw.GetImagePage()
    portrait := true
    
    if (ow > oh) {
      portrait = false
    }

    aspectRatio := float32(ow) / float32(oh)


    if strings.EqualFold(mode, "fit") {

      if (width < height) {
        portrait = false
      } else {
        portrait = true
      }

      if portrait {
        width = uint(float32(height) * aspectRatio)
      } else {
        height = uint(float32(width) * (1.0/aspectRatio))
      }

      mw.ResizeImage(width, height, imagick.FILTER_LANCZOS, 1)
    } else {

      if (width > height) {
        portrait = false
      } else {
        portrait = true
      }

      cropwidth := width
      cropheight := height

      if portrait {
        width = uint(float32(height) * aspectRatio)
      } else {
        height = uint(float32(width) * (1.0/aspectRatio))
      }
      
      mw.ResizeImage(width, height, imagick.FILTER_LANCZOS, 1)
      
      centerx := int((width / 2) - (cropwidth / 2))
      centery := int((height / 2) - (cropheight / 2))

      mw.CropImage(cropwidth, cropheight, centerx, centery)


    }

    resizedImage := mw.GetImageBlob()
    
    res.Write(resizedImage)
    
    defer mw.Destroy()

    // return params["width"] + "x" + params["height"] + "@" + params["_1"]
  })

  m.Run()
}

// apphost/300/200/www.kohactive.com/logo.png