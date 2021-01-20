func Native(imageSource string) (string) {
	fmt.Println("Start Downloading Image")
	cmd, err := exec.Command("/bin/sh", "copy_image.sh", imageSource, "/Users/mam/Documents/test/sample").Output()
    if err != nil {
    fmt.Printf("error %s", err)
	}
	output := string(cmd)
	fmt.Println("Image Downloading Successfully")
    return output
}

Native("")
