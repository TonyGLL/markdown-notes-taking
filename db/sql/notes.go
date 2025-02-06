package db

import "context"

const uploadNoteQuery = `
	INSERT INTO markdown_notes_schema.notes
	(mk, html)
	VALUES ($1, $2)
`

type UploadNoteParams struct {
	MK   string `json:"mk"`
	HTML string `json:"html"`
}

func (q *Queries) UploadNote(ctx context.Context, args UploadNoteParams) error {
	_, err := q.db.ExecContext(ctx, uploadNoteQuery, args.MK, args.HTML)
	return err
}

const getNoteQuery = `
	SELECT n.html FROM markdown_notes_schema.notes n WHERE n.id = $1
`

func (q *Queries) GetNote(ctx context.Context, id string) (string, error) {
	row := q.db.QueryRowContext(ctx, getNoteQuery, id)
	var html string
	err := row.Scan(&html)
	return html, err
}

const getNotesQuery = `
	SELECT * FROM markdown_notes_schema.notes n WHERE n.deleted IS NOT TRUE
`

func (q *Queries) GetNotes(ctx context.Context) ([]Note, error) {
	rows, err := q.db.QueryContext(ctx, getNotesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	notes := []Note{}
	for rows.Next() {
		var note Note
		if err := rows.Scan(
			&note.ID,
			&note.HTML,
			&note.MK,
			&note.CreatedAt,
			&note.UpdatedAt,
			&note.Deleted,
		); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return notes, nil
}
