package router

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/controller"
	"fmt"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	var ver = config.Config.Server.Ver

	// user
	e.POST(fmt.Sprintf("/%s/user/token", ver), controller.UserLogin) //Login and retrieve token
	e.POST(fmt.Sprintf("/%s/user", ver), controller.UserRegister)    // Register
	e.DELETE(fmt.Sprintf("/%s/user", ver), controller.UserDelete)    // Delete user, requires admin token
	e.GET(fmt.Sprintf("/%s/user", ver), controller.UserInfo)         // Get user information, requires token
	e.GET(fmt.Sprintf("/%s/verify", ver), controller.SendValidation) // Get user information, requires admin token

	// post
	e.POST(fmt.Sprintf("/%s/post", ver), controller.CreatePost)       // Create post, requires token
	e.GET(fmt.Sprintf("/%s/post", ver), controller.GetPosts)          // Get post, requires token
	e.GET(fmt.Sprintf("/%s/post/pid", ver), controller.GetPostByPID)  // Get post by id, requires token
	e.GET(fmt.Sprintf("/%s/post/uid", ver), controller.GetPostsByUID) // Get post by user id, requires token
	e.DELETE(fmt.Sprintf("/%s/post", ver), controller.DeletePost)     // Delete post, requires token

	// comment
	e.POST(fmt.Sprintf("/%s/comment", ver), controller.CreateComment)       // Create comment, requires token
	e.GET(fmt.Sprintf("/%s/comment/pid", ver), controller.GetCommentsByPID) // Get comment, requires token
	e.GET(fmt.Sprintf("/%s/comment/uid", ver), controller.GetCommentsByUID) // Get comment by user id, requires token
	e.DELETE(fmt.Sprintf("/%s/comment", ver), controller.DeleteComment)     // Delete comment, requires token

	// like
	e.POST(fmt.Sprintf("/%s/like", ver), controller.CreateLike)       // Like, requires token
	e.GET(fmt.Sprintf("/%s/like/pid", ver), controller.GetLikesByPID) // Get likes by post id, requires token
	e.GET(fmt.Sprintf("/%s/like/uid", ver), controller.GetLikesByUID) // Get likes by user id, requires token
	e.DELETE(fmt.Sprintf("/%s/like", ver), controller.DeleteLike)     // Cancel like, requires token
}
