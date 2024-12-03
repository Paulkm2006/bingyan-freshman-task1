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
	e.GET(fmt.Sprintf("/%s/user", ver), controller.UserInfo)         // Get user information
	e.GET(fmt.Sprintf("/%s/verify", ver), controller.SendValidation) // Get user information
	e.DELETE(fmt.Sprintf("/%s/user", ver), controller.UserDelete)    // Delete user

	// post
	e.POST(fmt.Sprintf("/%s/post", ver), controller.CreatePost)       // Create post
	e.GET(fmt.Sprintf("/%s/post", ver), controller.GetPosts)          // Get all posts
	e.GET(fmt.Sprintf("/%s/post/pid", ver), controller.GetPostByPID)  // Get post by id
	e.GET(fmt.Sprintf("/%s/post/uid", ver), controller.GetPostsByUID) // Get post by user id
	e.GET(fmt.Sprintf("/%s/post/nid", ver), controller.GetPostsByNID) // Get post by node id
	e.DELETE(fmt.Sprintf("/%s/post", ver), controller.DeletePost)     // Delete post

	// comment
	e.POST(fmt.Sprintf("/%s/comment", ver), controller.CreateComment)       // Create comment
	e.GET(fmt.Sprintf("/%s/comment/pid", ver), controller.GetCommentsByPID) // Get comment by post id
	e.GET(fmt.Sprintf("/%s/comment/uid", ver), controller.GetCommentsByUID) // Get comment by user id
	e.DELETE(fmt.Sprintf("/%s/comment", ver), controller.DeleteComment)     // Delete comment

	// like
	e.POST(fmt.Sprintf("/%s/like", ver), controller.CreateLike)       // Like
	e.GET(fmt.Sprintf("/%s/like/pid", ver), controller.GetLikesByPID) // Get likes by post id
	e.GET(fmt.Sprintf("/%s/like/uid", ver), controller.GetLikesByUID) // Get likes by user id
	e.DELETE(fmt.Sprintf("/%s/like", ver), controller.DeleteLike)     // Cancel like

	// node
	e.POST(fmt.Sprintf("/%s/node", ver), controller.CreateNode)                  // Create node
	e.POST(fmt.Sprintf("/%s/node/moderator", ver), controller.AddModerator)      // Add moderator
	e.GET(fmt.Sprintf("/%s/node", ver), controller.GetNodes)                     // Get all nodes
	e.GET(fmt.Sprintf("/%s/node/nid", ver), controller.GetNodeByNID)             // Get moderators
	e.DELETE(fmt.Sprintf("/%s/node/moderator", ver), controller.DeleteModerator) // Delete moderator
	e.DELETE(fmt.Sprintf("/%s/node", ver), controller.DeleteNode)                // Delete node

	// follow
	e.POST(fmt.Sprintf("/%s/follow", ver), controller.Follow)        // Follow
	e.GET(fmt.Sprintf("/%s/follow", ver), controller.GetFollows)     // Get follows
	e.GET(fmt.Sprintf("/%s/follower", ver), controller.GetFollowers) // Get followers
	e.DELETE(fmt.Sprintf("/%s/follow", ver), controller.Unfollow)    // Unfollow

}
