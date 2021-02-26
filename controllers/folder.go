package controllers

import (
	"gomvc/models"
	"gomvc/requests"
	"gomvc/services"
	"strconv"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type FolderController struct {
	Ctx     iris.Context
	Service services.UsersService
	GroupService services.GroupsService
	FileFolderService services.FileListFolderService
	LeaderFolder services.LeaderFoldersService
	MemberFolder services.MemberFolderService
	Session *sessions.Session
}

const IDSessionKey = "BACKEND_LOGGING_1990"

func (c *FolderController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(IDSessionKey, 0)
	return userID
}

func (c *FolderController) getCurrentUserRole() int64 {
	userRole := c.Session.GetInt64Default("roleSession", 3)
	return userRole
}

func (c *FolderController) getCurrentUsername() string {
	username := c.Session.GetStringDefault("usernameSession", "")
	return username
}

func (c *FolderController) isLoggedIn() bool {
	u := c.getCurrentUsername()
	id := c.getCurrentUserID()
	if u == "" || id <= 0 {
		return false
	}
	return true
}

var LoginStaticPathView =  mvc.Response { 
	Path: "/user/login",
}

func (c *FolderController) GetListfolder() mvc.Result  {
	if c.isLoggedIn() {
		roleUser := c.getCurrentUserRole()
		userID := c.getCurrentUserID()
		listFolder := c.LeaderFolder.GetByRole(roleUser, userID)

		return mvc.View {
			Name: "folder/index.html",
			Data: iris.Map{
				"Title": "List folder of user",
				"layout": true,
				"listFolder": listFolder,
			},
		}
	}
	return LoginStaticPathView
}

func (c *FolderController) GetMemberfolder() mvc.Result {
	if c.isLoggedIn() {
		idLeader := c.Ctx.URLParam("leader_id")
		idGroup := c.Ctx.URLParam("group_id")
		id, err := strconv.ParseInt(idLeader, 10, 64)
		groupID, error := strconv.ParseInt(idGroup, 10, 64)

		if err != nil || error != nil {
			return mvc.View {
				Name: "404page.html",
				Data: iris.Map{
					"Title": "Not found data",
					"layout": true,
				},
			}
		}

		infoLeader := c.LeaderFolder.GetByID(id)
		infoGroup := c.GroupService.GetByID(groupID)
		if (models.LeaderFolders{}) == infoLeader ||  (models.Groups{}) == infoGroup{
			return mvc.View {
				Name: "404page.html",
				Data: iris.Map{
					"Title": "Not found data",
					"layout": true,
				},
			}
		}

		listMember := c.MemberFolder.GetByIDLeader(id)
		return mvc.View {
			Name: "folder/member.html",
			Data: iris.Map{
				"Title": "List member of leader",
				"layout": true,
				"listMember": listMember,
				"infoLeader": infoLeader,
				"infoGroup": infoGroup,
			},
		}
	}
	return LoginStaticPathView
}

func (c *FolderController) GetCreatefolder() mvc.Result {
	if c.isLoggedIn() {
		msgLeader := c.Session.Get("ErrorsLeaderFolder")
		msgMember := c.Session.Get("ErrorsMemberFolder")
		groups := c.GroupService.GetAll()
		allFolder := c.FileFolderService.GetAll()
		state := c.Ctx.URLParam("state")

		if state != "fail_leader" {
			c.Session.Delete("ErrorsLeaderFolder")
		} 
		if state != "fail_member" {
			c.Session.Delete("ErrorsMemberFolder")
		} 

		return mvc.View {
			Name: "folder/create.html",
			Data: iris.Map{
				"Title": "Create folder page",
				"layout": true,
				"groups": groups,
				"allFolder": allFolder,
				"msg": msgLeader,
				"msgMember": msgMember,
				"state":state,
			},
		}
	}
	return LoginStaticPathView
}

func (c *FolderController) PostFolder() mvc.Result {
	if c.isLoggedIn() {
		var (
			leaderName = c.Ctx.FormValue("leaderName")
			leaderEmail = c.Ctx.FormValue("leaderEmail")
			leaderTelegram = c.Ctx.FormValue("leaderTelegram")
			leaderGroup = c.Ctx.FormValue("leaderGroup")
			leaderFolderName = c.Ctx.FormValue("hddLeaderFolderName")

			assignsMember = c.Ctx.FormValue("assignsMember")
			memberName = c.Ctx.FormValue("memberName")
			memberEmail = c.Ctx.FormValue("memberEmail")
			memberTelegram = c.Ctx.FormValue("memberTelegram")
			memberFolderName = c.Ctx.FormValue("hddMemberFolderName")
		)

		msgLeader := &requests.MessageLeaderFolder{
			LeaderName: leaderName,
			LeaderEmail: leaderEmail,
			LeaderTelegram:  leaderTelegram,
			GroupId:   leaderGroup,
			LeaderFolder: leaderFolderName,
		}

		if msgLeader.ValidateLeaderFolder() == false {
			c.Session.Set("ErrorsLeaderFolder", msgLeader.ErrorsLeader)
			return mvc.Response {
				Path: "/admin/folder?state=fail_leader",
			}
		}  

		if assignsMember == "on" {
			msgMember := &requests.MessageMemberFolder{
				MemberName: memberName,
				MemberEmail: memberEmail,
				MemberTelegram:  memberTelegram,
				MemberFolder: memberFolderName,
			}
			if msgMember.ValidateMemberFolder() == false {
				c.Session.Set("ErrorsMemberFolder", msgMember.ErrorsMember)
				return mvc.Response {
					Path: "/admin/folder?state=fail_member",
				}
			}
		}

		userID := c.getCurrentUserID()
		groupID, _ := strconv.ParseInt(leaderGroup, 10, 64)

		//insert database
		idLeader := c.LeaderFolder.CreateFolder(userID, groupID, leaderName, leaderEmail, leaderTelegram, leaderFolderName)
		if assignsMember == "on" && idLeader > 0 {
			idLeader64 := int64(idLeader)
			insertMember := c.MemberFolder.CreateFolder(idLeader64, memberName, memberEmail, memberTelegram, memberFolderName)
			if insertMember {
				return mvc.Response { 
					Path: "/admin/listfolder?state=success",
				}
			}
			return mvc.Response { 
				Path: "/admin/folder?state=error",
			}
		} else if idLeader > 0 {
			return mvc.Response { 
				Path: "/admin/listfolder?state=success",
			}
		} 
		return mvc.Response { 
			Path: "/admin/folder?state=error",
		}
	}
	return LoginStaticPathView
}
