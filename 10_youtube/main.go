package main
import (
	"fmt"
)

type YouTubeVedio struct{
	CloudStorageID int
	Title	string
	Description string
	URL	 string
}

type TaskYoutube struct {
	TaskID int
	TaskItem []YouTubeVedio
}

type YouTubeVedioResult struct{
	CloudStorageID int
	Status int
}

type TaskYoutubeResult struct {
	TaskID int
	TaskItem []YouTubeVedioResult
}
type SocialMedia struct{
	UserName	string
	MessagePosted	string

}

type TaskSocialMedia struct {
	TaskID	int
	TaskItem SocialMedia
}

func main() {
	fmt.Println("Hello, playground")
	mTasks := TaskYoutube{
		TaskID:123,
		TaskItem:[]YouTubeVedio{
			{CloudStorageID:111, Title:"hello youtube", Description:"xxx", URL:"www.google.com"},
			{CloudStorageID:222, Title:"hello youtube", Description:"xxx", URL:"www.google.com"},
		},
	}

	fmt.Println(mTasks)
	mTaskResults := TaskYoutubeResult {
		TaskID:123,
		TaskItem: []YouTubeVedioResult{
			{CloudStorageID:111, Status: 0},
			{CloudStorageID:222, Status: 1},
		},
	}
	fmt.Println(mTaskResults)

	mTasksMedia := TaskSocialMedia {
		TaskID:123,
		TaskItem: SocialMedia{UserName: "John_Lee", MessagePosted:"Hello Twitter "},
	}
	fmt.Println(mTasksMedia )
}
