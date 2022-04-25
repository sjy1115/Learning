package consts

const (
	CoursePrefix = "course/"
	AvatarPrefix = CoursePrefix + "avatar/"
	PdfPrefix    = CoursePrefix + "pdf/"
)

const (
	AuthHeader = "Authorization"
	AuthToken  = "jwt-token"
)

const (
	ROLE_STUDENT = 1
	ROLE_TEACHER = 2
)

const (
	RoleStudent = iota
	RoleTeacher
)
