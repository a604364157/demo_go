package models

import (
	"fmt"
	"strings"
)

func HandleTags(tags []string) map[string]int {
	tagsMap := make(map[string]int)
	for _, tag := range tags {
		tagList := strings.Split(tag, "&")
		for _, value := range tagList {
			tagsMap[value]++
		}
	}
	return tagsMap
}

func QueryArticleTags() []string {
	var tags []string
	rows, err := getDb().Raw("select tags from article").QueryRows(&tags)
	fmt.Println(rows, err)
	return tags
}
