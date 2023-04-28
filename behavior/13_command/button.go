// @Author Bing 
// @Desc
package main

type Button struct {
	Cmd Command
}

func (b *Button) Press() {
	b.Cmd.Execute()
}
