/*
   SLACK FILE-BOT

*/

package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK-BOT-TOKEN", "xoxb-4960522986976-4981832211509-bDOf5jryV74zlgqREYmqAFyj") // oAuth token[your App] to give permission to channel/workspace.

	os.Setenv("CHANNEL-ID", "C04TXDMG9HP") // your channel Id

	api := slack.New(os.Getenv("SLACK-BOT-TOKEN")) // creating New client/ slack bot

	channArr := []string{os.Getenv("CHANNEL-ID")} // channel Id is stored in string array

	fileARR := []string{"testfile.pdf"} // your file to be uploaded

	for i := 0; i < len(fileARR); i++ {
		params := slack.FileUploadParameters{ // Looping over the file . scanning line by line .
			Channels: channArr,   //Refer your channel name
			File:     fileARR[i], // Iterating your file
		}

		file, err := api.UploadFile(params) //UPLOADING THROUGH OUR BOT/CLIENT TOKEN

		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		fmt.Printf("Name: %s , URL: %s", file.Name, file.URL) // printing your file name and url for our refernce to check
	}

}
