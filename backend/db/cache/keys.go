package cache

import "fmt"

func UserTokenKey(id int) string {
	return fmt.Sprintf("user:token:%d", id)
}

func SignInKey(chapterId int) string {
	return fmt.Sprintf("chapter:%d", chapterId)
}
