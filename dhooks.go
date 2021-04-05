package main

import (
	"github.com/akamensky/argparse"
	dgo "github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func fatal(desc string, err error) {
	log.Fatal(desc+"\n", err)
}

func parseArgs() (bool, string) {
	parser := argparse.NewParser("dhooks", "Discord webhook message CLI tool.")

	setup := parser.NewCommand("setup", "Generate a sample config file.")
	output := setup.String("o", "output", &argparse.Options{
		Default: "./config.yaml",
	})

	send := parser.NewCommand("send", "Send a webhook message")
	config := send.String("c", "config", &argparse.Options{
		Required: true,
		Default:  "./config.yaml",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	if setup.Happened() {
		err = ioutil.WriteFile(*output, []byte(ExampleConfig), 0666)
		if err != nil {
			log.Fatal(err)
		}
		return false, ""
	} else if send.Happened() {
		return true, *config
	} else {
		parser.Help("test")
	}
	return false, ""
}

func splitUrl(url string) (id, token string) {
	split := strings.Split(url, "/")
	n := len(split)
	return split[n-2], split[n-1]
}

func handleEmbeds(embed *EmbedConfig) []*dgo.MessageEmbed {
	if embed != nil {
		log.Println(embed)
		color, err := strconv.ParseInt(embed.Color, 16, 64)
		if err != nil {
			color = 0
		}

		var fields []*dgo.MessageEmbedField
		for _, field := range embed.Fields {
			fields = append(fields, &dgo.MessageEmbedField{
				Name:  field.Name,
				Value: field.Value,
			})
		}

		var image *dgo.MessageEmbedImage
		if embed.Image != "" {
			image = &dgo.MessageEmbedImage{URL: embed.Image}
		}

		var thumb *dgo.MessageEmbedThumbnail
		if embed.Thumbnail != "" {
			thumb = &dgo.MessageEmbedThumbnail{URL: embed.Thumbnail}
		}

		return []*dgo.MessageEmbed{{
			URL:         embed.URL,
			Type:        "rich",
			Title:       embed.Title,
			Description: embed.Description,
			Color:       int(color),
			Fields:      fields,
			Image:       image,
			Thumbnail:   thumb,
		}}

	} else {
		return nil
	}
}

func ifExists(s, d string) string {
	if s != "" {
		return s
	} else {
		return d
	}
}

func executeWebhook(session *dgo.Session, config Config) error {
	id, token := splitUrl(config.WebhookURL)

	for _, msg := range config.Messages {

		embeds := handleEmbeds(msg.Embed)
		username := ifExists(msg.Username, config.WebhookUsername)
		avatarURL := ifExists(msg.AvatarURL, config.WebhookAvatarURL)

		params := dgo.WebhookParams{
			Content:         msg.Content,
			Username:        username,
			AvatarURL:       avatarURL,
			File:            msg.File,
			Embeds:          embeds,
			AllowedMentions: nil,
		}
		_, err := session.WebhookExecute(id, token, false, &params)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	send, configFile := parseArgs()
	if !send {
		os.Exit(0)
	}
	log.Println(configFile)

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		fatal("error reading file", err)
	}

	config := Config{}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fatal("error parsing YAML", err)
	}

	session, err := dgo.New()
	if err != nil {
		fatal("error creating dgo session", err)
	}

	err = executeWebhook(session, config)
	if err != nil {
		fatal("error executing webhook", err)
	}
}
