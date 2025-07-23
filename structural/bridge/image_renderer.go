package bridge

import "fmt"

// ImageRenderer определяем методы для рендеринга изображения.
type ImageRenderer interface {
	Render(data []byte)
}

// PngRenderer конкретный формат картинки PNG.
type PngRenderer struct{}

func (r *PngRenderer) Render(data []byte) {
	fmt.Println("Рендеринг изображения в формате PNG...")
}

// JpgRenderer конкретный формат картинки JPG.
type JpgRenderer struct{}

func (r *JpgRenderer) Render(data []byte) {
	fmt.Println("Рендеринг изображения в формате JPG...")
}

// ImageProcessor рендер для картинок.
type ImageProcessor struct {
	renderer ImageRenderer
}

func NewImageProcessor(renderer ImageRenderer) *ImageProcessor {
	return &ImageProcessor{renderer: renderer}
}

// ApplyFilter фильтр изображения.
func (ip *ImageProcessor) ApplyFilter(filterType string, data []byte) {
	fmt.Printf("Применяем фильтр '%s'...\n", filterType)
	ip.renderer.Render(data)
}

// changeFilter изменение фильтра без влияния на формат изображения.
func changeFilter(ip *ImageProcessor, filterType string, data []byte) {
	ip.ApplyFilter(filterType, data)
}

// changeImplementation изменение реализации без абстракции.
func changeImplementation(ip *ImageProcessor, renderer ImageRenderer) {
	ip.renderer = renderer
}

func StartBridgePattern() {
	pngRenderer := new(PngRenderer)
	jpgRenderer := new(JpgRenderer)
	imageData := make([]byte, 1024)

	processor := NewImageProcessor(pngRenderer)
	changeFilter(processor, "Grayscale", imageData)

	// меняем формат изображения, оставаясь на уровне абстракции.
	changeImplementation(processor, jpgRenderer)
	changeFilter(processor, "Sepia", imageData)
}
