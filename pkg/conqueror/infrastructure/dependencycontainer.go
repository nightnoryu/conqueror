package infrastructure

import (
	"context"

	"conqueror/pkg/conqueror/app"
	"conqueror/pkg/conqueror/app/query"
	"conqueror/pkg/conqueror/infrastructure/mysql"

	"github.com/jmoiron/sqlx"
)

type DependencyContainer interface {
	UserService() app.UserService
	SubjectService() app.SubjectService
	TaskService() app.TaskService
	NoteService() app.NoteService

	UserQueryService() query.UserQueryService
}

func NewDependencyContainer(ctx context.Context, db *sqlx.DB) (DependencyContainer, error) {
	conn, err := db.Connx(ctx)
	if err != nil {
		return nil, err
	}

	userRepository := mysql.NewUserRepository(ctx, conn)
	userService := app.NewUserService(userRepository)

	subjectRepository := mysql.NewSubjectRepository(ctx, conn)
	subjectService := app.NewSubjectService(subjectRepository, userRepository)

	taskRepository := mysql.NewTaskRepository(ctx, conn)
	taskService := app.NewTaskService(taskRepository, userRepository)

	noteRepository := mysql.NewNoteRepository(ctx, conn)
	noteService := app.NewNoteService(noteRepository, userRepository)

	userQueryService := mysql.NewUserQueryService(conn)

	return &dependencyContainer{
		userService:    userService,
		subjectService: subjectService,
		taskService:    taskService,
		noteService:    noteService,

		userQueryService: userQueryService,
	}, nil
}

type dependencyContainer struct {
	userService    app.UserService
	subjectService app.SubjectService
	taskService    app.TaskService
	noteService    app.NoteService

	userQueryService query.UserQueryService
}

func (container *dependencyContainer) UserService() app.UserService {
	return container.userService
}

func (container *dependencyContainer) SubjectService() app.SubjectService {
	return container.subjectService
}

func (container *dependencyContainer) TaskService() app.TaskService {
	return container.taskService
}

func (container *dependencyContainer) NoteService() app.NoteService {
	return container.noteService
}

func (container *dependencyContainer) UserQueryService() query.UserQueryService {
	return container.userQueryService
}
