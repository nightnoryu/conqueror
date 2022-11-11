package domain

import "conqueror/pkg/common/uuid"

type (
	UserID    uuid.UUID
	SubjectID uuid.UUID

	TimetableID      string
	LessonIntervalID string
	LessonID         string

	TaskID uuid.UUID

	NoteID uuid.UUID
)
