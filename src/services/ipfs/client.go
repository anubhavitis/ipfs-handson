package IPFS

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell
var hashmap map[string]string

const TempStorage = "./uploaded_files"

func init() {

	// connecting with local node of ipfs.
	sh = shell.NewShell("localhost:5001")

	// This will lead to data loss in server restart
	// can be easily handled by using redis or other db.
	hashmap = make(map[string]string)

	// Making a directory that will store temp data
	os.MkdirAll(TempStorage, os.ModePerm)
}

func AddFile(fileName string, file *multipart.File) (string, error) {
	fmt.Println("file name is:", fileName)
	// Add the file to IPFS
	cid, err := sh.Add(*file)
	if err != nil {
		return "", err
	}

	hashmap[cid] = fileName
	return cid, nil
}

func GetFile(cid string) (string, error) {

	fileName := fmt.Sprint(TempStorage, "/", hashmap[cid])
	content, err := sh.Cat(cid)
	if err != nil {
		return fileName, err
	}
	defer content.Close()

	// Create the file with the given name
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return fileName, err
	}
	defer file.Close()

	// Copy the content from IPFS to the local file
	_, err = io.Copy(file, content)
	if err != nil {
		return fileName, err
	}

	return fileName, nil
}
