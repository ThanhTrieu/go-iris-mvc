package requests

import (
	"regexp"
)

var rxEmail1 = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

var rxEmail2 = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
var rxIdGroup = regexp.MustCompile("^[0-9]{1,3}$")

type MessageLeaderFolder struct {
  LeaderName string
  LeaderEmail string
	LeaderTelegram string
	GroupId string
	LeaderFolder string
  ErrorsLeader  map[string]string
}

type MessageMemberFolder struct {
  MemberName string
  MemberEmail string
	MemberTelegram string
	MemberFolder string
  ErrorsMember  map[string]string
}

func (msgLeader *MessageLeaderFolder) ValidateLeaderFolder() bool {
  msgLeader.ErrorsLeader = make(map[string]string)

  matchEmail := rxEmail1.Match([]byte(msgLeader.LeaderEmail))
  if matchEmail == false {
    msgLeader.ErrorsLeader["LeaderEmail"] = "Please enter a valid email address"
  }

	matchName := isValidEmptyString(msgLeader.LeaderName)
	if matchName == false {
		msgLeader.ErrorsLeader["LeaderName"] = "Not be empty Your name"
	}

	matchTele := isValidEmptyString(msgLeader.LeaderTelegram)
	if matchTele == false {
		msgLeader.ErrorsLeader["LeaderTelegram"] = "Not be empty Your telegram"
	}

	matchGroupId := rxIdGroup.Match([]byte(msgLeader.GroupId))
	if matchGroupId == false {
		msgLeader.ErrorsLeader["GroupId"] = "Not be empty Your group"
	}

	matchFolderName := isValidEmptyString(msgLeader.LeaderFolder)
	if matchFolderName == false {
		msgLeader.ErrorsLeader["LeaderFolder"] = "Not be empty Your folder name"
	}

  return len(msgLeader.ErrorsLeader) == 0
}

func (msg *MessageMemberFolder) ValidateMemberFolder() bool {

  msg.ErrorsMember = make(map[string]string)

  matchEmail := rxEmail2.Match([]byte(msg.MemberEmail))
  if matchEmail == false {
    msg.ErrorsMember["MemberEmail"] = "Please enter a valid email address"
  }

	matchName := isValidEmptyString(msg.MemberName)
	if matchName == false {
		msg.ErrorsMember["MemberName"] = "Not be empty Your name"
	}

	matchTele := isValidEmptyString(msg.MemberTelegram)
	if matchTele == false {
		msg.ErrorsMember["MemberTelegram"] = "Not be empty Your telegram"
	}

	matchFolderName := isValidEmptyString(msg.MemberFolder)
	if matchFolderName == false {
		msg.ErrorsMember["MemberFolder"] = "Not be empty Your folder name"
	}

  return len(msg.ErrorsMember) == 0
}

func isValidEmptyString(s string) bool {
	var (
		hasMinLen  = false
	)
	if len(s) >= 2 {
		hasMinLen = true
	}

	return hasMinLen
}
