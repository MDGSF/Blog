package models

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/MDGSF/Blog/u"
	"github.com/astaxie/beego"
)

func init() {

	AllPosts = make([]*TPost, 0)

	AllPostsFileName = make(map[string]*TPost)

	AllPostsName = make(map[string]*TPost)

	AllPostsTags = make(map[string]*TPost)

	LoadAllPostsDirectory()
}

// AllPosts store all posts information.
var AllPosts []*TPost

// AllPostsFileName key: file name, value : TPost.
var AllPostsFileName map[string]*TPost

// AllPostsName key: post name, value : TPost.
var AllPostsName map[string]*TPost

// AllPostsTags key: tag name, value : TPost.
var AllPostsTags map[string]*TPost

// LoadAllPostsDirectory load all posts from all PostDirectory.
func LoadAllPostsDirectory() {
	postDirs := beego.AppConfig.Strings("PostDirectory")
	beego.Info("postDirs =", postDirs)
	for _, postDir := range postDirs {
		if !u.IsDir(postDir) {
			continue
		}

		LoadOnePostDirectory(postDir)
	}
}

var sema = make(chan struct{}, 20)
var done = make(chan struct{})

// TFileInfo used in inner to transfer file information.
type TFileInfo struct {
	Dir      string
	FileInfo os.FileInfo
}

// LoadOnePostDirectory load all posts from dir directory.
func LoadOnePostDirectory(dir string) {

	fileInfoChan := make(chan TFileInfo)
	var n sync.WaitGroup
	n.Add(1)
	go walkDir(dir, &n, fileInfoChan)
	go func() {
		n.Wait()
		close(fileInfoChan)
	}()

loop:
	for {
		select {
		case fileInfo, ok := <-fileInfoChan:
			if !ok {
				break loop
			}

			newPost := &TPost{}
			newPost.FileName = fileInfo.FileInfo.Name()
			fileBase := path.Base(newPost.FileName)
			fileExt := path.Ext(newPost.FileName)
			beego.Info("dir =", fileInfo.Dir)
			beego.Info("FileName =", newPost.FileName)
			beego.Info("fileBase =", fileBase)
			beego.Info("fileExt =", fileExt)

			fileNameOnly := strings.TrimSuffix(fileBase, fileExt)
			beego.Info("fileNameOnly =", fileNameOnly)

			fileNameParts := strings.Split(fileNameOnly, "-")
			if len(fileNameParts) < 4 {
				// invalid file
				continue
			}

			strYear := fileNameParts[0]
			strMonth := fileNameParts[1]
			strDay := fileNameParts[2]
			beego.Info("strYear, strMonth, strDay =", strYear, strMonth, strDay)

			iYear, _ := strconv.Atoi(strYear)
			iMonth, _ := strconv.Atoi(strMonth)
			iDay, _ := strconv.Atoi(strDay)

			// check valid

			time.Date(iYear, time.Month(iMonth), iDay, 0, 0, 0, 0, time.Local)
		}
	}
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func walkDir(dir string, n *sync.WaitGroup, fileInfoChan chan<- TFileInfo) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileInfoChan)
		} else {
			fileInfoChan <- TFileInfo{Dir: dir, FileInfo: entry}
		}
	}
}

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		beego.Error("read dir failed, err =", err)
		return nil
	}
	return entries
}
