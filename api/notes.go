package api

import (
	"database/sql"
	"io"
	"net/http"
	"strings"

	db "github.com/TonyGLL/markdown-note-taking/db/sql"
	"github.com/TonyGLL/markdown-note-taking/util"
	"github.com/gin-gonic/gin"
)

func (s *Server) uploadFile(ctx *gin.Context) {
	//* 1. Obtener el archivo del request
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Archivo no proporcionado"})
		return
	}

	//* 2. Se verifica que el archivo sea tipo .md (Markdown)
	filename := strings.ToLower(file.Filename)
	if !strings.HasSuffix(filename, ".md") {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Solo se permiten archivos .md"})
		return
	}

	//* 3. Validar tamaño (hasta 1 MB)
	if file.Size > 1<<20 { // 1 MB en bytes
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El archivo supera 1 MB"})
		return
	}

	//* 4. Se lee el archivo
	uploadedFile, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el archivo"})
		return
	}
	defer uploadedFile.Close()

	//* 5. Se valida que el archivo no este vacio
	markdownContent, err := io.ReadAll(uploadedFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar el archivo"})
		return
	}

	if len(markdownContent) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El archivo está vacío"})
		return
	}

	//* 6. Se convierte el archivo .md a HTML
	htmlFile := util.MDToHTML(markdownContent)

	args := db.UploadNoteParams{
		HTML: string(htmlFile),
		MK:   string(markdownContent),
	}

	err = s.store.UploadNote(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "El archivo se ha subido exitosamente."})
}

func (s *Server) checkGrammar(ctx *gin.Context) {

}

type getNoteRequest struct {
	ID string `uri:"id" binding:"required"`
}

func (s *Server) getNote(ctx *gin.Context) {
	var req getNoteRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	htmlFile, err := s.store.GetNote(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.Header("Content-Type", "text/html")
	ctx.Data(http.StatusOK, "text/html", []byte(htmlFile))
}
