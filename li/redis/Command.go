package redis

type Command struct {
	Command string
	len int
	split string
	exec []string
}