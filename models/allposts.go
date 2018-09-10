package models

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"

	"github.com/MDGSF/Blog/u"
	"github.com/astaxie/beego"
)

func init() {

	curYear = time.Now().Year()

	AllPosts = make([]*TPost, 0)

	AllPostsFileName = make(map[string]*TPost)

	AllPostsName = make(map[string]*TPost)

	AllPostsTags = make(map[string][]*TPost)

	MonthPosts = make(map[string][]*TPost)

	LoadAllPostsDirectory()
}

var curYear int

// AllPosts store all posts information.
var AllPosts []*TPost

// ByTime sort AllPosts array by post time.
type ByTime []*TPost

func (s ByTime) Len() int      { return len(s) }
func (s ByTime) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByTime) Less(i, j int) bool {
	return s[i].Time.Before(s[j].Time)
}

// AllPostsFileName key: file name, value : TPost.
var AllPostsFileName map[string]*TPost

// AllPostsName key: post name, value : TPost.
var AllPostsName map[string]*TPost

// AllPostsTags key: tag name, value: post array.
var AllPostsTags map[string][]*TPost

// MonthPosts key: year-month, value: post array.
var MonthPosts map[string][]*TPost

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

			newPost := NewPost(fileInfo.Dir, fileInfo.FileInfo.Name())
			if newPost == nil {
				continue
			}

			AllPosts = append(AllPosts, newPost)
		}
	}

	sort.Sort(ByTime(AllPosts))

	for k, post := range AllPosts {
		AllPostsFileName[post.FileName] = AllPosts[k]
		AllPostsName[post.Title] = AllPosts[k]

		for _, tag := range post.Tags {
			if len(tag) == 0 {
				continue
			}

			_, ok := AllPostsTags[tag]
			if !ok {
				AllPostsTags[tag] = make([]*TPost, 0)
			}
			AllPostsTags[tag] = append(AllPostsTags[tag], AllPosts[k])
		}

		_, ok := MonthPosts[post.YearMonth]
		if !ok {
			MonthPosts[post.YearMonth] = make([]*TPost, 0)
		}
		MonthPosts[post.YearMonth] = append(MonthPosts[post.YearMonth], AllPosts[k])
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
