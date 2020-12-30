package posts

import (
	"fmt"
)

// Post A blog post
type Post struct {
	Name      string
	Content   []byte
	Extension string
}

func (post Post) String() string {
	return fmt.Sprintf("%v (size: %v)", post.Name, len(post.Content))
}
