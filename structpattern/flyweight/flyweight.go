package flyweight

import "fmt"

// 享元工厂类
type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

var imageFactory *ImageFlyweightFactory
func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = new(ImageFlyweightFactory)
		imageFactory.maps = map[string]*ImageFlyweight{}
	}

	return imageFactory
}

func (f *ImageFlyweightFactory) Get(fileName string) *ImageFlyweight {
	image := f.maps[fileName]
	if image == nil {
		image = NewImageFlyweight(fileName)
		f.maps[fileName] = image
	}

	return image
}

// ===========================

// 享元类
type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	data := fmt.Sprintf("image data %s", filename)

	imageFlyweight := new(ImageFlyweight)
	imageFlyweight.data = data
	return imageFlyweight
}

func (i *ImageFlyweight)Data() string{
	return i.data
}

// ===========================

// 享元具体类
type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	imageViewer := new(ImageViewer)
	imageViewer.ImageFlyweight = image

	return imageViewer
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}