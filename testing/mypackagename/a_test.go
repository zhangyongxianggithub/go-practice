package mypackagename

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestA(t *testing.T) {
	log.Println("TestA running")
}

func TestB(t *testing.T) {
	log.Println("TestB running")
}

// func TestDBFeatureA(t *testing.T) {
// 	defer models.TestDBManager.Reset()
//
// 	// Do the tests
// }
// func TestDBFeatureB(t *testing.T) {
// 	defer models.TestDBManager.Reset()
//
// 	// Do the tests
// }
// func TestMain(m *testing.M) {
// 	models.TestDBManager.Setup()
// 	// os.Exit() does not respect defer statements
// 	code := m.Run()
// 	models.TestDBManager.Exit()
// 	os.Exit(code)
// }
