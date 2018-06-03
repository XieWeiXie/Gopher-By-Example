package main

import "go-example-for-live/ten/app"

func main() {
	NewApp := app.NewConfigInstance("2")
	NewApp.Instance("git@github.com:wuxiaoxiaoshen/Resume.git")

}
