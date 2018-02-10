package main

import (
	"flag"
	"fmt"
	"reflect"
	"github.com/golang/glog"

)
type YouTubeVedio struct{
	GCloudBucketname		string
	GCloudFileName			string
	Title					string
	CategoryId				string
	Description 			string
	CourseURL	 			string
}
// 避免没有引用fmt的编译错误
var _ = fmt.Println

func main() {
	//
	flag.Parse()

	glog.Info("hello, glog")
	y := YouTubeVedio {
		GCloudBucketname:"apistoragetest",
		GCloudFileName: "my_cute_boy.mp4",
		Title:"hello youtube",
		CategoryId: "21",
		Description:"GreaterCommons",
		CourseURL:"www.google.com",
	}
	fmt.Println(reflect.TypeOf(y))
	judgmentType(1, 2.2, "learing", y)



	glog.Flush()
}

func judgmentType(items ...interface{}) {
	for k, v := range items {
		switch v.(type) {

		case YouTubeVedio:
			fmt.Printf("Student1, %d[%v]\n", k, v)
			value := reflect.ValueOf(v)
			vv := value.FieldByName("GCloudFileName")
			fmt.Println("vv:", vv)
		case *YouTubeVedio:
			fmt.Printf("Student2, %d[%p]\n", k, v)
		}
	}
}