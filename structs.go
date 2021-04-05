package main

type Config struct {
	WebhookURL       string           `yaml:"wh-url"`
	WebhookUsername  string           `yaml:"wh-username"`
	WebhookAvatarURL string           `yaml:"wh-avatar-url"`
	Messages         []*MessageConfig `yaml:"messages"`
}

type MessageConfig struct {
	Username  string       `yaml:"username"`
	AvatarURL string       `yaml:"avatar-url"`
	Content   string       `yaml:"content"`
	File      string       `yaml:"file"`
	Embed     *EmbedConfig `yaml:"embed"`
}

type EmbedConfig struct {
	Color       string              `yaml:"color"`
	Title       string              `yaml:"title"`
	Description string              `yaml:"description"`
	URL         string              `yaml:"url"`
	Fields      []*EmbedFieldConfig `yaml:"fields"`
	Image       string              `yaml:"image"`
	Thumbnail   string              `yaml:"thumbnail"`
}

type EmbedFieldConfig struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}
