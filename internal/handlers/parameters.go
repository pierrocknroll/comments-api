package handlers

import "strconv"

// TODO For now there is no check for these parameters

func HandleTargetParameter(target string) (string, error) {
	return target, nil
}

func HandleMessageParameter(message string) (string, error) {
	return message, nil
}

func HandleUserIdParameter(userId string) (int, error) {
	var userIdInt, err = strconv.Atoi(userId)
	return userIdInt, err
}

func HandleCommentIdParameter(commentId string) (int, error) {
	var commentIdInt, err = strconv.Atoi(commentId)
	return commentIdInt, err
}
