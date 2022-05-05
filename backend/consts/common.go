package consts

const (
	CoursePrefix = "course/"
	AvatarPrefix = CoursePrefix + "avatar/"
	PdfPrefix    = CoursePrefix + "pdf/"
	VideoPrefix  = CoursePrefix + "video/"
)

const (
	FileTypeAvatar = "avatar"
	FileTypePdf    = "pdf"
	FileTypeVideo  = "video"
)

const (
	AuthHeader = "Authorization"
	AuthToken  = "jwt-token"
)

const (
	RoleStudent = iota
	RoleTeacher
)

const (
	ExerciseTypeMultipleChoice = "multiple_choice"
	ExerciseTypeJudge          = "judge"
	ExerciseTypeFillsUp        = "fills_up"
)
