package db

import "context"

type NotesQuerier interface {
	UploadNote(ctx context.Context, args UploadNoteParams) error
	GetNote(ctx context.Context, id string) (string, error)
}

var _ NotesQuerier = (*Queries)(nil)
