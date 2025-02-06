package db

import "context"

type NotesQuerier interface {
	UploadNote(ctx context.Context, args UploadNoteParams) error
	GetNote(ctx context.Context, id string) (string, error)
	GetNotes(ctx context.Context) ([]Note, error)
}

var _ NotesQuerier = (*Queries)(nil)
