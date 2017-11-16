package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/djthorpe/EnterpriseTechnologyDay/client"
	"github.com/djthorpe/EnterpriseTechnologyDay/util"
)

var (
	FlagDeveloperKey = flag.String("key", "AIzaSyBv-X-KZMrHSaFa6ooY4E2RdbFR6v_3ubE", "YouTube API Key")
	FlagDebug        = flag.Bool("debug", false, "Debug")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Create a YouTube client
	youtube, err := client.NewYouTubeServiceFromDeveloperKey(*FlagDeveloperKey, *FlagDebug)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(-1)
	}

	videos_call := youtube.API.Search.List("id,snippet").EventType("live").Type("video").MaxResults(20)
	if response, err := videos_call.Do(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(-1)
	} else {
		// Gather video id's into an array
		watching_now := make(map[string]uint64, len(response.Items))
		video_list := ""
		for _, row := range response.Items {
			watching_now[row.Id.VideoId] = 0
			video_list += row.Id.VideoId + ","
		}
		call2 := youtube.API.Videos.List("id,liveStreamingDetails").Id(video_list)
		if response2, err := call2.Do(); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(-1)
		} else {
			for _, row := range response2.Items {
				watching_now[row.Id] = row.LiveStreamingDetails.ConcurrentViewers
			}
		}
		output := util.NewOutput("id", "title", "concurrent viewers")
		for _, row := range response.Items {
			output.AppendMap(map[string]interface{}{
				"id":                 row.Id.VideoId,
				"title":              row.Snippet.Title,
				"concurrent viewers": watching_now[row.Id.VideoId],
			})
		}
		output.RenderASCII()
	}

}
