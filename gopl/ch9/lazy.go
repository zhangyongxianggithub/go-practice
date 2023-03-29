package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image
var mu1 sync.Mutex
var mu2 sync.RWMutex

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png": loadIcon("spades.png"),
	}
}

// 并发不安全
func Icon(name string) image.Image {
	mu1.Lock()
	defer mu1.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}
func Icon2(name string) image.Image {
	mu2.RLock()
	if icons != nil {
		icon := icons[name]
		mu2.RUnlock()
		return icon
	}
	mu2.RUnlock()
	mu2.Lock()
	defer mu2.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]

}

var loadIconsOnce sync.Once

func Icon3(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
func loadIcon(name string) image.Image {
	return nil
}
