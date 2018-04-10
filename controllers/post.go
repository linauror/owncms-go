package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/linauror/owncms-go/models"
)

type PostController struct {
	BaseController
}

func (c *PostController) Info() {
	idstr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idstr)
	post, _ := models.GetPostById(id)
	post.NewView()

	if _tagIds := post.Tag; len(_tagIds) > 0 {
		tagIds := strings.Split(_tagIds, ",")
		tags := models.GetTagsByIds(tagIds)
		c.Data["tags"] = tags
	}

	nextPost, prevPost := models.GetNextAndProvPost(id)

	commentFilter := make(map[string]string)
	commentFilter["pid"] = idstr
	commentLists, _ := models.GetAllComment(int64(1), int64(1000), commentFilter)

	c.Data["nextPost"] = nextPost
	c.Data["prevPost"] = prevPost
	c.Data["commentLists"] = commentLists
	c.Data["post"] = post
	c.TplName = "post.html"
}

func (c *PostController) Comment() {
	comment := models.Comment{}
	comment.Uid = c.Uid
	comment.Pid, _ = c.GetInt("pid")

	post, err := models.GetPostById(comment.Pid)
	if err != nil {
		c.ShowTip(err.Error())
		return
	}

	if c.Uid > 0 {
		comment.Username = c.AuthUser.Username
		comment.Usermail = c.AuthUser.Usermail
		comment.Userurl = c.AuthUser.Userurl
	} else {
		comment.Username = c.GetString("username")
		comment.Usermail = c.GetString("usermail")
		comment.Userurl = c.GetString("userurl")
	}
	comment.Content = c.GetString("content")

	id, err := comment.Pub()
	if err != nil {
		c.ShowTip(err.Error())
		return
	} else {
		post.NewComment()
		c.Redirect(fmt.Sprintf("/post/%d/%s#comments_id_%d", comment.Pid, post.Slug, id), 302)
	}
}
