package models

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	headerSplit = "---"
)

// TPost one post information.
type TPost struct {
	ReqURL            string
	Dir               string
	FileName          string
	Time              time.Time
	Title             string
	Author            string
	Categories        string
	Tags              []string
	ContentWithHeader []byte
	Content           []byte
}

// NewPost create a new post.
func NewPost(Dir, FileName string) *TPost {
	newPost := &TPost{}
	newPost.Dir = Dir
	newPost.FileName = FileName
	newPost.ReqURL = "/posts/" + FileName

	fileBase := path.Base(newPost.FileName)
	fileExt := path.Ext(newPost.FileName)
	beego.Info("dir =", Dir)
	beego.Info("FileName =", newPost.FileName)
	beego.Info("fileBase =", fileBase)
	beego.Info("fileExt =", fileExt)

	fileNameOnly := strings.TrimSuffix(fileBase, fileExt)
	beego.Info("fileNameOnly =", fileNameOnly)

	fileNameParts := strings.Split(fileNameOnly, "-")
	if len(fileNameParts) < 4 {
		// invalid file
		return nil
	}

	strYear := fileNameParts[0]
	strMonth := fileNameParts[1]
	strDay := fileNameParts[2]
	beego.Info("strYear, strMonth, strDay =", strYear, strMonth, strDay)

	iYear, _ := strconv.Atoi(strYear)
	iMonth, _ := strconv.Atoi(strMonth)
	iDay, _ := strconv.Atoi(strDay)

	if iYear > curYear || iYear < 1991 {
		// invalid year
		return nil
	}

	if iMonth < 1 || iMonth > 12 {
		// invalid month
		return nil
	}

	newPost.Time = time.Date(iYear, time.Month(iMonth), iDay, 0, 0, 0, 0, time.Local)

	filePath := filepath.Join(Dir, FileName)
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		beego.Error("read file failed, err = ", err)
		return nil
	}
	newPost.ContentWithHeader = content

	bReader := bytes.NewReader(content)
	ioReader := bufio.NewReader(bReader)

	bEnterHeader := false
	for {
		line, _, err := ioReader.ReadLine()
		if err != nil {
			break
		}

		if string(line) == headerSplit {
			if !bEnterHeader {
				bEnterHeader = true
			} else {
				break
			}
		} else {
			if !bEnterHeader {
				// invalid file
				return nil
			}

			lineParts := strings.Split(string(line), ":")
			if len(lineParts) < 2 {
				continue
			}

			switch lineParts[0] {
			case "layout":
			case "title":
				newPost.Title = lineParts[1]
			case "date":
			case "author":
				newPost.Author = lineParts[1]
			case "comments":
			case "categories":
				newPost.Categories = lineParts[1]
			case "tags":
				tags := strings.Split(lineParts[1], ",")
				newPost.Tags = tags
			case "description":
			case "published":
			}
		}
	}

	if len(newPost.Author) == 0 {
		newPost.Author = "MDGSF"
	}

	contentParts := strings.SplitN(string(content), headerSplit, 3)
	if len(contentParts) != 3 {
		return nil
	}
	// beego.Info("contentParts[0] =", contentParts[0]) // empty
	// beego.Info("contentParts[1] =", contentParts[1]) // header
	// beego.Info("contentParts[2] =", contentParts[2]) // content

	newPost.Content = []byte(contentParts[2])

	return newPost
}
