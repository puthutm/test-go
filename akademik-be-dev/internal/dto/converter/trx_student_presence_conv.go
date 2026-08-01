// Package converter
package converter

import (
	"unsia.ac.id/akademic_be/internal/dto"
	"unsia.ac.id/akademic_be/internal/model"
)

func ConvertPresenceToResponse(presence model.Presence) *dto.PresenceRes {
	return &dto.PresenceRes{
		UseOpenSession:             presence.UseOpenSession,
		OpenSessionPercentage:      presence.OpenSessionPercentage,
		UseDocumentMaterial:        presence.UseDocumentMaterial,
		DocumentMaterialPercentage: presence.DocumentMaterialPercentage,
		UseQuiz:                    presence.UseQuiz,
		QuizPercentage:             presence.QuizPercentage,
		UseTask:                    presence.UseTask,
		TaskPercentage:             presence.TaskPercentage,
		UseVideo:                   presence.UseVideo,
		VideoPercentage:            presence.VideoPercentage,
		UseUTS:                     presence.UseUTS,
		UTSPercentage:              presence.UTSPercentage,
		UseUAS:                     presence.UseUAS,
		UASPercentage:              presence.UASPercentage,
		UseComment:                 presence.UseComment,
		CommentPercentage:          presence.CommentPercentage,
	}
}

func ConvertTrxStudentPresenceBySessionToResponse(p model.TrxStudentPresenceBySession) dto.TrxStudentPresenceBySessionResponse {
	return dto.TrxStudentPresenceBySessionResponse{
		StudentID:               p.StudentID,
		StudentName:             p.StudentName,
		StudentNIM:              p.StudentNIM,
		OpenSessionPercentage:   p.OpenSessionPercentage,
		DocumentMaterialPercent: p.DocumentMaterialPercent,
		QuizPercentage:          p.QuizPercentage,
		TaskPercentage:          p.TaskPercentage,
		VideoPercentage:         p.VideoPercentage,
		UTSPercentage:           p.UTSPercentage,
		UASPercentage:           p.UASPercentage,
		CommentPercentage:       p.CommentPercentage,
		TotalPercentage:         p.TotalPercentage,
		PresenceFlag:            p.PresenceFlag,
	}
}
