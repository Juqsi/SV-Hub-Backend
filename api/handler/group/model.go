package group

type Group struct {
	Id              string   `json:"id" db:"id"`
	Name            string   `json:"name" db:"name"`
	Invitationtoken string   `json:"invitationtoken" db:"invitationtoken"`
	Members         []Member `json:"members"`
	Parent          string   `json:"parent" db:"parent"`
}

type Member struct {
	Id       string `json:"id" db:"id"`
	Forename string `json:"forename" db:"forename"`
	Lastname string `json:"lastname" db:"lastname"`
	Username string `json:"username" db:"username"`
	Telenum  string `json:"telenum" db:"telenum"`
	Role     string `json:"role" db:"role"`
	UserID   string `json:"user_id" db:"user"`
	Group    string `db:"group"`
}

type MemberRole string

const (
	MEMBER_OWNER     MemberRole = "owner"
	MEMBER_MODERATOR MemberRole = "moderator"
	MEMBER_MEMBER    MemberRole = "member"
	MEMBER_GUEST     MemberRole = "guest"
)

func (m MemberRole) String() string {
	return string(m)
}

func memberRoleFromString(role string) (MemberRole, bool) {
	switch role {
	case MEMBER_OWNER.String():
		return MEMBER_OWNER, true
	case MEMBER_MEMBER.String():
		return MEMBER_MEMBER, true
	case MEMBER_MODERATOR.String():
		return MEMBER_MODERATOR, true
	case MEMBER_GUEST.String():
		return MEMBER_GUEST, true
	default:
		return MemberRole(""), false
	}
}
